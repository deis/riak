package clustersrv

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/arschles/assert"
	"github.com/gorilla/mux"
)

func TestSuccessfulRoundTrip(t *testing.T) {
	mut := new(sync.Mutex)
	lockID := NewLockID()
	router := mux.NewRouter()
	registerLockHandler(router, mut, lockID)
	registerUnlockHandler(router, mut, lockID)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/lock", bytes.NewReader(nil))
	assert.NoErr(t, err)
	router.ServeHTTP(w, r)
	assert.Equal(t, w.Code, http.StatusOK, "response code")
	idStr := string(w.Body.Bytes())
	assert.True(t, len(idStr) > 0, "returned lock ID was empty")
	assert.True(t, lockID.Equals(idStr), fmt.Sprintf("internal lock ID (%s) was not equal returned lock ID (%s)", lockID.id, idStr))

	w = httptest.NewRecorder()
	r, err = http.NewRequest("DELETE", "/lock/"+idStr, bytes.NewReader(nil))
	assert.NoErr(t, err)
	router.ServeHTTP(w, r)
	assert.Equal(t, w.Code, http.StatusOK, "response code")
}
