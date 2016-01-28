package clustersrv

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	ClusterServerHostDiscoveryEnvVar = "DEIS_RIAK_CLUSTER_SERVICE_HOST"
	ClusterServerPortDiscoveryEnvVar = "DEIS_RIAK_CLUSTER_SERVICE_PORT"
)

var (
	ErrMissingHostDiscoveryEnvVar = errors.New("missing cluster server host environment variable")
	ErrMissingPortDiscoveryEnvVar = errors.New("missing cluster server port environment variable")
)

func ClusterServerURLFromEnv() (string, error) {
	clusterSrvHost := os.Getenv(ClusterServerHostDiscoveryEnvVar)
	clusterSrvPort := os.Getenv(ClusterServerPortDiscoveryEnvVar)
	if clusterSrvHost == "" {
		return "", ErrMissingHostDiscoveryEnvVar
	}
	if clusterSrvPort == "" {
		return "", ErrMissingPortDiscoveryEnvVar
	}
	return clusterSrvHost + ":" + clusterSrvPort, nil
}

func AcquireLock(httpClient *http.Client, clusterSrvURLBase string) (string, error) {
	urlStr := fmt.Sprintf("%s/lock", clusterSrvURLBase)
	req, err := http.NewRequest("POST", urlStr, bytes.NewReader(nil))
	if err != nil {
		return "", err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	lockID, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(lockID), nil
}

func ReleaseLock(httpClient *http.Client, clusterSrvURLBase, lockID string) error {
	urlStr := fmt.Sprintf("%s/lock/%s", clusterSrvURLBase, lockID)
	req, err := http.NewRequest("DELETE", urlStr, bytes.NewReader(nil))
	if err != nil {
		return err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
