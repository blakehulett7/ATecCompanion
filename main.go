package main

import (
	"os"
	"os/exec"
)

func main() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
	startPrompt()
}
