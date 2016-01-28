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

func TestUnlockNotLocked(t *testing.T) {
	mut := new(sync.Mutex)
	mut.Lock()
	lockID := NewLockID()
	router := mux.NewRouter()
	registerUnlockHandler(router, mut, lockID)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("DELETE", "/lock", bytes.NewReader(nil))
	assert.NoErr(t, err)
	router.ServeHTTP(w, r)
	assert.Equal(t, w.Code, http.StatusBadRequest, "response code")

	w := httptest.NewRecorder()
	r, err = http.NewRequest("DELETE", "/lock/noexist", bytes.NewReader(nil))
	assert.NoErr(t, err)
	router.ServeHTTP(w, r)
	assert.Equal(t, w.Code, http.StatusBadRequest, "response code")
}

func TestUnlockAlreadyLocked(t *testing.T) {
	mut := new(sync.Mutex)
	mut.Lock()
	lockID := NewLockID()
	lockID.Generate()
	router := mux.NewRouter()
	registerUnlockHandler(router, mut, lockID)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("DELETE", "/lock/"+lockID.id, bytes.NewReader(nil))
	assert.NoErr(t, err)
	router.ServeHTTP(w, r)
	assert.Equal(t, w.Code, http.StatusOK, "response code")
}
