// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphql

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	// "github.com/net-auto/resourceManager/graph/graphql/directive"
	"github.com/facebookincubator/symphony/pkg/log"
	"github.com/facebookincubator/symphony/pkg/telemetry/ocgql"
	"github.com/facebookincubator/symphony/pkg/viewer"
	"github.com/net-auto/resourceManager/ent"
	"github.com/net-auto/resourceManager/ent/privacy"
	"github.com/net-auto/resourceManager/graph/graphql/generated"
	"github.com/net-auto/resourceManager/graph/graphql/resolver"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.uber.org/zap"
)

// HandlerConfig configures graphql handler.
type HandlerConfig struct {
	Client *ent.Client
	Logger log.Logger
}

func init() {
	// TODO directive
	// views := append(
	// 	ocgql.DefaultServerViews,
	// 	directive.ServerDeprecatedCountByObjectInputField,
	// )
	views := ocgql.DefaultServerViews

	for _, v := range views {
		v.TagKeys = append(v.TagKeys,
			viewer.KeyTenant,
			viewer.KeyUser,
			viewer.KeyUserAgent,
		)
	}
}

// NewHandler creates a graphql http handler.
func NewHandler(cfg HandlerConfig) (http.Handler, func(), error) {
	rsv := resolver.New(
		resolver.Config{
			Client: cfg.Client,
			Logger: cfg.Logger,
		},
	)

	// TODO directive
	// views := append(
	// 	ocgql.DefaultServerViews,
	// 	directive.ServerDeprecatedCountByObjectInputField,
	// )
	views := ocgql.DefaultServerViews

	if err := view.Register(views...); err != nil {
		return nil, nil, fmt.Errorf("registering views: %w", err)
	}
	closer := func() { view.Unregister(views...) }

	router := mux.NewRouter()
	router.Use(func(handler http.Handler) http.Handler {
		timeouter := http.TimeoutHandler(handler, 3*time.Minute, "request timed out")
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h := timeouter
			if websocket.IsWebSocketUpgrade(r) {
				h = handler
			}
			h.ServeHTTP(w, r)
		})
	})

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: rsv,
				// TODO directive
				// Directives: directive.New(cfg.Logger),
				Directives: generated.DirectiveRoot{},
			},
		),
	)
	srv.SetErrorPresenter(errorPresenter(cfg.Logger))
	srv.SetRecoverFunc(recoverFunc(cfg.Logger))
	srv.Use(ocgql.Tracer{})
	srv.Use(ocgql.Metrics{})

	router.Path("/graphiql").
		Handler(
			ochttp.WithRouteTag(
				playground.Handler(
					"GraphQL playground",
					"/graph/query",
				),
				"graphiql",
			),
		)
	router.Path("/query").
		Handler(
			ochttp.WithRouteTag(
				gziphandler.GzipHandler(srv),
				"query",
			),
		)

	return router, closer, nil
}

func errorPresenter(logger log.Logger) graphql.ErrorPresenterFunc {
	return func(ctx context.Context, err error) *gqlerror.Error {
		gqlerr := graphql.DefaultErrorPresenter(ctx, err)
		if strings.Contains(err.Error(), privacy.Deny.Error()) {
			gqlerr.Message = "Permission denied"
		} else if _, ok := err.(*gqlerror.Error); !ok {
			logger.For(ctx).Error("graphql internal error", zap.Error(err))
			gqlerr.Message = "Sorry, something went wrong"
		}
		return gqlerr
	}
}

func recoverFunc(logger log.Logger) graphql.RecoverFunc {
	return func(ctx context.Context, err interface{}) error {
		logger.For(ctx).Error("graphql panic recovery",
			zap.Any("error", err),
			zap.Stack("stack"),
		)
		return errors.New("internal system error")
	}
}
