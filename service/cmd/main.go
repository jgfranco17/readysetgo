package main

import (
	"flag"
	"fmt"
	"strings"

	logging "github.com/jgfranco17/readysetgo/core/pkg/logging"
)

func main() {
	command := flag.String("command", "help", "Command to run")
	logLevel := flag.String("log", "info", "Logging level for CLI")
	flag.Parse()
	if _, ok := logging.LOG_LEVELS[*logLevel]; ok {
		fmt.Printf("Failed to set logging level: invalid log level '%s' provided\n", *logLevel)
		return
	}
	log := logging.CreateLogger(strings.ToUpper(*logLevel))
	if flag.NArg() == 0 {
		fmt.Printf("Hello, user!\n")
	} else {
		switch *command {
		case "help":
			fmt.Printf("Welcome to readysetgo!\n")
		default:
			log.Debugf("No command specified...")
		}
	}
}
