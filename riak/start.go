package riak

import (
	"fmt"
	"os"
	"os/exec"
)

type ErrAcquireLock struct {
	origErr error
}

func (e ErrAcquireLock) Error() string {
	return fmt.Sprintf("Couldn't acquire the cluster level lock (%s)", e.origErr)
}

// Start grabs the cluster lock, starts the riak server, joins it to the riak cluster, then releases the lock and returns. Returns any error (ErrAcquireLock if the lock couldn't be acquired), or nil if everything went well. After any return, successful or otherwise, always attempts once to release the lock, and if that fails, logs that fact
func Start() error {
	bootCmd := exec.Command("/bin/start_riak")
	bootCmd.Stdout = os.Stdout
	bootCmd.Stderr = os.Stderr
	if err := bootCmd.Run(); err != nil {
		return err
	}
	return nil
}
