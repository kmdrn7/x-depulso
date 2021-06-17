package main

import (
	"fmt"
	"time"
)

func WriteLog(message string){
	fmt.Print(time.Now(), " => ")
	fmt.Print(message)
	fmt.Println()
}
