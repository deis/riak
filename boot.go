package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/deis/riak/src/cs"
	"github.com/deis/riak/src/riak"
	"github.com/deis/riak/src/stanchion"
)

func main() {
	app := cli.NewApp()
	app.Name = "Deis Riak"
	app.Usage = "Binary to launch and configure all Deis Riak components"
	app.Commands = []cli.Command{
		{
			Name:        "riak",
			Description: "Configures and launches Riak",
			Action:      riak.Action,
		},
		{
			Name:        "riak-cs",
			Description: "Configures and launches Riak CS",
			Action:      cs.Action,
		},
		{
			Name:        "riak-stanchion",
			Description: "Configures and launches Riak Stanchion",
			Action:      stanchion.Action,
		},
	}
	app.Run(os.Args)
}
