package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/deis/riak/chans"
	"github.com/deis/riak/clustersrv"
	"github.com/deis/riak/riak"
)

const (
	clusterServerPortEnvVar = "CLUSTER_SERVER_HTTP_PORT"
	riakMasterEnvVar        = "RIAK_MASTER"
)

func main() {
	cmdDoneCh := make(chan error)
	if os.Getenv(riakMasterEnvVar) != "1" {
		log.Printf("Starting as a member node")
		// non-bootstrap nodes should start a riak server and join
		go func() {
			httpClient := &http.Client{}
			clusterServerURL, err := clustersrv.ClusterServerURLFromEnv()
			if err != nil {
				cmdDoneCh <- err
				return
			}

			if err := riak.Start(); err != nil {
				cmdDoneCh <- err
				return
			}

			if err := riak.Join(httpClient, clusterServerURL); err != nil {
				cmdDoneCh <- err
				return
			}
			close(cmdDoneCh)
		}()
	} else {
		log.Printf("Starting as a bootstrap node")
		// bootstrap nodes should just start a riak server
		go func() {
			if err := riak.Start(); err != nil {
				cmdDoneCh <- err
				return
			}
			close(cmdDoneCh)
		}()
	}

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
