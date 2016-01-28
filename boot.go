package main

import (
	"log"
	"net/http"
	"os"

	"github.com/deis/riak/chans"
	"github.com/deis/riak/clustersrv"
	"github.com/deis/riak/config"
	"github.com/deis/riak/riak"
)

func main() {
	conf, err := config.Get()
	if err != nil {
		log.Printf("Error getting config (%s)", err)
		os.Exit(1)
	}

	cmdDoneCh := make(chan error)
	serverDoneCh := make(chan error)
	if !conf.RiakMaster {
		// non-bootstrap nodes should start a riak server and join
		log.Printf("Starting as a bootstrap node")
		go func() {
			httpClient := &http.Client{}
			clusterServerURL := clustersrv.URLFromConfig(conf)
			if err != nil {
				cmdDoneCh <- err
				return
			}

			if err := riak.Start(false); err != nil {
				cmdDoneCh <- err
				return
			}

			if err := riak.Join(httpClient, clusterServerURL, true); err != nil {
				cmdDoneCh <- err
				return
			}
			close(cmdDoneCh)
		}()
	} else {
		// bootstrap nodes should start (not join) a riak server and start the cluster server
		log.Printf("Starting as a bootstrap node")

		go func() {
			if err := riak.Start(true); err != nil {
				cmdDoneCh <- err
				return
			}
			close(cmdDoneCh)
		}()
		log.Printf("Cluster server starting on port %d", conf.ClusterServerHTTPPort)
		go clustersrv.Start(conf.ClusterServerHTTPPort, serverDoneCh)
	}

	if err := chans.JoinErrs(serverDoneCh, cmdDoneCh); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
