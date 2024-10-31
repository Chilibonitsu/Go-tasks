package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Process struct {
	PID  int
	VSZ  int
	Name string
}

func main() {
	cmd := exec.Command("ps", "-e", "-o", "pid,rss,comm=", "--sort=-rss")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	processes := []Process{}

	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		pid, _ := strconv.Atoi(fields[0])
		vsz, _ := strconv.Atoi(fields[1])
		name := strings.Join(fields[2:], " ")
		processes = append(processes, Process{PID: pid, VSZ: vsz, Name: name})
	}
	fmt.Println("NAME PID VSZ")
	for i := 0; i < 8; i++ {
		fmt.Println(processes[i].Name, processes[i].PID, processes[i].VSZ)
	}
	fmt.Println()
	fmt.Printf("Process \"%s\" with PID %d consumes the most memory ", processes[0].Name, processes[0].PID)
	fmt.Println()
	fmt.Printf("Kill process %s with PID %d? y/n", processes[0].Name, processes[0].PID)
	fmt.Println()
	var input string
	fmt.Scanln(&input)
	if input == "y" || input == "Y" {
		//
		process, err := os.FindProcess(int(processes[0].PID))
		if err != nil {
			log.Fatal(err)
		}
		errKill := process.Kill()

		if errKill != nil {
			fmt.Println(errKill)
		}
		//

		// //
		// if err := syscall.Kill(int(processes[0].PID), syscall.SIGKILL); err != nil {
		// 	fmt.Print(err)
		// }
		// //

		// //
		// cmd = exec.Command("sudo", "kill", "-9", strconv.Itoa(processes[0].PID))
		// //cmd = exec.Command("pkill", "firefox")
		// err = cmd.Run()
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// //
	}

}
