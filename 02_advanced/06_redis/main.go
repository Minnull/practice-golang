package main

import (
	"flag"
	"log"

	"github.com/Minnull/practice-golang/02_advanced/06_redis/validator"
)

func main() {
	config := validator.Config{}

	flag.StringVar(&config.SourceHost, "sourceHost", "", "Source Redis host")
	flag.StringVar(&config.SourcePort, "sourcePort", "", "Source Redis port")
	flag.StringVar(&config.SourcePassword, "sourcePassword", "", "Source Redis password")
	flag.StringVar(&config.TargetHost, "targetHost", "", "Target Redis host")
	flag.StringVar(&config.TargetPort, "targetPort", "", "Target Redis port")
	flag.StringVar(&config.TargetPassword, "targetPassword", "", "Target Redis password")
	flag.StringVar(&config.ExecutionMode, "mode", "batch", "Execution mode: 'batch' for all commands at once, 'step' for one at a time")

	flag.Parse()

	if config.SourceHost == "" || config.SourcePort == "" || config.TargetHost == "" || config.TargetPort == "" {
		log.Fatal("All host and port parameters must be provided")
	}

	err := validator.RunValidation(config)
	if err != nil {
		log.Fatalf("Validation failed: %v", err)
	}
}
