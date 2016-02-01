package stanchion

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/deis/riak/src/replace"
)

const (
	confFilePath = "/etc/stanchion/stanchion.conf"
)

func Action(ctx *cli.Context) {
	conf, err := getConfig()
	if err != nil {
		log.Printf("Error: getting config (%s)", err)
		os.Exit(1)
	}
	confFile, err := ioutil.ReadFile(confFilePath)
	if err != nil {
		log.Printf("Error: reading config file at %s (%s)", confFilePath, err)
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
	}
	newConfFile := replace.String(string(confFile), replacements...)
	if err := ioutil.WriteFile(confFilePath, []byte(newConfFile), os.ModePerm); err != nil {
		log.Printf("Error: writing new config file to %s (%s)", confFilePath, err)
		os.Exit(1)
	}
}
