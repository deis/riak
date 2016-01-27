package clustersrv

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

func Start(port int, doneCh chan<- error) {
	var mut sync.Mutex
	lockID := NewLockID()
	hostStr := fmt.Sprintf(":%d", port)
	router := mux.NewRouter()
	router.Handle(StartHandlerPath(), NewStartHandler(&mut, lockID)).Methods("POST")
	router.Handle(EndHandlerPath(), NewEndHandler(&mut, lockID)).Methods("DELETE")

	if err := http.ListenAndServe(hostStr, router); err != nil {
		doneCh <- err
	}
}
