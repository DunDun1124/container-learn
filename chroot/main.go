package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

const ShellToUse = "/bin/bash"

func shellOut(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

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

	out, s, err := shellOut("ls -l")
	if err != nil {
		fmt.Println("Failed to execute command:", err)
	} else {
		fmt.Println("result is:", out+s)
	}

	err = oldRootF.Chdir()
	if err != nil {
		fmt.Printf("chdir() err: %v", err)
	}
	err = syscall.Chroot(".")
	if err != nil {
		fmt.Printf("chroot back err: %v", err)
	}

	out, s, err = shellOut("ls -l")
	if err != nil {
		fmt.Println("Failed to execute command:", err)
		return
	} else {
		fmt.Println("result is:", out+s)
	}

}
