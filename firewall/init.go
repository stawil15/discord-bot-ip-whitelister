package firewall

import (
	"strings"

	"github.com/geekloper/discord-bot-ip-whitelister/config"
	"github.com/geekloper/discord-bot-ip-whitelister/logger"
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
	logger.Debug("✅ Firewall initialized successfully.", "SERVICES", defaultServices)
}
