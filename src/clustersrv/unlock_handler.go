package clustersrv

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

const (
	lockIDKey = "lock_id"
)

func registerUnlockHandler(r *mux.Router, mut *sync.Mutex, lockID *LockID) {
	r.Handle("/lock/{"+lockIDKey+"}", newUnlockHandler(mut, lockID)).Methods("DELETE")
}

func newUnlockHandler(mut *sync.Mutex, lockID *LockID) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lid, ok := mux.Vars(r)[lockIDKey]
		if !ok {
			http.Error(w, "missing lock ID in path", http.StatusBadRequest)
			return
		}
		if !lockID.Equals(lid) {
			http.Error(w, "invalid lock ID "+lid, http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		mut.Unlock()
	})
}
