package config

import (
	bootapi "github.com/ALiuGuanyan/micro-boot/api"
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
)

type Share struct {
	Name string `json:"name" yaml:"name"`
	APIs bootapi.APIs `json:"apis" yaml:"apis"`
}


func (s *Share) BindFlags(fs *bootflag.FlagSet)  {
	fs.StringVar(&s.Name, "name", "sharesvc", "specify the micro service name")
}

func (s *Share) Parse() (err error) { return nil }

