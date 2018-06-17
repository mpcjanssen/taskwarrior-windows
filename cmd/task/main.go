package main

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	newArgs := make([]string, len(os.Args)+1)
	newArgs[0] = "task"
	for idx, arg := range os.Args[1:] {
		newArgs[idx+1] = strings.Replace(arg, "'", "\\'", -1)
	}
	cmd := exec.Command("bash", "-c", strings.Join(newArgs," "))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()

	status := cmd.ProcessState.Sys().(syscall.WaitStatus)
	os.Exit(status.ExitStatus())

}
