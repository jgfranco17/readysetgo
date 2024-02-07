package main

import (
	"flag"
	"fmt"
)

func main() {
	command := flag.String("command", "help", "Command to run")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Printf("Hello, user!\n")
	} else {
		switch *command {
		case "help":
			fmt.Printf("Welcome to readysetgo!\n")
		default:
			fmt.Printf("Hello, user!\n")
		}
	}
}
