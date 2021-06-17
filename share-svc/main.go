package main

import (
	"github.com/al8n/shareable-notes/share-svc/config"
	"github.com/al8n/shareable-notes/share-svc/pkg/server"
	boot "github.com/al8n/micro-boot"
	"log"
)

func main() {
	var (
		cfg = config.GetConfig()
		srv = server.GetServer()
	)

	boot.SetDefaultConfigFileType(".yml")
	boot.SetDefaultConfigFileName("config")

	bt, err := boot.New("share", srv, boot.Root{
		Start:          	&boot.Config{
			Configurator: cfg,
		},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	if err := bt.Execute(); err != nil {
		log.Fatal(err)
		return
	}
}
