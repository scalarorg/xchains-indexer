package main

import (
	"log"

	"github.com/scalarorg/xchains-indexer/cmd"
)

func main() {
	// simplest main as recommended by the Cobra package
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("Failed to execute. Err: %v", err)
	}
}
