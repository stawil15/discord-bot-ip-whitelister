package firewall

import (
	"strings"

	"github.com/geekloper/discord-bot-ip-whitelister/logger"
	"github.com/geekloper/discord-bot-ip-whitelister/utils"
)

// Check if UFW is installed
func isUFWInstalled() {
	_, err := utils.RunCommand("which ufw")
	if err != nil {
		logger.Fatal("❌ UFW is NOT installed. Install it using: sudo apt install ufw")
	}

	logger.Info("✅ UFW is installed.")
}

// Check if UFW is active
func isUFWActive() {
	output, err := utils.RunCommand("sudo ufw status")
	if err != nil {
		logger.Fatal("❌ Can't check UFW Status, error: %v", err)
	}

	if strings.Contains(string(output), "Status: active") {
		logger.Info("✅ UFW is ACTIVE.")
	} else {
		logger.Fatal("❌ UFW is NOT active. Enable it using: sudo ufw enable")
	}
}

func dumpAllUFWRules() {
	logger.Debug("Dump all UFW rules")
	logger.Debug("=================")
	output, err := utils.RunCommand("sudo ufw status verbose")
	if err != nil {
		logger.Fatal("❌ Can't check UFW Status, error: %v", err)
	}
	logger.Debug(output)
	logger.Debug("=================")
}
