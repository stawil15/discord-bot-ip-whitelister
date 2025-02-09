package firewall

import (
	"fmt"
	"log"
	"log/slog"
	"strings"

	"github.com/geekloper/discord-bot-ip-whitelister/database"
	"github.com/geekloper/discord-bot-ip-whitelister/errors"
	"github.com/geekloper/discord-bot-ip-whitelister/utils"
)

var (
	defaultPorts     []string
	defaultProtocols []string
)

func InitFirewall(servicePorts, serviceProtocols string) {
	// Check UFW Status
	IsUFWInstalled()
	IsUFWActive()

	defaultPorts = strings.Split(servicePorts, ",")
	defaultProtocols = strings.Split(serviceProtocols, ",")

	if len(defaultPorts) != len(defaultProtocols) {
		log.Fatal("Mismatch between number of SERVICE_PORTS and SERVICE_PROTOCOLS")
	}
}

// Allow IP in UFW and update database
func AllowUFWRule(ip, discordUser string) error {
	if !utils.ValidateIP(ip) {
		return errors.ErrInvalidIpFormat
	}

	if ok, dbIp := database.UserIpExists(discordUser); ok {
		DeleteUFWRule(dbIp)
	}

	for i := 0; i < len(defaultPorts); i++ {
		port := defaultPorts[i]
		proto := defaultProtocols[i]

		cmd := fmt.Sprintf("sudo ufw allow proto %s from %s to any port %s", proto, ip, port)
		output, err := utils.RunCommand(cmd)
		if err != nil {
			return fmt.Errorf("failed to allow rule for port %s: %v\n%s", port, err, output)
		}
	}
	err := database.AddRule(ip, discordUser, "allow")
	if err != nil {
		return fmt.Errorf("failed to store rule in database: %v", err)
	}

	slog.Info("IP rule added successfully", "ip", ip)
	return nil
}

// Deny IP in UFW and update database
func DenyUFWRule(ip, discordUser string) error {
	if !utils.ValidateIP(ip) {
		return errors.ErrInvalidIpFormat
	}

	if ok, dbIp := database.UserIpExists(discordUser); ok {
		DeleteUFWRule(dbIp)
	}

	for _, port := range defaultPorts {
		cmd := fmt.Sprintf("sudo ufw deny from %s to any port %s", ip, port)
		output, err := utils.RunCommand(cmd)
		if err != nil {
			return fmt.Errorf("failed to deny rule for port %s: %v\n%s", port, err, output)
		}
	}
	err := database.AddRule(ip, discordUser, "deny")
	if err != nil {
		return fmt.Errorf("failed to store rule in database: %v", err)
	}
	slog.Info("IP rule denied successfully", "ip", ip)
	return nil
}

// Delete UFW rule and update database
func DeleteUFWRule(ip string) error {
	if !utils.ValidateIP(ip) {
		return errors.ErrInvalidIpFormat
	}
	for i := 0; i < len(defaultPorts); i++ {
		port := defaultPorts[i]
		proto := defaultProtocols[i]

		cmd := fmt.Sprintf("sudo ufw delete allow proto %s from %s to any port %s", proto, ip, port)
		output, err := utils.RunCommand(cmd)
		if err != nil {
			return fmt.Errorf("failed to delete rule for port %s (%s): %v\n%s", port, proto, err, output)
		}
	}
	err := database.RemoveRule(ip)
	if err != nil {
		return fmt.Errorf("failed to remove rule from database: %v", err)
	}
	slog.Info("IP rule deleted successfully", "ip", ip)
	return nil
}
