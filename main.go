package main

import (
	_ "github.com/google/gopacket/layers"
	_ "github.com/joho/godotenv/autoload"
)

const (
	defaultSnapLen = 262144
)

func main() {
	config := GetConfig()
	PrintMessage(config)
	Run(config)
}