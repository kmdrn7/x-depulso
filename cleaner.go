package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func RunCleaner(day time.Duration, cleanPath string) {
	var files []os.FileInfo

	tmpFiles, err := ioutil.ReadDir(cleanPath)
	if err != nil {
		return
	}

	for _, file := range tmpFiles {
		if file.Mode().IsRegular() {
			if time.Now().Sub(file.ModTime()) > 24*day*time.Hour {
				files = append(files, file)
			}
		}
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".pcap") {
			fileFullPath := cleanPath + file.Name()
			err := os.Remove(fileFullPath)
			if err != nil {
				fmt.Println("Error deleting file. ", err)
				return
			}
			WriteLog("Deleting file " + fileFullPath)
		}
	}
}