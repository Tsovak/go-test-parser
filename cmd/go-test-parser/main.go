package main

import (
	"log"

	"github.com/tsovak/go-test-parser/cmd/go-test-parser/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal("execution failed:", err)
	}
}
