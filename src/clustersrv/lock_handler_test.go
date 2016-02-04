package clustersrv

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/arschles/assert"
	"github.com/gorilla/mux"
)

func TestNotLocked(t *testing.T) {
	mut := new(sync.Mutex)
	lockID := NewLockID()
	router := mux.NewRouter()
	registerLockHandler(router, mut, lockID)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/lock", bytes.NewReader(nil))
	assert.NoErr(t, err)
	router.ServeHTTP(w, r)

	assert.Equal(t, w.Code, http.StatusOK, "response code")
	assert.Equal(t, string(w.Body.Bytes()), lockID.id, "lock ID")
}
