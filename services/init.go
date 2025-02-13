// services/init.go
package services

import (
	"strings"

	"github.com/geekloper/discord-bot-ip-whitelister/logger"
)

var adminIDs map[string]struct{}

// InitServices initializes required services.
func InitServices(envAdminIDs string) {
	// Split the string into a slice
	ids := strings.Split(envAdminIDs, ",")

	adminIDs = make(map[string]struct{})
	for _, id := range ids {
		adminIDs[strings.TrimSpace(id)] = struct{}{}
	}

	logger.Debug("✅ Services initialized successfully.", "adminIDs", adminIDs)
}
