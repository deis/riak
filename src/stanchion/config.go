package stanchion

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	AppName = "deis-riak-stanchion"
)

type config struct {
	ConfFilePath        string `envconfig:"CONF_FILE" required:"true"`
	NumPorts            int    `envconfig:"NUM_PORTS" required:"true"`
	RiakHost            string `envconfig:"DEIS_RIAK_SERVICE_HOST" required:"true"`
	RiakPort            int    `envconfig:"DEIS_RIAK_SERVICE_PORT" required:"true"`
	ListenHost          string `envconfig:"LISTEN_HOST" required:"true"`
	ListenPort          int    `envconfig:"LISTEN_PORT" required:"true"`
	AdminKeyLocation    string `envconfig:"ADMIN_KEY_LOCATION" default:"/var/run/secrets/deis/riak-stanchion/admin/access-key-id"`
	AdminSecretLocation string `envconfig:"ADMIN_SECRET_LOCATION" default:"/var/run/secrets/deis/riak-stanchion/admin/access-secret-key"`
}

func getConfig() (*config, error) {
	var ret config
	if err := envconfig.Process(AppName, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
