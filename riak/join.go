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

// Join acquires the cluster level lock (using httpClient issuing requests against clusterServerBaseURL), then takes the following steps:
//
// - Joins the riak server to the existing riak cluster by calling "riak-admin cluster join riak@DISCOVERY_IP", where DISCOVERY_IP is an IP address obtained by lookup up IPs for the deis-riak-discovery host
// - Plans the new cluster and commits by running "riak-admin cluster plan"
// - Commits the plan by running "riak-admin cluster commit"
//
// If lock acquisition failed, returns ErrAcquireLock. In all other cases, Join attempts to release the cluster level lock (and logs failures to do so) and returns. Any error returned other than ErrAcquireLock will be related to starting or executing /bin/join_riak.
func Join(httpClient *http.Client, clusterServerBaseURL string) error {
	discoveryAddr, err := getDiscoveryIP(discoveryHostName)
	if err != nil {
		return err
	}
	log.Printf("Got %s IP %s", discoveryHostName, discoveryAddr)

	log.Printf("Attempting to acquire cluster lock")
	lockID, err := clustersrv.AcquireLock(httpClient, clusterServerBaseURL)
	if err != nil {
		return &ErrAcquireLock{origErr: err}
	}
	defer func() {
		log.Printf("Attempting to release the cluster lock")
		if err := clustersrv.ReleaseLock(httpClient, clusterServerBaseURL, lockID); err != nil {
			log.Printf("Error releasing lock ID %s (%s)", lockID, err)
		} else {
			log.Printf("... released")
		}
	}()
	log.Printf("... acquired")

	log.Printf("Attempting to join the riak cluster")
	joinCmd := exec.Command("riak-admin", "cluster", "join", "riak@"+discoveryAddr)
	joinCmd.Stdout = os.Stdout
	joinCmd.Stderr = os.Stderr
	if err := joinCmd.Run(); err != nil {
		return err
	}
	log.Printf("... joined")

	log.Printf("Attempting to plan the cluster")
	planCmd := exec.Command("riak-admin", "cluster", "plan")
	planCmd.Stdout = os.Stdout
	planCmd.Stderr = os.Stderr
	if err := planCmd.Run(); err != nil {
		return err
	}
	log.Printf("... planned")

	log.Printf("Attempting to commit the new plan")
	commitCmd := exec.Command("riak-admin", "cluster", "commit")
	commitCmd.Stdout = os.Stdout
	commitCmd.Stderr = os.Stderr
	if err := commitCmd.Run(); err != nil {
		return err
	}
	log.Printf("... committed")

	return nil
}
