package integration_test

import (
	"context"
	"flag"
	"strings"
	"testing"
	"time"

	"github.com/kwilteam/kwil-db/test/integration"
	"github.com/kwilteam/kwil-db/test/specifications"
)

var dev = flag.Bool("dev", false, "run for development purpose (no tests)")

var drivers = flag.String("drivers", "http,cli", "comma separated list of drivers to run")

// Here we make clear the services will be used in each stage
var basicServices = []string{integration.ExtContainer, "node0", "node1", "node2"}
var newServices = []string{integration.Ext3Container, "node3"}

// NOTE: allServices will be sorted by docker-compose(in setup), so the order is not reliable
var allServices = []string{integration.ExtContainer, integration.Ext3Container, "node0", "node1", "node2", "node3"}

func TestLocalDevSetup(t *testing.T) {
	if !*dev {
		t.Skip("skipping local dev setup")
	}

	// running forever for local development

	ctx := context.Background()

	opts := []integration.HelperOpt{
		integration.WithBlockInterval(time.Second),
		integration.WithValidators(4),
		integration.WithNonValidators(0),
	}

	helper := integration.NewIntHelper(t, opts...)
	helper.Setup(ctx, allServices)
	defer helper.Teardown()

	helper.WaitForSignals(t)
}

func TestKwildDatabaseIntegration(t *testing.T) {
	ctx := context.Background()

	opts := []integration.HelperOpt{
		integration.WithBlockInterval(time.Second),
		integration.WithValidators(4),
		integration.WithNonValidators(0),
	}

	testDrivers := strings.Split(*drivers, ",")
	for _, driverType := range testDrivers {
		t.Run(driverType+"_driver", func(t *testing.T) {
			helper := integration.NewIntHelper(t, opts...)
			helper.Setup(ctx, basicServices)
			defer helper.Teardown()

			node0Driver := helper.GetUserDriver(ctx, "node0", driverType)
			node1Driver := helper.GetUserDriver(ctx, "node1", driverType)
			node2Driver := helper.GetUserDriver(ctx, "node2", driverType)

			// Create a new database and verify that the database exists on other nodes
			specifications.DatabaseDeploySpecification(ctx, t, node0Driver)
			specifications.DatabaseVerifySpecification(ctx, t, node1Driver, true)
			specifications.DatabaseVerifySpecification(ctx, t, node2Driver, true)

			specifications.ExecuteDBInsertSpecification(ctx, t, node0Driver)
			specifications.ExecuteDBUpdateSpecification(ctx, t, node1Driver)
			specifications.ExecuteDBDeleteSpecification(ctx, t, node2Driver)

			// specifications.ExecutePermissionedActionSpecification(ctx, t, invalidUserDriver)

			specifications.DatabaseDropSpecification(ctx, t, node1Driver)
		})
	}
}

func TestKwildValidatorRemoval(t *testing.T) {
	ctx := context.Background()

	// In this test, we will have a set of 4 validators, where 3 of the
	// validators are required to remove one.
	const numVals, numNonVals = 4, 0
	opts := []integration.HelperOpt{
		integration.WithValidators(numVals),
		integration.WithNonValidators(numNonVals),
	}

	testDrivers := strings.Split(*drivers, ",")
	for _, driverType := range testDrivers {
		if driverType == "http" {
			continue // admin service cannot use http
		}

		t.Run(driverType+"_driver", func(t *testing.T) {
			helper := integration.NewIntHelper(t, opts...)
			helper.Setup(ctx, allServices)
			defer helper.Teardown()

			node0Driver := helper.GetOperatorDriver(ctx, "node0", driverType)
			node1Driver := helper.GetOperatorDriver(ctx, "node1", driverType)
			node2Driver := helper.GetOperatorDriver(ctx, "node2", driverType)
			targetPubKey := helper.NodePrivateKey("node3").PubKey().Bytes()

			/* Remove node 3 (4 validators, nodes 0, 1, and 2 remove node 3)
			- node 0 votes to remove
			- node 3 is still a validator
			- node 1 votes to remove
			- node 3 is still a validator
			- node 2 votes to remove
			- node 3 is no longer a validator
			*/
			specifications.ValidatorNodeRemoveSpecificationV4R1(ctx, t, node0Driver, node1Driver, node2Driver, targetPubKey) // joiner is a validator at node
		})
	}
}

