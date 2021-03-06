import {deleteAllocationStrategy, testStrategy, createAllocationStrategy, findAllocationStrategyId} from '../graphql-queries';
import {getUniqueName} from '../test-helpers';

test('create and call JS strategy', async () => {
    let poolName = getUniqueName('testJSstrategy');
    let strategyId = await createAllocationStrategy(
        poolName,
        'function invoke() {return {vlan: userInput.desiredVlan};}',
        'js');
    let strategyOutput = await testStrategy(strategyId, { ResourcePoolName: 'testpool'}, poolName, [], {desiredVlan: 85} );
    expect(strategyOutput.stdout.vlan).toBe(85);
});

test('create and call Py strategy', async () => {
    let poolName = getUniqueName('testJSstrategy');
    let strategyId = await createAllocationStrategy(
        poolName,
        'log(json.dumps({ \'respool\': resourcePool[\'ResourcePoolName\'], \'currentRes\': currentResources }))\nreturn {\'vlan\': userInput[\'desiredVlan\']}',
        'py');
    let strategyOutput = await testStrategy(strategyId,
        { ResourcePoolName: poolName},
        poolName, [], {desiredVlan: 11} );
    expect(strategyOutput.stdout.vlan).toBe(11);
});

test('delete strategy', async () => {
    let poolName = getUniqueName('testJSstrategy');
    let strategyId = await createAllocationStrategy(
        poolName,
        'function invoke() {return {vlan: userInput.desiredVlan};}', 'js');
    let foundStrategyId = await findAllocationStrategyId(poolName);
    expect(foundStrategyId).toBe(strategyId);
    await deleteAllocationStrategy(strategyId);
    foundStrategyId = await findAllocationStrategyId(poolName);
    expect(foundStrategyId).not.toBeTruthy();
});

test('simple ipv4 prefix strategy', async () => {
    let poolName = getUniqueName('testJSstrategy');
    let ipv4PrefixStrategyId = await findAllocationStrategyId('ipv4_prefix');
    let x = await testStrategy(ipv4PrefixStrategyId,
        {prefix: 8, address: '10.0.0.0'},
        poolName,
        [], {desiredSize: 8388608});
    expect(x.stdout.address).toBe('10.0.0.0');
    expect(x.stdout.prefix).toBe(9);
});

test('ipv4 prefix strategy one resource already claimed', async () => {
    let poolName = getUniqueName('testJSstrategy');
    let ipv4PrefixStrategyId = await findAllocationStrategyId('ipv4_prefix');
    let allocated = await testStrategy(ipv4PrefixStrategyId,
        {prefix: 8, address: '10.0.0.0'},
        poolName,
        [{Properties: { prefix: 9, address: '10.0.0.0'},
            Status: 'claimed',
            UpdatedAt: '2020-08-18 11:38:48.0 +0200 CEST'
        }], {desiredSize: 8388608});
    expect(allocated.stdout.address).toBe('10.128.0.0');
    expect(allocated.stdout.prefix).toBe(9);
});

test('ipv4 prefix strategy pool has no capacity left', async () => {
    const poolName = getUniqueName('testJSstrategy');
    const allocatedResources = [
        {Properties: { prefix: 9, address: '10.0.0.0'},
            Status: 'claimed',
            UpdatedAt: '2020-08-18 11:38:48.0 +0200 CEST'
        },
        {Properties: { prefix: 9, address: '10.128.0.0'},
            Status: 'claimed',
            UpdatedAt: '2020-08-18 11:38:48.0 +0200 CEST'
        }];
    let ipv4PrefixStrategyId = await findAllocationStrategyId('ipv4_prefix');
    let allocated = await testStrategy(ipv4PrefixStrategyId,
        {prefix: 8, address: '10.0.0.0'},
        poolName, allocatedResources, {desiredSize: 8388608});
    expect(allocated).not.toBeTruthy();
});

