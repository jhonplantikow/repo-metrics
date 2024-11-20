package daemons

import (
	"os"
	"path/filepath"
	"repo-metrics/pkg/application/adapters/file"
	"repo-metrics/pkg/application/handlers"
	"repo-metrics/pkg/application/services"

	"gofr.dev/pkg/gofr/config"
	"gofr.dev/pkg/gofr/logging"
)

type Daemon struct {
	config config.Config
	logger logging.Logger
}

func NewStartupDaemon(c config.Config, logger logging.Logger) (*Daemon, error) {
	return &Daemon{config: c, logger: logger}, nil
}

func (d *Daemon) Start() error {
	d.logger.Info("starting startup daemon...")

	//make sure if dir exist
	err := os.MkdirAll(filepath.Dir(d.config.Get("IN_REPOS")), 0755)
	if err != nil {
		return err
	}

	fin, err := os.OpenFile(d.config.Get("IN_REPOS"), os.O_RDONLY|os.O_CREATE, 0444)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(d.config.Get("OUT_REJECTED")), 0755)
	if err != nil {
		return err
	}

	fReject, err := os.OpenFile(d.config.Get("OUT_REJECTED"), os.O_WRONLY|os.O_CREATE, 0222)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(d.config.Get("OUT_REPOS")), 0755)
	if err != nil {
		return err
	}

	fout, err := os.OpenFile(d.config.Get("OUT_REPOS"), os.O_WRONLY|os.O_CREATE, 0222)
	if err != nil {
		return err
	}

	//commits
	fReader := file.NewFileReader(fin, fReject)
	cService := services.NewCommitService(fReader)
	wService := services.NewWriterService(fout)
	cHandler := handlers.NewCommitHandler(cService, wService)

	if err := cHandler.ActivityScoreCLI(); err != nil {
		return err
	}

	return nil
}
