package riak

import (
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
	"github.com/deis/riak/src/chans"
	"github.com/deis/riak/src/clustersrv"
)

func Action(ctx *cli.Context) {
	conf, err := getConfig()
	if err != nil {
		log.Printf("Error getting config (%s)", err)
		os.Exit(1)
	}

	cmdDoneCh := make(chan error)
	serverDoneCh := make(chan error)
	if !conf.RiakMaster {
		// non-bootstrap nodes should start a riak server and join
		log.Printf("Starting as a non-bootstrap node")
		go func() {
			httpClient := &http.Client{}
			clusterServerURL := clustersrv.URLFromConfig(conf.ClusterServerHost, conf.ClusterServerPort)
			if err != nil {
				cmdDoneCh <- err
				return
			}

			if err := Start(); err != nil {
				cmdDoneCh <- err
				return
			}

			if err := Join(httpClient, clusterServerURL); err != nil {
				cmdDoneCh <- err
				return
			}
		}()
	} else {
		// bootstrap nodes should start (not join) a riak server and start the cluster server
		log.Printf("Starting as a bootstrap node")

		go func() {
			if err := Start(); err != nil {
				cmdDoneCh <- err
				return
			}
		}()
		log.Printf("Cluster server starting on port %d", conf.ClusterServerHTTPPort)
		go clustersrv.Start(conf.ClusterServerHTTPPort, serverDoneCh)
	}

	if err := chans.JoinErrs(serverDoneCh, cmdDoneCh); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
