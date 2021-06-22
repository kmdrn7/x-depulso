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
	javaExecutable, errLookPath := exec.LookPath("java")
	if errLookPath != nil {
		fmt.Println("Cannot find Java executable in your system. Try install Java first")
		return errors.New("error processing PCAP")
	}

	if _, err := os.Stat(config.CICFlowmeterPath); os.IsNotExist(err) {
		fmt.Println("Cannot find CICFlowmeter.jar in your system.")
		return errors.New("error processing PCAP")
	}

	cmd := &exec.Cmd{
		Path:   javaExecutable,
		Args:   []string{javaExecutable, "-jar", config.CICFlowmeterPath, source},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	// set CFM Environment Variables
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
