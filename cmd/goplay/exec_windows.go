// +build windows

package open

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func invokeBrowser(input string) *exec.Cmd {
	rundll32Path = filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe")
	cmd := exec.Command(rundll32Path, "url.dll,FileProtocolHandler", input)
	return cmd
}
