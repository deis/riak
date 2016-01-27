package clustersrv

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// Start starts the cluster server on the given port. The server will block execution of this function as long as it runs. If it stops, Start sends an error on doneCh
func Start(port int, doneCh chan<- error) {
	var mut sync.Mutex
	lockID := NewLockID()
	hostStr := fmt.Sprintf(":%d", port)
	router := mux.NewRouter()
	router.Handle(lockHandlerPath(), newLockHandler(&mut, lockID)).Methods("POST")
	router.Handle(unlockHandlerPath(), newUnlockHandler(&mut, lockID)).Methods("DELETE")

	if err := http.ListenAndServe(hostStr, router); err != nil {
		doneCh <- err
		return
	}
	close(doneCh)
}
