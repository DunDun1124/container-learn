package main

import (
	"os"
	"syscall"
)

func main() {
	oldRootF, err := os.Open("/")
	defer oldRootF.Close()
	if err != nil {
		println("fail to open root: %v\n", err)
	}

	// change working dir to old root
	err =
	if err != nil {
		println("fail to chroot %v\n", err)
	}

	err = oldRootF.Chdir()
	if err != nil {
		println("chdir() err: %v", err)
	}
	err = syscall.Chroot(".")
	if err != nil {
		println("chroot back err: %v", err)
	}
}
