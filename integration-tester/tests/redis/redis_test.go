package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"testing"

	"github.com/kelda/kelda/integration-tester/util"
)

const (
	// The hostname of the master of the Redis cluster. This is how we identify
	// the container that we should run the test from.
	masterHostname = "redis-master"

	masterPassword     = "password"
	expConnectedSlaves = 3
)

var connectedSlavesRegex = regexp.MustCompile(`connected_slaves:(\d+)`)

func TestRedis(t *testing.T) {
	clnt, err := util.GetDefaultDaemonClient()
	if err != nil {
		t.Fatalf("couldn't get api client: %s", err)
	}
	defer clnt.Close()

	containers, err := clnt.QueryContainers()
	if err != nil {
		t.Fatalf("couldn't query containers: %s", err)
	}

	var redisMasterID string
	for _, c := range containers {
		if c.Hostname == masterHostname {
			redisMasterID = c.BlueprintID
			break
		}
	}
	if redisMasterID == "" {
		t.Fatal("Failed to find master container")
	}

	redisInfoBytes, err := exec.Command("kelda", "ssh", redisMasterID,
		"redis-cli", "-a", masterPassword, "info").CombinedOutput()
	redisInfo := string(redisInfoBytes)
	fmt.Printf("Master node info:\n%s\n", redisInfo)
	if err != nil {
		t.Fatalf("Failed to get deployment info: %s", err.Error())
	}

	connectedSlavesMatch := connectedSlavesRegex.FindAllStringSubmatch(redisInfo, -1)
	if len(connectedSlavesMatch) != 1 {
		t.Fatal("Failed to find number of connected slaves")
	}

	connectedSlavesStr := connectedSlavesMatch[0][1]
	connectedSlaves, err := strconv.Atoi(connectedSlavesStr)
	if err != nil {
		t.Fatalf("Failed to parse number of connected slaves (%s): %s",
			connectedSlavesStr, err)
	}

	if connectedSlaves != expConnectedSlaves {
		t.Fatalf("Wrong number of connected slaves: expected %d, got %d",
			expConnectedSlaves, connectedSlaves)
	}
}
