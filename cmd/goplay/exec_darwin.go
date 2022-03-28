// +build darwin

package main

import (
	"os/exec"
)

func invokeBrowser(input string) *exec.Cmd {
	return exec.Command("open", input)
}
