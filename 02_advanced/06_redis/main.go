package main

import (
	"flag"
	"log"
)

func main() {
	config := Config{}

	flag.StringVar(&config.SourceHost, "sourceHost", "127.0.0.1", "Source Redis host")
	flag.StringVar(&config.SourcePort, "sourcePort", "6379", "Source Redis port")
	flag.StringVar(&config.SourcePassword, "sourcePassword", "", "Source Redis password")
	flag.StringVar(&config.TargetHost, "targetHost", "127.0.0.1", "Target Redis host")
	flag.StringVar(&config.TargetPort, "targetPort", "6380", "Target Redis port")
	flag.StringVar(&config.TargetPassword, "targetPassword", "", "Target Redis password")
	flag.StringVar(&config.ExecutionMode, "mode", "batch", "Execution mode: 'batch' for all commands at once, 'step' for one at a time")

	flag.Parse()

	if config.SourceHost == "" || config.SourcePort == "" || config.TargetHost == "" || config.TargetPort == "" {
		log.Fatal("All host and port parameters must be provided")
	}

	commands := initCommands()

	err := RunValidation(config, commands)
	if err != nil {
		log.Fatalf("Validation failed: %v", err)
	}
}

func initCommands() []Command {
	return []Command{
		{
			// 一组命令，支持多条指令
			WriteCommand: "set k1 a; set k1 b",
			// 一组命令，支持多条指令
			VerifyCommand: "GET k1; type k1; exists k1",
		},
		{
			WriteCommand:  "EXPIRE k1 100",
			VerifyCommand: "TTL keyToExpire; get k1",
		},
	}
}
