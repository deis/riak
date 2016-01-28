package riak

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/deis/riak/clustersrv"
)

type ErrAcquireLock struct {
	origErr error
}

func (e ErrAcquireLock) Error() string {
	return fmt.Sprintf("Couldn't acquire the cluster level lock (%s)", e.origErr)
}

// Join acquires the cluster level lock (using httpClient issuing requests against clusterServerBaseURL), joins the riak server to the existing riak cluster by running /bin/join_riak, and then releases the same cluster lock.
//
// If lock acquisition failed, returns ErrAcquireLock. In all other cases, Join attempts to release the cluster level lock (and logs failures to do so) and returns. Any error returned other than ErrAcquireLock will be related to starting or executing /bin/join_riak.
//
// If joinWait is true, adds "WAIT_AFTER_JOIN=1" to the env. Otherwise adds "WAIT_AFTER_JOIN=0" to the env.
func Join(httpClient *http.Client, clusterServerBaseURL string, joinWait bool) error {
	lockID, err := clustersrv.AcquireLock(httpClient, clusterServerBaseURL)
	if err != nil {
		return &ErrAcquireLock{origErr: err}
	}
	defer func() {
		if err := clustersrv.ReleaseLock(httpClient, clusterServerBaseURL, lockID); err != nil {
			log.Printf("Error releasing lock ID %s (%s)", lockID, err)
		}
	}()
	joinCmd := exec.Command("/bin/join_riak")
	joinCmd.Stdout = os.Stdout
	joinCmd.Stderr = os.Stderr
	if err := joinCmd.Run(); err != nil {
		return err
	}
	return nil
}
