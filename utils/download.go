package utils

import (
	"os/exec"
	"runtime"
)

func DownloadFromLink(url string) {
	var cmd *exec.Cmd

	if runtime.GOOS == "linux" {
		cmd = exec.Command("xdg-open", url)
	}
	if runtime.GOOS == "darwin" {
		cmd = exec.Command("open", url)
	}
	if runtime.GOOS == "windows" {
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	}

	if cmd != nil {
		cmd.Start()
	}
}