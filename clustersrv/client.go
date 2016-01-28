package clustersrv

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/deis/riak/config"
)

func URLFromConfig(conf *config.Config) string {
	return fmt.Sprintf("http://%s:%d", conf.ClusterServerHost, conf.ClusterServerPort)
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
