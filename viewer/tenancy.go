// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copied from inventory and modified to use our ent.Client

package viewer

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/net-auto/resourceManager/pools/allocating_strategies"
	"github.com/net-auto/resourceManager/psql"
	"runtime"
	"sync"

	"github.com/facebook/ent/dialect"
	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/facebookincubator/symphony/pkg/ent/migrate"
	"github.com/facebookincubator/symphony/pkg/log"
	fb_viewer "github.com/facebookincubator/symphony/pkg/viewer"
	"github.com/net-auto/resourceManager/ent"

	"github.com/jackc/pgx"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"gocloud.dev/server/health"
	"gocloud.dev/server/health/sqlhealth"
)

// Tenancy provides tenant client for key.
type Tenancy interface {
	ClientFor(context.Context, string) (*ent.Client, error)
}

// FixedTenancy returns a fixed client.
type FixedTenancy struct {
	client *ent.Client
}

// NewFixedTenancy creates fixed tenancy from client.
func NewFixedTenancy(client *ent.Client) FixedTenancy {
	return FixedTenancy{client}
}

// ClientFor implements Tenancy interface.
func (f FixedTenancy) ClientFor(context.Context, string) (*ent.Client, error) {
	return f.Client(), nil
}

// Client returns the client stored in fixed tenancy.
func (f FixedTenancy) Client() *ent.Client {
	return f.client
}

// CacheTenancy is a tenancy wrapper cashing underlying clients.
type CacheTenancy struct {
	tenancy  Tenancy
	initFunc func(*ent.Client)
	clients  map[string]*ent.Client
	mu       sync.RWMutex
}

// NewCacheTenancy creates a tenancy cache.
func NewCacheTenancy(tenancy Tenancy, initFunc func(*ent.Client)) *CacheTenancy {
	return &CacheTenancy{
		tenancy:  tenancy,
		initFunc: initFunc,
		clients:  map[string]*ent.Client{},
	}
}

// ClientFor implements Tenancy interface.
func (c *CacheTenancy) ClientFor(ctx context.Context, name string) (*ent.Client, error) {
	c.mu.RLock()
	client, ok := c.clients[name]
	c.mu.RUnlock()
	if ok {
		return client, nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if client, ok := c.clients[name]; ok {
		return client, nil
	}
	client, err := c.tenancy.ClientFor(ctx, name)
	if err != nil {
		return client, err
	}
	if c.initFunc != nil {
		c.initFunc(client)
	}
	c.clients[name] = client
	return client, nil
}

// SetLogger sets tenancy logger.
func (m *PsqlTenancy) SetLogger(logger log.Logger) {
	m.logger = logger
}

// CheckHealth implements health.Checker interface.
func (c *CacheTenancy) CheckHealth() error {
	if checker, ok := c.tenancy.(health.Checker); ok {
		return checker.CheckHealth()
	}
	return nil
}

// MySQLTenancy provides logical database per tenant.
type PsqlTenancy struct {
	health.Checker
	logger        log.Logger
	dsnUrl        string
	tenantMaxConn int
	closers       []func()
	tenantManager *TenantService
}

// NewPsqlTenancy creates psql tenancy for data source.
func NewPsqlTenancy(dsn string, tenantMaxConn int) (*PsqlTenancy, error) {
	_, err := pgx.ParseConnectionString(dsn)
	if err != nil {
		return nil, fmt.Errorf("parsing dsn: %w", err)
	}
	db := psql.Open(dsn)
	checker := sqlhealth.New(db)
	tenancy := &PsqlTenancy{
		Checker:       checker,
		dsnUrl:        dsn,
		tenantMaxConn: tenantMaxConn,
		logger:        log.NewNopLogger(),
		closers:       []func(){checker.Stop},
		tenantManager: NewTenantService(db),
	}
	runtime.SetFinalizer(tenancy, func(tenancy *PsqlTenancy) {
		for _, closer := range tenancy.closers {
			closer()
		}
	})
	return tenancy, nil
}

// ClientFor implements Tenancy interface.
func (m *PsqlTenancy) ClientFor(ctx context.Context, name string) (*ent.Client, error) {
	client := ent.NewClient(ent.Driver(entsql.OpenDB(dialect.Postgres, m.dbFor(name))))

	if exists, err := m.tenantManager.Exist(ctx, name); !exists && err == nil {
		m.logger.For(ctx).Info("Creating db for new tenant", zap.String("tenant", name))
		if _, err := m.tenantManager.Create(ctx, name); err != nil {
			return nil, err
		}
	} else if exists && err == nil {
		// Do nothing, db already in place
	} else {
		return nil, err
	}

	m.logger.For(ctx).Debug("Invoking db migration for tenant", zap.String("tenant", name))
	if err := m.migrate(ctx, client); err != nil {
		return nil, err
	}

	m.logger.For(ctx).Debug("Loading built-in resource types for tenant", zap.String("tenant", name))
	if err := pools.LoadBuiltinTypes(ctx, client); err != nil {
		return nil, err
	}

	return client, nil
}

func (m *PsqlTenancy) migrate(ctx context.Context, client *ent.Client) error {
	ctx, span := trace.StartSpan(ctx, "tenancy.Migrate")
	defer span.End()
	if err := client.Schema.Create(ctx,
		migrate.WithFixture(false),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		m.logger.For(ctx).Error("tenancy migrate", zap.Error(err))
		span.SetStatus(trace.Status{Code: trace.StatusCodeUnknown, Message: err.Error()})
		return fmt.Errorf("running tenancy migration: %w", err)
	}
	return nil
}

func (m *PsqlTenancy) dbFor(name string) *sql.DB {
	dbName, err := psql.ReplaceDbName(m.dsnUrl, fb_viewer.DBName(name))
	if (err != nil) {
		m.logger.Background().Error("Unable to create DB for tenant: ", zap.Error(err))
		panic("Unable to create DB for tenant: " + err.Error())
	}
	m.dsnUrl = dbName
	db := psql.Open(m.dsnUrl)
	db.SetMaxOpenConns(m.tenantMaxConn)
	m.closers = append(m.closers, psql.RecordStats(db))
	return db
}
