package main

import(
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
	"os"
)

// main takes all build flags or environment variables, defines
// the configuration and runs the client.
func main() {

	app := cli.NewApp()
	app.Name = "client"
	app.Usage = "instantiate a new client to make requests"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:   "file",
			Aliases: []string{"f"},
			EnvVars: []string{"REQUEST_FILE"},
		},
		&cli.StringFlag{
			Name:   "queue",
			Aliases: []string{"q"},
			Value: "requests",
			EnvVars: []string{"QUEUE"},
		},
		&cli.StringFlag{
			Name:   "connection",
			Aliases: []string{"c"},
			Value: "amqp://guest:guest@localhost:5672/",
			EnvVars: []string{"CONNECTION"},
		},

	}

	app.Action = func(ctx *cli.Context) error {

		service := DefineService(ctx)

		if err := service.Run(); err != nil {
			return fmt.Errorf("could not start service %w", err)
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("error occurred %v", err)
	}
}