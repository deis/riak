package clustersrv

import (
	"net/http"
	"sync"
)

func StartHandlerPath() string {
	return "/start"
}

func NewStartHandler(mut *sync.Mutex, lockID *LockID) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mut.Lock()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(lockID.Generate()))
	})
}
