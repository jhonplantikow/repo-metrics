package main

import (
	"fmt"
	"repo-metrics/pkg/application/daemons"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/config"
	"gofr.dev/pkg/gofr/logging"
)

func main() {
	app := gofr.NewCMD()

	app.SubCommand("startup", func(c *gofr.Context) (interface{}, error) {
		logger := logging.NewLogger(logging.DEBUG)

		if err := startDaemon(app.Config, logger, "startup"); err != nil {
			return nil, err
		}

		return "daemon started successfully", nil
	},
		gofr.AddDescription("Starts the processing metrics repository"),
		gofr.AddHelp("Use this command to start the processing metrics repository"),
	)

	app.Run()
}

func startDaemon(c config.Config, logger logging.Logger, appName string) error {
	d, err := DaemonFor(c, logger, appName)
	if err != nil {
		logger.Errorf("error creating daemon for %s: %v", appName, err)
		return err
	}

	logger.Infof("starting %s daemon...", appName)

	if err := d.Start(); err != nil {
		logger.Error(err)
	}

	return nil
}

func DaemonFor(c config.Config, logger logging.Logger, appName string) (*daemons.Daemon, error) {
	switch appName {
	case "startup":
		return daemons.NewStartupDaemon(c, logger)
	default:
		return nil, fmt.Errorf("unknown app name: %s", appName)
	}
}
