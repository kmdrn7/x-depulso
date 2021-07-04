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
PCAP File Retention : %s days
CICFlowmeter Path : %s
Kafka Topic : %s
Kafka Host : %s
Kafka Port : %s
Sensor Serial : %s
MLServer Url : %s
=========================================
::::::::::::: ===========================
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
			strconv.Itoa(int(config.DaysRetention)),
			config.CICFlowmeterPath,
			config.KafkaTopic,
			config.KafkaHost,
			strconv.Itoa(int(config.KafkaPort)),
			config.SensorSerial,
			config.MLServerUrl,
		),
	)
}
