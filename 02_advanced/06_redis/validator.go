package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	SourceHost     string
	SourcePort     string
	SourcePassword string
	TargetHost     string
	TargetPort     string
	TargetPassword string
	ExecutionMode  string // 控制执行模式
}

type Command struct {
	WriteCommand  string
	VerifyCommand string
}

// RunValidation 运行验证流程
func RunValidation(config Config, commands []Command) error {
	sourceClient := createRedisClient(config.SourceHost, config.SourcePort, config.SourcePassword)
	targetClient := createRedisClient(config.TargetHost, config.TargetPort, config.TargetPassword)

	ctx := context.Background()

	for i, cmd := range commands {
		log.Printf("Executing command %d, cmd:[%s]...", i+1, cmd.WriteCommand)

		// 执行写入命令到源集群
		if _, err := ExecuteCommand(ctx, sourceClient, cmd.WriteCommand); err != nil {
			return fmt.Errorf("failed to write data to source: %v", err)
		}

		// 验证源集群中的数据
		log.Printf("source redis check cmd: [%s]", cmd.VerifyCommand)
		sourceResult, err := ExecuteCommand(ctx, sourceClient, cmd.VerifyCommand)
		if err != nil {
			return fmt.Errorf("source data verification failed: %v", err)
		}

		// 验证目标集群中的数据
		log.Printf("target redis check cmd: [%s]", cmd.VerifyCommand)
		targetResult, err := ExecuteCommand(ctx, targetClient, cmd.VerifyCommand)
		if err != nil {
			return fmt.Errorf("target data verification failed: %v", err)
		}

		// 对比源集群和目标集群的数据
		if sourceResult != targetResult {
			return fmt.Errorf("data mismatch for command %s: sourceResult='%s', targetResult='%s'", cmd.VerifyCommand, sourceResult, targetResult)
		} else {
			log.Printf("Command %d verified successfully.command %s: sourceResult='%s', targetResult='%s'", i+1, cmd.VerifyCommand, sourceResult, targetResult)
		}

		// 根据执行模式判断是否停顿
		if config.ExecutionMode == "step" {
			log.Println("Press Enter to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}

	return nil
}

// convertToInterfaceSlice 将字符串 slice 转换为 interface{} slice
func convertToInterfaceSlice(args []string) []interface{} {
	result := make([]interface{}, len(args))
	for i, v := range args {
		result[i] = v
	}
	return result
}

func ExecuteCommand(ctx context.Context, client *redis.Client, commands string) (string, error) {
	cmdList := strings.Split(commands, ";")
	var cmdResult bytes.Buffer

	for i, command := range cmdList {
		command = strings.TrimSpace(command)
		if command == "" {
			continue
		}

		parts := strings.Fields(command) // 使用 Fields 分割命令
		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]
		args := parts[1:]

		redisCmd := client.Do(ctx, append([]interface{}{cmd}, convertToInterfaceSlice(args)...)...)
		cmdResultTemp, err := redisCmd.Result()
		if err != nil && err != redis.Nil {
			return "", fmt.Errorf("command failed: %s, error: %v", command, err)
		}

		fmt.Printf("Command: %s\nResult: %v\n", command, cmdResultTemp)

		if i > 0 {
			cmdResult.WriteString("; ")
		}
		cmdResult.WriteString(fmt.Sprintf("%v", cmdResultTemp))
	}

	return cmdResult.String(), nil
}

func createRedisClient(host, port, password string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
}
