package cs

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	AppName = "deis-riak-cs"
)

type config struct {
	NumPorts            int    `envconfig:"NUM_PORTS" required:"true"`
	ListenHost          string `envconfig:"LISTEN_HOST" required:"true"`
	ListenPort          int    `envconfig:"LISTEN_PORT" required:"true"`
	RiakHost            string `envconfig:"RIAK_HOST" default:"localhost"`
	RiakProtobufPort    int    `envconfig:"RIAK_PROTOBUF_PORT" default:"8098"`
	StanchionHost       string `envconfig:"DEIS_RIAK_STANCHION_SERVICE_HOST" required:"true"`
	StanchionPort       int    `envconfig:"DEIS_RIAK_STANCHION_SERVICE_PORT" required:"true"`
	AdminKeyLocation    string `envconfig:"ADMIN_KEY_LOCATION" default:"/var/run/secrets/deis/riak-cs/admin/access-key-id"`
	AdminSecretLocation string `envconfig:"ADMIN_SECRET_LOCATION" default:"/var/run/secrets/deis/riak-cs/admin/access-secret-key"`
}

func getConfig() (*config, error) {
	var ret config
	if err := envconfig.Process(AppName, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
