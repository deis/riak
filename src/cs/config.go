package cs

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	AppName = "deis-riak-cs"
)

type config struct {
	ListenHost          string `envconfig:"LISTEN_HOST" required:"true"`
	ListenPort          int    `envconfig:"LISTEN_PORT" required:"true"`
	StanchionHost       string `envconfig:"DEIS_RIAK_STANCHION_SERVICE_HOST" required:"true"`
	StanchionPort       int    `envconfig:"DEIS_RIAK_STANCHION_SERVICE_PORT" required:"true"`
	AdminKeyLocation    string `envconfig:"ADMIN_KEY_LOCATION" default:"/var/run/secrets/deis/riak-cs/admin-user"`
	AdminSecretLocation string `envconfig:"ADMIN_SECRET_LOCATION" default:"/var/run/secrets/deis/riak-cs/admin-secret"`
}

func getConfig() (*config, error) {
	var ret config
	if err := envconfig.Process(AppName, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
