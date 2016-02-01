package stanchion

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	AppName = "deis-riak-stanchion"
)

type config struct {
	ListenHost          string `envconfig:"LISTEN_HOST" required:"true"`
	ListenPort          int    `envconfig:"LISTEN_PORT" required:"true"`
	StanchionURL        string `envconfig:"STANCHION_URL" required:"true"`
	AdminKeyLocation    string `envconfig:"ADMIN_KEY_LOCATION" default:"/var/run/secrets/deis/riak-stanchion/admin-user"`
	AdminSecretLocation string `envconfig:"ADMIN_SECRET_LOCATION" default:"/var/run/secrets/deis/riak-stanchion/admin-secret"`
}

func getConfig() (*config, error) {
	var ret config
	if err := envconfig.Process(AppName, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
