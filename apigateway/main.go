package main

import (
	"github.com/ALiuGuanyan/margin/apigateway/config"
	"github.com/ALiuGuanyan/margin/apigateway/internal/server"
	boot "github.com/ALiuGuanyan/micro-boot"
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