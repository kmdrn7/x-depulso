package main

import (
	"encoding/json"
	"github.com/kelseyhightower/envconfig"
	"io/ioutil"
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

type SensorResponse struct {
	Status int32 `json:"status"`
	Data struct {
		Config string `json:"config"`
	}
}

type ConfigResponse struct {
	KafkaTopic string `json:"KAFKA_TOPIC"`
	KafkaHost string `json:"KAFKA_HOST"`
	KafkaPort int32 `json:"KAFKA_PORT"`
	ListenInterface string `json:"LISTEN_INTERFACE"`
}

func UpdateConfigFromServer(config *Config) {
	// Fetch sensor's configuration data from server
	res, errRes := ApiClient.Get(config.MLServerUrl+"/api/v1/sensor/"+config.SensorSerial+"/config", nil)
	if errRes != nil {
		panic(errRes)
	}
	b, errBody := ioutil.ReadAll(res.Body)
	if errBody != nil {
		panic(errBody)
	}

	// Unmarshal response body
	var body SensorResponse
	err = json.Unmarshal(b, &body)
	if err != nil {
		panic(err)
		return
	}

	// Unmarshal configuration parameters
	var configResponse ConfigResponse
	err = json.Unmarshal([]byte(body.Data.Config), &configResponse)
	if err != nil {
		panic(err)
		return
	}

	// Update local config with data fetched from server
	config.KafkaHost = configResponse.KafkaHost
	config.KafkaPort = configResponse.KafkaPort
	config.KafkaTopic = configResponse.KafkaTopic
	config.ListenInterface = configResponse.ListenInterface
}