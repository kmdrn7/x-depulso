package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	CronSpec string `required:"true" split_words:"true"`
	Interval int32 `required:"true" split_words:"true"`
	ListenInterface string `required:"true" split_words:"true"`
	Promisc bool `required:"true" split_words:"true"`
	WriteLocation string `required:"true" split_words:"true"`
}

func GetConfig() *Config {
	var config Config
	err := envconfig.Process("DEPULSO", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &config
}