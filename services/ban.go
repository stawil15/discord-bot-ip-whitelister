package services

import (
	"github.com/geekloper/discord-bot-ip-whitelister/database"
	"github.com/geekloper/discord-bot-ip-whitelister/errors"
	"github.com/geekloper/discord-bot-ip-whitelister/firewall"
	"github.com/geekloper/discord-bot-ip-whitelister/logger"
)

// Ban applies UFW deny rule, and updates the database.
func BanUser(userID, userAdminID string) error {
	// Check if userAdminID is admin
	if _, exists := adminIDs[userAdminID]; !exists {
		return errors.ErrUserNotAdmin
	}

	// Check if user already has an existing whitelisted IP
	if exists, ip, _ := database.UserExists(userID); exists {

		err := firewall.DenyUFWRule(ip)
		if err != nil {
			logger.Error("Failed to deny UFW rule", "ip", ip, "userID", userID, "error", err)
			return err
		}

		// Store the new rule in the database after successful firewall application
		if err := database.AddRule(ip, userID, "deny"); err != nil {
			logger.Error("Failed to ban user IP in database", "ip", ip, "userID", userID, "error", err)
			return err
		}

		logger.Info("User banned", "ip", ip, "userID", userID)

		return nil
	} else {
		return errors.ErrUserDBNotFound
	}
}
