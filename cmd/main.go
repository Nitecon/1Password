package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

type ConfigData struct {
	VaultLocations []string
}

func main() {
	fmt.Println("Starting up...")
	cu, err := user.Current()
	if err != nil {
		log.Printf("Must be a valid user account in order to use this tool: %s\n", err)
		os.Exit(1)
	}
	confPath := fmt.Sprintf("%s%s%s", cu.HomeDir, string(os.PathSeparator), ".1Password-config.json")
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		log.Println("No config file found creating initial setup...")
		ioutil.WriteFile
	}
}
