package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

func getCSVPath(sourcePath string, destPath string) string {
	var name string
	_, fileName := path.Split(sourcePath)
	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		name = fileName[:pos]
	}
	return destPath + name + "_Flow.csv"
}

func RunFlowmeter(source string, config *Config) error {
	cfmExecutable, errLookPath := exec.LookPath("CICFlowMeter")
	if errLookPath != nil {
		fmt.Println("Cannot find CICFlowMeter executable in your system. Try add cicflowmeter-cli package in your $PATH")
		return errors.New("error processing PCAP")
	}

	cmd := &exec.Cmd{
		Path:   cfmExecutable,
		Args:   []string{cfmExecutable, source},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	// set CFM Environment Variables
	cmd.Env = append(cmd.Env, "DEFAULT_JVM_OPTS=\"\"-Djava.library.path=$APP_HOME/lib/native\"\"")
	cmd.Env = append(cmd.Env, "KAFKA_TOPIC="+config.KafkaTopic)
	cmd.Env = append(cmd.Env, "KAFKA_HOST="+config.KafkaHost)
	cmd.Env = append(cmd.Env, "KAFKA_PORT="+strconv.Itoa(int(config.KafkaPort)))

	errStart := cmd.Start()
	if errStart != nil {
		return errStart
	}

	errWait := cmd.Wait()
	if errWait != nil {
		return errWait
	}

	return nil
}
