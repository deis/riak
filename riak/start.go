package riak

import (
	"os"
	"os/exec"
)

// Start starts the riak server by running /bin/start_riak. If startWait is true, passes WAIT_AFTER_START in the env to the script. Otherwise passes WAIT_AFTER_START=0. Returns any error starting or running the script.
func Start(startWait bool) error {
	bootCmd := exec.Command("/bin/start_riak")
	bootCmd.Stdout = os.Stdout
	bootCmd.Stderr = os.Stderr
	startWaitEnv := startWaitEnvName + "=0"
	if startWait {
		startWaitEnv = startWaitEnvName + "=1"
	}
	bootCmd.Env = append(bootCmd.Env, startWaitEnv)
	if err := bootCmd.Run(); err != nil {
		return err
	}
	return nil
}
