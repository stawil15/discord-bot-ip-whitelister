package firewall

import (
	"log"
	"log/slog"
	"strings"

	"github.com/geekloper/discord-bot-ip-whitelister/utils"
)

// Check if UFW is installed
func IsUFWInstalled() {
	_, err := utils.RunCommand("which ufw")
	if err != nil {
		log.Fatal("❌ UFW is NOT installed. Install it using: sudo apt install ufw")
	}

	slog.Info("✅ UFW is installed.")
}

// Check if UFW is active
func IsUFWActive() {
	output, err := utils.RunCommand("sudo ufw status")
	if err != nil {
		log.Fatalf("❌ Can't check UFW Status, error: %v", err)
	}

	if strings.Contains(string(output), "Status: active") {
		slog.Info("✅ UFW is ACTIVE.")
	} else {
		log.Fatal("❌ UFW is NOT active. Enable it using: sudo ufw enable")
	}
}
