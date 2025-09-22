package utils

import (
	"net"
	"strings"
	"os/exec"
)

func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "127.0.0.1"
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

func GetPublicIP() string {
	out, err := exec.Command("curl", "-4", "ifconfig.me").Output()
	if err != nil {
		return "0.0.0.0"
	}
	return strings.TrimSpace(string(out))
}

