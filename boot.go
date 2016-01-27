package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"sync"

	"github.com/deis/riak/clustersrv"
	"github.com/gorilla/mux"
)

const (
	clusterServerPortEnvVar = "CLUSTER_SERVER_HTTP_PORT"
)

func main() {
	cmdDoneCh := make(chan error)
	go func() {
		bootCmd := exec.Command("/bin/start_riak")
		bootCmd.Stdout = os.Stdout
		bootCmd.Stderr = os.Stderr
		if err := bootCmd.Run(); err != nil {
			cmdDoneCh <- err
		}
	}()

	serverDoneCh := make(chan error)
	if os.Getenv("RIAK_MASTER") == "1" {
		go func() {
			var mut sync.Mutex
			lockID := clustersrv.NewLockID()
			httpPort, err := strconv.Atoi(os.Getenv(clusterServerPortEnvVar))
			if err != nil {
				httpPort = clustersrv.DefaultHTTPPort
			}
			hostStr := fmt.Sprintf(":%d", httpPort)
			log.Printf("Serving cluster planner on %s", hostStr)
			router := mux.NewRouter()
			router.Handle(clustersrv.StartHandlerPath(), clustersrv.NewStartHandler(&mut, lockID)).Methods("POST")
			router.Handle(clustersrv.EndHandlerPath(), clustersrv.NewEndHandler(&mut, lockID)).Methods("DELETE")

			if err := http.ListenAndServe(hostStr, router); err != nil {
				serverDoneCh <- err
			}
		}()
	}

	select {
	case err := <-cmdDoneCh:
		log.Printf("Error running riak start script (%s)", err)
		os.Exit(1)
	case err := <-serverDoneCh:
		log.Printf("Error running plan/commit server (%s)", err)
		os.Exit(1)
	}
}
