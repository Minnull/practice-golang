package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func checkSSTFile(filePath string) bool {
	var cmd *exec.Cmd

	if runtime.GOOS == "darwin" {
		// macOS
		cmd = exec.Command("rocksdb_sst_dump", "--file="+filePath, "--verify_checksum")
	} else if runtime.GOOS == "linux" {
		// Ubuntu
		cmd = exec.Command("sst_dump", "--file="+filePath, "--verify_checksum")
	} else {
		fmt.Println("Unsupported OS")
		return false
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing sst_dump:", err)
		return false
	}

	output := out.String()
	if bytes.Contains([]byte(output), []byte("is not a valid SST file")) {
		return false
	}

	return true
}

func main() {
	filePath := "/Users/Desktop/testsst/1-42674.sst"
	if checkSSTFile(filePath) {
		fmt.Println("Checksum verification passed.")
	} else {
		fmt.Println("Checksum verification failed.")
	}
}
