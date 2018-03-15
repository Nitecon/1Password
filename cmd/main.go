package main

import (
	"fmt"

	"github.com/Nitecon/1Password/config"
)

func main() {
	fmt.Println("Starting up...")

	config.LoadFromDisk()
}
