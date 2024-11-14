package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {

	parentPID := os.Getppid()
	parentProc, _ := os.FindProcess(parentPID)

	parentProc.Signal(syscall.SIGXCPU)
	fmt.Printf("Дочерний процесс %d отправляет SIGXCPU процессу %d\n", os.Getpid(), parentPID)

	parentProc.Signal(syscall.SIGTERM)
	fmt.Printf("Дочерний процесс %d отправляет SIGTERM процессу %d\n", os.Getpid(), parentPID)

	//time.Sleep(10 * time.Second)

}
