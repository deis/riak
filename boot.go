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
	bootCmd := exec.Command("/bin/start_riak")
	bootCmd.Stdout = os.Stdout
	bootCmd.Stderr = os.Stderr
	if err := bootCmd.Run(); err != nil {
		log.Printf("Error running riak start script (%s)", err)
		os.Exit(1)
	}

	if os.Getenv("RIAK_MASTER") == "1" {
		var mut sync.Mutex
		httpPort, err := strconv.Atoi(os.Getenv("PLANSRV_HTTP_PORT"))
		if err != nil {
			httpPort = plansrv.DefaultHTTPPort
		}
		hostStr := fmt.Sprintf(":%d", httpPort)
		log.Printf("Serving cluster planner on %s", hostStr)
		mux := http.NewServeMux()
		mux.Handle("/plan_and_commit", plansrv.NewPlanAndCommitHandler(&mut))

		go http.ListenAndServe(hostStr, mux)
	}
}
