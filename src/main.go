package main

import (
	"log"
	"lws-client/src/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Printf("Command failed with: %s\n", err.Error())
		return
	}
}
