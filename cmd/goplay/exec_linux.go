// +build linux

package main

import (
	"os/exec"
)

// http://sources.debian.net/src/xdg-utils/1.1.0~rc1%2Bgit20111210-7.1/scripts/xdg-open/
// http://sources.debian.net/src/xdg-utils/1.1.0~rc1%2Bgit20111210-7.1/scripts/xdg-mime/

func invokeBrowser(input string) *exec.Cmd {
	return exec.Command("xdg-open", input)
}
