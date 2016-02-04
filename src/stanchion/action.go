package stanchion

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
	"github.com/deis/riak/src/replace"
)

func Action(ctx *cli.Context) {
	conf, err := getConfig()
	if err != nil {
		log.Printf("Error: getting config (%s)", err)
		os.Exit(1)
	}
	confFile, err := ioutil.ReadFile(conf.ConfFilePath)
	if err != nil {
		log.Printf("Error: reading config file at %s (%s)", conf.ConfFilePath, err)
		os.Exit(1)
	}

	adminKey, err := ioutil.ReadFile(conf.AdminKeyLocation)
	if err != nil {
		log.Printf("Error: reading admin key at %s (%s)", conf.AdminKeyLocation, err)
		os.Exit(1)
	}
	adminSecret, err := ioutil.ReadFile(conf.AdminSecretLocation)
	if err != nil {
		log.Printf("Error: reading admin secret at %s (%s)", conf.AdminSecretLocation, err)
		os.Exit(1)
	}

	replacements := []replace.Replacement{
		replace.FmtReplacement("listener = 127.0.0.1:8085", "listener = %s:%d", conf.ListenHost, conf.ListenPort),
		replace.FmtReplacement("admin.key = admin-key", "admin.key = %s", string(adminKey)),
		replace.FmtReplacement("admin.secret = admin-secret", "admin.secret = %s", string(adminSecret)),
		replace.FmtReplacement("erlang.max_ports = 65536", "erlang.max_ports = %d", conf.NumPorts),
		replace.FmtReplacement("riak_host = 127.0.0.1:8087", "riak_host = %s:%d", conf.RiakHost, conf.RiakPort),
	}
	newConfFile := replace.String(string(confFile), replacements...)
	if err := ioutil.WriteFile(conf.ConfFilePath, []byte(newConfFile), os.ModePerm); err != nil {
		log.Printf("Error: writing new config file to %s (%s)", conf.ConfFilePath, err)
		os.Exit(1)
	}
	log.Printf("Final config file:\n%s", newConfFile)

	log.Printf("Starting Riak Stanchion...")
	startCmd := exec.Command("start_riak_stanchion")
	startCmd.Stdout = os.Stdout
	startCmd.Stderr = os.Stderr
	if err := startCmd.Start(); err != nil {
		log.Printf("Error: starting Riak Stanchion (%s)", err)
		os.Exit(1)
	}
	if err := startCmd.Wait(); err != nil {
		log.Printf("Error: running Riak Stanchion (%s)", err)
		os.Exit(1)
	}
	log.Printf("Error: Riak Stanchion exited without error, should run forever")
	os.Exit(1)
}
