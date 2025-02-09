package utils

import (
	"net"
	"os/exec"
	"strings"
)

// RunCommand safely executes a command and returns its output.
func RunCommand(cmd string) (string, error) {
	args := strings.Fields(cmd) // Split command string into separate arguments
	out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
	return string(out), err
}

// validateIP checks if the given IP address is valid (IPv4 or IPv6).
func ValidateIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
