package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"sync"

	"github.com/deis/riak/plansrv"
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
			httpPort, err := strconv.Atoi(os.Getenv("PLANSRV_HTTP_PORT"))
			if err != nil {
				httpPort = plansrv.DefaultHTTPPort
			}
			hostStr := fmt.Sprintf(":%d", httpPort)
			log.Printf("Serving cluster planner on %s", hostStr)
			mux := http.NewServeMux()
			mux.Handle("/plan_and_commit", plansrv.NewPlanAndCommitHandler(&mut))

			if err := http.ListenAndServe(hostStr, mux); err != nil {
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
