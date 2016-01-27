package clustersrv

import (
	"log"
	"net/http"
	"os/exec"
	"sync"
)

func NewPlanAndCommitHandler(mut *sync.Mutex) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("you must POST to this endpoint"))
		}
		mut.Lock()
		defer mut.Unlock()
		planCmd := exec.Command("riak-admin", "cluster", "plan")
		commitCmd := exec.Command("riak-admin", "cluster", "commit")
		planOut, err := planCmd.CombinedOutput()
		if err != nil {
			http.Error(w, string(planOut), http.StatusInternalServerError)
			return
		}
		log.Printf("Cluster successfully planned")
		log.Println(string(planOut))
		commitOut, err := commitCmd.CombinedOutput()
		if err != nil {
			http.Error(w, string(commitOut), http.StatusInternalServerError)
			return
		}
		log.Printf("Cluster successfully committed")
		log.Println(string(planOut))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
}
