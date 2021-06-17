package main

import (
	"github.com/al8n/shareable-notes/apigateway/config"
	"github.com/al8n/shareable-notes/apigateway/internal/server"
	boot "github.com/al8n/micro-boot"
	"log"
)

func main()  {
	var (
		cfg = config.GetConfig()
	)

	boot.SetDefaultConfigFileType(".yml")
	boot.SetDefaultConfigFileName("config")

	bt, err := boot.New("gateway", &server.Server{}, boot.Root{
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