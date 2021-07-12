package main

import (
	"github.com/gojek/heimdall/v7/httpclient"
	_ "github.com/google/gopacket/layers"
	_ "github.com/joho/godotenv/autoload"
	"time"
)

const (
	defaultSnapLen = 262144
)

var ApiClient = httpclient.NewClient(
	httpclient.WithHTTPTimeout(10000 * time.Millisecond),
)

func main() {
	WriteLog("Application is booting...")
	config := GetConfig()
	WriteLog("Fetching config from server...")
	UpdateConfigFromServer(config)
	WriteLog("Config fetched from server...")
	PrintMessage(config)
	Run(config)
}