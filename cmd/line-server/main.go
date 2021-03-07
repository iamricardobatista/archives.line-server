package main

import (
	"log"
	"os"

	"github.com/nihildacta/line-server/internal/http"
	"github.com/nihildacta/line-server/pkg/file"
	"github.com/urfave/cli/v2"
)

// main entry point of the program
// sets expected parameters and help description
func main() {
	app := &cli.App{
		Name:  "line-server",
		Usage: "Serves lines from a file via an http interface",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file_path",
				Usage:   "path to the file to serve",
				Aliases: []string{"fp"},
				EnvVars: []string{"FILE_PATH"},
			},
			&cli.IntFlag{
				Name:    "server_port",
				Usage:   "Server port",
				Value:   8889,
				Aliases: []string{"sp"},
				EnvVars: []string{"SERVER_PORT"},
			},
		},
		Action: action,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// action business logic for the server
// reads the lines index into memory
// servers the file lines using an http server
func action(c *cli.Context) error {
	file, err := file.ReadFileLines(c.String("file_path"))
	if err != nil {
		return err
	}

	server := http.New(file)
	return server.ListenAndServe(c.Int("server_port"))
}
