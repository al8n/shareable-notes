package config

import (
	bootapi "github.com/ALiuGuanyan/micro-boot/api"
)

type ShareService struct {
	Name string `json:"name" yaml:"name"`
	APIs bootapi.APIs `json:"apis" yaml:"apis"`
}
