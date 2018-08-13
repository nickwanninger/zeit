package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	start := time.Now()

	if len(os.Args) == 1 {
		fmt.Println("Usage: zeit <command>")
		os.Exit(-1)
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	cmd.Wait()

	duration := time.Since(start)

	fmt.Println("")
	fmt.Println(duration)

}
