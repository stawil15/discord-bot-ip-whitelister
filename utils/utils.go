package utils

import (
	"net"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
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

// OptionsToMap converts a slice of options into a map for easier lookup
func OptionsToMap(options []*discordgo.ApplicationCommandInteractionDataOption) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}
	return optionMap
}
