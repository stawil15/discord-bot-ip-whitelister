package firewall

import (
	"fmt"
	"strings"

	"github.com/geekloper/discord-bot-ip-whitelister/config"
	"github.com/geekloper/discord-bot-ip-whitelister/logger"
	"github.com/geekloper/discord-bot-ip-whitelister/utils"
)

var defaultServices []string

func InitFirewall(services string) {
	// Check UFW Status
	isUFWInstalled()
	isUFWActive()

	if config.DebugMode() {
		dumpAllUFWRules()
	}

	defaultServices = strings.Split(services, ",")

	// Vérifier que chaque entrée est bien formée (port/protocol)
	for _, service := range defaultServices {
		if !strings.Contains(service, "/") {
			logger.Fatal("Invalid format for service, please use format port/protocol", "SERVICE", service)
		}
	}

	logger.Info("✅ Firewall initialized successfully.", "SERVICES", defaultServices)
}

// AllowUFWRule applies a UFW allow rule for a given IP.
func AllowUFWRule(ip string) error {
	if err := applyUFWRule(ip, "allow"); err != nil {
		logger.Error("Failed to allow UFW rule", "ip", ip, "error", err)
		return err
	}

	logger.Info("Successfully added UFW allow rule", "ip", ip)
	return nil
}

// DenyUFWRule applies a UFW deny rule for a given IP.
func DenyUFWRule(ip string) error {
	if err := applyUFWRule(ip, "deny"); err != nil {
		logger.Error("Failed to deny UFW rule", "ip", ip, "error", err)
		return err
	}

	logger.Info("Successfully added UFW deny rule", "ip", ip)
	return nil
}

// DeleteUFWRule removes an IP from UFW.
func DeleteUFWRule(ip string) error {
	if err := removeUFWRule(ip); err != nil {
		logger.Error("Failed to delete UFW rule", "ip", ip, "error", err)
		return err
	}

	logger.Info("Successfully removed UFW rule", "ip", ip)
	return nil
}

// Helper function to apply UFW allow or deny rules
func applyUFWRule(ip, action string) error {
	for _, service := range defaultServices {
		parts := strings.Split(service, "/")
		if len(parts) != 2 {
			return fmt.Errorf("Invalid format for service: %s", service)
		}
		port, proto := parts[0], parts[1]

		cmd := fmt.Sprintf("sudo ufw %s proto %s from %s to any port %s", action, proto, ip, port)
		output, err := utils.RunCommand(cmd)
		if err != nil {
			return fmt.Errorf("❌ failed to %s rule for port %s: %v\n%s", action, port, err, output)
		}
	}

	return nil
}

// Helper function to remove UFW rules
func removeUFWRule(ip string) error {
	for _, service := range defaultServices {
		parts := strings.Split(service, "/")
		if len(parts) != 2 {
			return fmt.Errorf("Invalid format for service: %s", service)
		}
		port, proto := parts[0], parts[1]

		cmd := fmt.Sprintf("sudo ufw delete proto %s from %s to any port %s", proto, ip, port)
		output, err := utils.RunCommand(cmd)
		if err != nil {
			return fmt.Errorf("❌ failed to delete rule for port %s: %v\n%s", port, err, output)
		}
	}

	return nil
}
