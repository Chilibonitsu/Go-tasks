package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	// Установка таймера для отправки сигнала
	//time.Sleep(2 * time.Second)

	parentPID := os.Getppid()
	parentProc, _ := os.FindProcess(parentPID)

	parentProc.Signal(syscall.SIGUSR1)
	fmt.Printf("Дочерний процесс %d отправляет SIGUSR1 процессу %d\n", os.Getpid(), parentPID)

	parentProc.Signal(syscall.SIGUSR2)
	fmt.Printf("Дочерний процесс %d отправляет SIGUSR2 процессу %d\n", os.Getpid(), parentPID)
	//time.Sleep(10 * time.Second)
	//os.Exit(0)
}
