package server 

import (
	"os"
	"os/exec"
	"fmt"
)

var cmd *exec.Cmd

func startTunnel() {
    cmd = exec.Command("ssh", "-N", "-R", "127.0.0.1:8000:localhost:8000", "ookami@192.168.1.42")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    go func() {
        if err := cmd.Run(); err != nil {
            fmt.Println("Tunnel error:", err)
        }
    }()
}

// func stopTunnel() {
//     if cmd != nil && cmd.Process != nil {
//         cmd.Process.Kill()
//     }
// }
