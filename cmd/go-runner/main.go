package main

import (
	"flag"
	"log"

	"github.com/tsovak/go-test-parser/cmd/go-runner/cmd"
)

func main() {
	flag.Parse()
	executor := cmd.NewExecutor()
	err := executor.Execute(flag.CommandLine)
	if err != nil {
		log.Fatal("execution failed: ", err)
	}
}
