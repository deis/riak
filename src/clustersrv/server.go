package clustersrv

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// Start starts the cluster server on the given port. The server will block execution of this function as long as it runs. If it stops, Start sends an error on doneCh
func Start(port int, doneCh chan<- error) {
	mut := new(sync.Mutex)
	lockID := NewLockID()
	hostStr := fmt.Sprintf(":%d", port)
	router := mux.NewRouter()
	registerLockHandler(router, mut, lockID)
	registerUnlockHandler(router, mut, lockID)

	if err := http.ListenAndServe(hostStr, router); err != nil {
		doneCh <- err
		return
	}
	close(doneCh)
}
