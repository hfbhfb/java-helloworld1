package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// 定义要运行的命令
	// cmd := exec.Command("sh", "-c", "pwd")
	cmd := exec.Command("sh", "buildrunjava.sh")

	// 获取命令的输出
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// 打印输出
	fmt.Println(string(output))
}
