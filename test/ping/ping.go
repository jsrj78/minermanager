package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	// 要执行的命令和参数
	for i := 144; i < 255; i++ {

		args := "10.7.6." + strconv.Itoa(i) // 这里的参数表示发送 4 个 ICMP 包给 google.com

		var cmd *exec.Cmd

		switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("ping", "-n", "3", args)
		case "linux":
			cmd = exec.Command("ping", "-c", "3", args)
		default:
			fmt.Errorf("unsupported operating system")
		}

		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Command failed:", err)

		}

		//fmt.Println("Command output:", string(output))

		successful := strings.Contains(string(output), "Lost = 0") &&
			strings.Contains(string(output), "Received = 3")

		if successful {
			fmt.Println(successful, args)
		} else {
			fmt.Println("===", args)
		}

	}
}
