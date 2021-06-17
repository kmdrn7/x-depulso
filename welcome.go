package main

import (
	"fmt"
	"strconv"
)

func PrintMessage(config *Config){
	message := `
	____  __________  __  ____   _____ ____ 
   / __ \/ ____/ __ \/ / / / /  / ___// __ \
  / / / / __/ / /_/ / / / / /   \__ \/ / / /
 / /_/ / /___/ ____/ /_/ / /______/ / /_/ / 
/_____/_____/_/    \____/_____/____/\____/

Running with configurations :::::::::::::
=========================================
Listen Interface : %s
Processing Interval : %s
Cron Spec : %s
Promiscuous Mode : %s 
PCAP Write Location : %s
CSV Write Location : %s
Kafka Topic : %s
Kafka Host : %s
Kafka Port : %s
=========================================
	`
	fmt.Println(
		fmt.Sprintf(
			message,
			config.ListenInterface,
			strconv.Itoa(int(config.Interval)),
			config.CronSpec,
			strconv.FormatBool(bool(config.Promisc)),
			config.WriteLocation,
			config.WriteCsvLocation,
			config.KafkaTopic,
			config.KafkaHost,
			strconv.Itoa(int(config.KafkaPort)),
		),
	)
}
