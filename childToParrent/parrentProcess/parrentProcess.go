package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Обработчики сигналов
func handleSIGINT() {
	fmt.Println()
	fmt.Println("Получен сигнал SIGINT")
}

func handleSIGUSR1() {
	fmt.Println("Получен сигнал SIGUSR1")
}

func handleSIGUSR2() {
	fmt.Println("Получен сигнал SIGUSR2")
}

func main() {
	timer := time.NewTimer(5 * time.Second)
	sigChan := make(chan os.Signal, 5)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGXCPU, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		runChild(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		runChild2(ctx)

	}()

	wg.Wait()
	//fmt.Println(len(sigChan))
	defer close(sigChan)
	for {
		select {
		case sig := <-sigChan:
			switch sig {
			case syscall.SIGUSR1:
				handleSIGUSR1()
			case syscall.SIGUSR2:
				handleSIGUSR2()
			case syscall.SIGINT:
				fmt.Println("Получен сигнал SIGINT")
				fmt.Println("Завершаем работу")
				return
			case syscall.SIGXCPU:
				fmt.Println("Получен сигнал SIGXCPU")
			case syscall.SIGTERM:
				fmt.Println("Получен сигнал SIGTERM")
			}
		case <-timer.C:
			fmt.Println("Выходим из программы по таймеру")
			return
		}
	}

}
func runChild(ctx context.Context) {

	child := exec.CommandContext(ctx, "./child")
	child.Dir = "../child"
	child.Stdout = os.Stdout
	child.Stderr = os.Stderr

	err := child.Start()
	if err != nil {
		fmt.Println("Ошибка при запуске дочернего процесса:", err)
		return
	}

	child.Wait()
}
func runChild2(ctx context.Context) {

	child := exec.CommandContext(ctx, "./child2")
	child.Dir = "../child2"
	fmt.Println(os.Getwd())
	child.Stdout = os.Stdout
	child.Stderr = os.Stderr
	fmt.Println("Запуск дочернего процесса 2")
	err := child.Start()
	if err != nil {
		fmt.Println("Ошибка при запуске дочернего процесса:", err)
		return
	}

	child.Wait()
}
