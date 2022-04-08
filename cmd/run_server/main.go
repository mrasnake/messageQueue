package main

import (
	"fmt"
	"github.com/mrasnake/messageQueue/cmd/run_server/transport"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)


// main takes all build flags or environment variables, defines
// the configuration and runs the server.
func main() {

	app := cli.NewApp()
	app.Name = "client"
	app.Usage = "instanciate a new client to make requests"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
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
		&cli.StringFlag{
			Name:   "logs",
			Aliases: []string{"l"},
			Value: fmt.Sprintf("./logfile-%v.log", time.Now().String()),
			EnvVars: []string{"LOG_FILE"},
		},
	}

	app.Action = func(ctx *cli.Context) error {

		service := transport.defineSettings(ctx)

		if err := service.Run(); err != nil {
			return fmt.Errorf("could not start service %w", err)
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("error occurred %v", err)
	}
}