// Package main runner for app
package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/activatedio/cs"
	"github.com/activatedio/cs/sources"
	"github.com/activatedio/cs/sources/yaml"
	"github.com/activatedio/exechc"
	"github.com/alecthomas/kong"
)

const (
	configKeyPrefix = "exechc"
)

// CLI command line arguments
var CLI struct {
	Run struct {
		ConfigPath string `help:"Config path."`
	} `cmd:"" help:"Run server."`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "run":

		err := run(CLI.Run.ConfigPath)

		if err != nil {
			log.Fatal(err)
		}

	default:
		panic(ctx.Command())
	}
}

func run(configPath string) error {

	cfg := cs.NewConfig()

	cfg.AddSource(sources.NewSource(configKeyPrefix, &exechc.Runtime{
		Port: 8080,
		Host: "localhost",
	}))

	if configPath != "" {
		cfg.AddSource(yaml.NewSourceFromPath(configPath, ""))
	}

	cfg.AddLateBindingSource(sources.NewEnvLateBindingSource(strings.ToUpper(configKeyPrefix)))

	res := &exechc.Runtime{}

	err := cfg.Read(configKeyPrefix, res)

	if err != nil {
		return err
	}

	chk := exechc.NewChecker(res)

	svr := exechc.NewServer(res, chk)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Shutting down")
		exechc.Must(svr.Shutdown())
	}()

	return svr.Start()

}
