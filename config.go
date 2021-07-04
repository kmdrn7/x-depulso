package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	CronSpec string `required:"true" split_words:"true" default:"0 * * * * *"`
	CICFlowmeterPath string `required:"true" envconfig:"CICFLOWMETER_PATH" default:"/app/CICFlowmeter.jar"`
	DaysRetention int32 `required:"true" split_words:"true" default:"2"`
	Interval int32 `required:"true" split_words:"true" default:"60"`
	KafkaTopic string `required:"true" split_words:"true"`
	KafkaHost string `required:"true" split_words:"true"`
	KafkaPort int32 `required:"true" split_words:"true"`
	ListenInterface string `required:"true" split_words:"true"`
	MLServerUrl string `required:"true" envconfig:"MLSERVER_URL"`
	Promisc bool `required:"true" split_words:"true" default:"true"`
	SensorSerial string `required:"true" split_words:"true"`
	WriteLocation string `required:"true" split_words:"true" default:"/data/"`
	WriteCsvLocation string `required:"true" split_words:"true" default:"/data/csv/"`
}

func GetConfig() *Config {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &config
}