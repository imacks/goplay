// +build !darwin,!windows,!linux

package main

import (
	"os/exec"
)

func invokeBrowser(input string) *exec.Cmd {
	panic("cannot invoke browser on this platform")
}
