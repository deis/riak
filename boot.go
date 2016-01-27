package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/deis/riak/chans"
	"github.com/deis/riak/clustersrv"
	"github.com/deis/riak/riak"
)

const (
	clusterServerPortEnvVar = "CLUSTER_SERVER_HTTP_PORT"
)

func main() {
	cmdDoneCh := make(chan error)
	go riak.Start(cmdDoneCh)

	serverDoneCh := make(chan error)
	if os.Getenv("RIAK_MASTER") == "1" {
		httpPort, err := strconv.Atoi(os.Getenv(clusterServerPortEnvVar))
		if err != nil {
			httpPort = clustersrv.DefaultHTTPPort
		}
		hostStr := fmt.Sprintf(":%d", httpPort)
		log.Printf("Serving cluster planner on %s", hostStr)
		go clustersrv.Start(httpPort, serverDoneCh)
	}

	if err := chans.JoinErrs(serverDoneCh, cmdDoneCh); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
