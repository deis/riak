package clustersrv

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/arschles/assert"
	"github.com/gorilla/mux"
)

func TestClientAcquireLock(t *testing.T) {
	mut := new(sync.Mutex)
	lockID := NewLockID()
	router := mux.NewRouter()
	registerLockHandler(router, mut, lockID)
	registerUnlockHandler(router, mut, lockID)
	srv := httptest.NewServer(router)
	defer srv.Close()

	baseURL := srv.URL
	lid, err := AcquireLock(http.DefaultClient, baseURL)
	assert.NoErr(t, err)
	assert.Equal(t, lockID.id, lid, "resulting lock IDs")
	assert.NoErr(t, ReleaseLock(http.DefaultClient, baseURL, lid))
}
