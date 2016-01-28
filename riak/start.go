package riak

import (
	"os"
	"os/exec"
)

// Start starts the riak server by running /bin/start_riak
func Start() error {
	bootCmd := exec.Command("/bin/start_riak")
	bootCmd.Stdout = os.Stdout
	bootCmd.Stderr = os.Stderr
	if err := bootCmd.Run(); err != nil {
		return err
	}
	return nil
}
