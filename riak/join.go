package riak

import (
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/deis/riak/clustersrv"
)

func Join(httpClient *http.Client, clusterServerBaseURL string) error {
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
