package main

import (
	"fmt"
)

func RunProcessor(config *Config, fullPath string) {
	WriteLog("Start processing and streaming " + fullPath)
	err := RunFlowmeter(fullPath, config)
	if err != nil {
		fmt.Println("Error processing PCAP")
		return
	}
	WriteLog("Finished processing and streaming " + fullPath)
}
