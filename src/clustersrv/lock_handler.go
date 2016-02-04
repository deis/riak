package clustersrv

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

func registerLockHandler(router *mux.Router, mut *sync.Mutex, lockID *LockID) {
	router.Handle("/lock", newLockHandler(mut, lockID))
}

func newLockHandler(mut *sync.Mutex, lockID *LockID) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mut.Lock()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(lockID.Generate()))
	})
}