test('ipv4 strategy just get an IP', async () => {
    const poolName = getUniqueName('testJSstrategy');
    let ipv4StrategyId = await findAllocationStrategyId('ipv4');
    let allocated = await testStrategy(ipv4StrategyId,
        {prefix: 8, address: '10.0.0.0'},
        poolName, [], {});
    expect(allocated.stdout.address).toBe('10.0.0.0');
});


test('simple ipv6 prefix strategy', async () => {
    let poolName = getUniqueName('testJSstrategy');
    let ipv6PrefixStrategyId = await findAllocationStrategyId('ipv6_prefix');
    let allocated = await testStrategy(ipv6PrefixStrategyId,
        {prefix: 120, address: 'dead::'},
        poolName,
        [], {desiredSize: 101});
    expect(allocated.stdout.address).toBe('dead::');
    expect(allocated.stdout.prefix).toBe(121);
});

test('simple ipv6 strategy', async () => {
    let poolName = getUniqueName('testJSstrategy');
    let ipv6StrategyId = await findAllocationStrategyId('ipv6');
    let allocated = await testStrategy(ipv6StrategyId,
        {prefix: 120, address: 'dead::'},
        poolName,
        [], {subnet: true});
    expect(allocated.stdout.address).toBe('dead::1');
});

test('ipv4-rd strategy', async () => {
    let poolName = getUniqueName('testJSstrategy');
    let strategyId = await findAllocationStrategyId('route_distinguisher');
    let allocated = await testStrategy(strategyId,
        {},
        poolName,
        [], {ipv4: '1.2.3.4', assignedNumber: 2});

    expect(allocated.stdout.rd).toBe('1.2.3.4:2');
});

test('ipv4-rd strategy duplicate already claimed', async () => {
    let poolName = getUniqueName('testJSstrategy');
    const claimed = [{Properties: {rd: '1.2.3.4:2'},
        Status: 'claimed',
        UpdatedAt: '2020-08-30 11:38:48.0 +0200 CEST'
    }];

    let strategyId = await findAllocationStrategyId('route_distinguisher');
    let allocated = await testStrategy(strategyId,
        {},
        poolName,
        claimed, {ipv4: '1.2.3.4', assignedNumber: 2});

    expect(allocated).not.toBeTruthy();
});


test('as-rd strategy', async () => {
    let poolName = getUniqueName('testJSstrategy');
    let strategyId = await findAllocationStrategyId('route_distinguisher');
    let allocated = await testStrategy(strategyId,
        {},
        poolName,
        [], {asNumber: 22, assignedNumber: 288});

    expect(allocated.stdout.rd).toBe('22:288');
});

test('vlan range strategy', async () => {
    let poolName = getUniqueName('testJSstrategy');
    let strategyId = await findAllocationStrategyId('vlan_range');
    let allocated = await testStrategy(strategyId,
        {from: 0, to: 4095},
        poolName,
        [], {desiredSize: 101});

    expect(allocated.stdout).toMatchObject({from: 0, to:100});
});

test('vlan range strategy range partly claimed', async () => {
    let poolName = getUniqueName('testJSstrategy');
    let strategyId = await findAllocationStrategyId('vlan_range');
    const claimed = [
        {
            Properties: {from: 0, to: 1000},
            Status: 'claimed',
            UpdatedAt: '2020-08-30 11:38:48.0 +0200 CEST'
        },];

        let allocated = await testStrategy(strategyId,
        {from: 0, to: 4095},
        poolName,
        claimed, {desiredSize: 101});

    expect(allocated.stdout).toMatchObject({from: 1001, to:1101});
});

test('vlan strategy', async () => {
    let poolName = getUniqueName('testJSstrategy');
    let strategyId = await findAllocationStrategyId('vlan');
    let allocated = await testStrategy(strategyId,
        {from: 0, to: 4095},
        poolName,
        [], {});

    expect(allocated.stdout.vlan).toBe(0);
});
