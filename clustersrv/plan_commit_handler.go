package clustersrv

import (
	"net/http"
	"os/exec"
	"sync"
)

func NewPlanAndCommitHandler(mut *sync.Mutex) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mut.Lock()
		defer mut.Unlock()
		planCmd := exec.Command("riak-admin", "cluster", "plan")
		commitCmd := exec.Command("riak-admin", "cluster", "commit")
		planOut, err := planCmd.CombinedOutput()
		if err != nil {
			http.Error(w, string(planOut), http.StatusInternalServerError)
			return
		}
		commitOut, err := commitCmd.CombinedOutput()
		if err != nil {
			http.Error(w, string(commitOut), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
}
