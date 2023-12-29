package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// 注意,使用goland时，需要在 setting -> go -> build tag 切换 os 为 linux
func main() {
	oldRootF, err := os.Open("/")
	defer oldRootF.Close()
	if err != nil {
		fmt.Printf("fail to open root: %v\n", err)
	}

	err = syscall.Chroot("/root/new-root")
	if err != nil {
		fmt.Printf("fail to chroot %v\n", err)
	}

	cmd := exec.Command("/bin/ls", "-l")
	_, err = cmd.Output()
	if err != nil {
		fmt.Println("Failed to execute command:", err)
		//return
	}

	err = oldRootF.Chdir()
	if err != nil {
		fmt.Printf("chdir() err: %v", err)
	}
	err = syscall.Chroot("/")
	if err != nil {
		fmt.Printf("chroot back err: %v", err)
	}

	cmd = exec.Command("/bin/ls", "-l")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to execute command:", err)
		return
	}

	fmt.Println(string(output))
}
