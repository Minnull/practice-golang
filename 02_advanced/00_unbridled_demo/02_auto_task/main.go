package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Task struct {
	Name                    string
	Action                  func() error
	Executed                bool
	Status                  string
	AwaitingInputOnFirstRun bool
}

func main() {
	tasks := []Task{
		{"1. 启动准备工作", startTask, false, "", false},
		{"2. 监控任务", monitorTask, false, "", true},
		{"3. 飞机起飞", airTask, false, "", false},
		{"4. 飞机落地", air2Task, false, "", false},
		{"5. 完成飞行", over, false, "", false},
	}

	maxTaskLength := getMaxLength(tasks)
	statusWidth := 10
	totalWidth := maxTaskLength + statusWidth + 16

	clearScreen()
	displayMainMenu(totalWidth)
	for i := range tasks {
		tasks[i].Status = "等待中"
	}
	updateUI(tasks, maxTaskLength, statusWidth, totalWidth, true, false)

	for i := range tasks {
		updateUI(tasks, maxTaskLength, statusWidth, totalWidth, false, tasks[i].AwaitingInputOnFirstRun)
		if tasks[i].AwaitingInputOnFirstRun && !confirmTaskExecution() {
			fmt.Println("程序已终止。")
			return
		}

		tasks[i].Status = "执行中"
		updateUI(tasks, maxTaskLength, statusWidth, totalWidth, false, false)

		err := tasks[i].Action()
		if err != nil {
			tasks[i].Status = "失败"
		} else {
			tasks[i].Status = "成功"
			tasks[i].Executed = true
		}
		updateUI(tasks, maxTaskLength, statusWidth, totalWidth, false, false)

		time.Sleep(2 * time.Second)
	}

	fmt.Println("\n所有步骤执行完毕。")
}

func getMaxLength(tasks []Task) int {
	maxLength := 0
	for _, task := range tasks {
		if len(task.Name) > maxLength {
			maxLength = len(task.Name)
		}
	}
	return maxLength
}

func displayMainMenu(totalWidth int) {
	border := "═"
	fmt.Printf("╔%s╗\n", strings.Repeat(border, totalWidth))
	fmt.Printf(" %s \n", padCenter("auto 工作流", totalWidth-2))
	fmt.Printf("╚%s╝\n", strings.Repeat(border, totalWidth))
}

func padCenter(text string, width int) string {
	padding := width - len(text)
	if padding <= 0 {
		return text
	}
	leftPadding := padding / 2
	rightPadding := padding - leftPadding
	return fmt.Sprintf("%s%s%s", strings.Repeat(" ", leftPadding), text, strings.Repeat(" ", rightPadding))
}

func getTaskStatusIcon(executed bool, status string) string {
	if executed && status == "成功" {
		return "✔"
	} else if status == "执行中" {
		return "🔄"
	} else if status == "失败" {
		return "✖"
	} else {
		return " "
	}
}

func updateUI(tasks []Task, maxLength, statusWidth, totalWidth int, firstUpdate, awaitingInput bool) {
	if firstUpdate {
		clearScreen()
		displayMainMenu(totalWidth)
	}

	for i, task := range tasks {
		moveCursorToLine(i + 4)
		fmt.Printf("[%s] %-*s %-*s\n", getTaskStatusIcon(task.Executed, task.Status), statusWidth-2, task.Status, maxLength, task.Name)
	}

	if awaitingInput {
		moveCursorToEnd(len(tasks) + 6) // Move cursor below tasks
		fmt.Print("按 Enter 键继续执行下一步，或输入其他内容终止程序：")
	}
}

func moveCursorToLine(n int) {
	fmt.Printf("\033[%d;0H", n)
}

func moveCursorToEnd(n int) {
	fmt.Printf("\033[%d;0H", n)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func confirmTaskExecution() bool {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input == ""
}

func startTask() error {
	time.Sleep(2 * time.Second)
	return nil
}

func monitorTask() error {
	time.Sleep(2 * time.Second)
	return nil
}

func airTask() error {
	time.Sleep(3 * time.Second)
	fmt.Println("第一次日志")
	return nil
}

func air2Task() error {
	time.Sleep(1 * time.Second)
	fmt.Println("测试日志更新输出")
	return nil
}

func over() error {
	time.Sleep(2 * time.Second)
	fmt.Println("完成任务")
	return nil
}