func TestKwildValidatorUpdatesIntegration(t *testing.T) {
	ctx := context.Background()

	const expiryBlocks = 10
	const blockInterval = time.Second
	const numVals, numNonVals = 3, 1
	opts := []integration.HelperOpt{
		integration.WithValidators(numVals),
		integration.WithNonValidators(numNonVals),
		integration.WithJoinExpiry(expiryBlocks),
		integration.WithBlockInterval(blockInterval),
		integration.WithGas(), // must give the joining node some gas too
	}

	expiryWait := 2 * expiryBlocks * blockInterval

	testDrivers := strings.Split(*drivers, ",")
	for _, driverType := range testDrivers {
		if driverType == "http" {
			continue // admin service cannot use http
		}

		t.Run(driverType+"_driver", func(t *testing.T) {
			helper := integration.NewIntHelper(t, opts...)
			helper.Setup(ctx, allServices)
			defer helper.Teardown()

			node0Driver := helper.GetOperatorDriver(ctx, "node0", driverType)
			node1Driver := helper.GetOperatorDriver(ctx, "node1", driverType)
			joinerDriver := helper.GetOperatorDriver(ctx, "node3", driverType)
			joinerPrivKey := helper.NodePrivateKey("node3")
			joinerPubKey := joinerPrivKey.PubKey().Bytes()

			// Start the network with 3 validators & 1 Non-validator
			specifications.CurrentValidatorsSpecification(ctx, t, node0Driver, 3)

			/*
				Join Expiry:
				- Node3 requests to join
				- No approval from other nodes
				- Join request should expire after 15 blocks
			*/
			specifications.ValidatorJoinExpirySpecification(ctx, t, joinerDriver, joinerPubKey, expiryWait)

			/*
			 Join Process:
			 - Node3 requests to join
			 - Requires at least 2 nodes to approve
			 - Consensus reached, Node3 is a Validator
			*/
			specifications.ValidatorNodeJoinSpecification(ctx, t, joinerDriver, joinerPubKey, 3)
			// Node 0,1 approves
			specifications.ValidatorNodeApproveSpecification(ctx, t, node0Driver, joinerPubKey, 3, 3, false)
			specifications.ValidatorNodeApproveSpecification(ctx, t, node1Driver, joinerPubKey, 3, 4, true)
			specifications.CurrentValidatorsSpecification(ctx, t, node0Driver, 4)

			/*
			 Leave Process:
			 - node3 issues a leave request -> removes it from the validator list
			 - Validator set count should be reduced by 1
			*/
			specifications.ValidatorNodeLeaveSpecification(ctx, t, joinerDriver)

			/*
			 Rejoin: (same as join process)
			*/
			specifications.ValidatorNodeJoinSpecification(ctx, t, joinerDriver, joinerPubKey, 3)
			// Node 0, 1 approves
			specifications.ValidatorNodeApproveSpecification(ctx, t, node0Driver, joinerPubKey, 3, 3, false)
			specifications.ValidatorNodeApproveSpecification(ctx, t, node1Driver, joinerPubKey, 3, 4, true)
			specifications.CurrentValidatorsSpecification(ctx, t, node0Driver, 4)
		})
	}
}

func TestKwildNetworkSyncIntegration(t *testing.T) {
	ctx := context.Background()

	opts := []integration.HelperOpt{
		integration.WithValidators(4),
		integration.WithBlockInterval(time.Second),
	}

	testDrivers := strings.Split(*drivers, ",")
	for _, driverType := range testDrivers {
		t.Run(driverType+"_driver", func(t *testing.T) {
			helper := integration.NewIntHelper(t, opts...)
			helper.Setup(ctx, basicServices)
			defer helper.Teardown()

			node0Driver := helper.GetUserDriver(ctx, "node0", driverType)
			node1Driver := helper.GetUserDriver(ctx, "node1", driverType)
			node2Driver := helper.GetUserDriver(ctx, "node2", driverType)

			// Create a new database and verify that the database exists on other nodes
			specifications.DatabaseDeploySpecification(ctx, t, node0Driver)
			time.Sleep(time.Second * 2) // need time to sync
			specifications.DatabaseVerifySpecification(ctx, t, node1Driver, true)
			specifications.DatabaseVerifySpecification(ctx, t, node2Driver, true)

			// Insert 1 User and 1or2 Posts
			specifications.ExecuteDBInsertSpecification(ctx, t, node0Driver)

			// Spin up node 4 and ensure that the database is synced to node4
			/*
				1. Generate config for node 4: place it in the homedir/newNode
				2. Run docker compose up on the new node and get the container
				3. Get the node driver
				4. Verify that the database exists on the new node
			*/
			helper.RunDockerComposeWithServices(ctx, newServices)
			node3Driver := helper.GetUserDriver(ctx, "node3", driverType)

			/*
			   1. This checks if the database exists on the new node
			   2. Verify if the user and posts are synced to the new node
			*/
			time.Sleep(time.Second * 4) // need time to catch up
			specifications.DatabaseVerifySpecification(ctx, t, node3Driver, true)

			expectPosts := 1
			specifications.ExecuteDBRecordsVerifySpecification(ctx, t, node3Driver, expectPosts)

			// NOTE: integration tests shows that we need somewhere to track the
			// test state, so we can verify across nodes
		})
	}
}
