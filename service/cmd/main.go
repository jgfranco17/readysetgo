package main

import (
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	command := flag.String("command", "help", "Command to run")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Printf("Hello, user!\n")
	} else {
		switch *command {
		case "help":
			fmt.Printf("Welcome to readysetgo!\n")
		default:
			logger.Debugf("No command specified...")
		}
	}
}
