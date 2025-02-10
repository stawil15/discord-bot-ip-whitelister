package services

import (
	"github.com/geekloper/discord-bot-ip-whitelister/database"
	apperrors "github.com/geekloper/discord-bot-ip-whitelister/errors"
	"github.com/geekloper/discord-bot-ip-whitelister/firewall"
	"github.com/geekloper/discord-bot-ip-whitelister/logger"
	"github.com/geekloper/discord-bot-ip-whitelister/utils"
)

// WhitelistIP validates the IP, removes existing rules if necessary, applies the new rule, and updates the database.
func WhitelistIP(ip string, userID string) error {
	// Validate IP format
	if !utils.ValidateIP(ip) {
		logger.Warn("Invalid IP format attempted for whitelisting", "ip", ip, "userID", userID)
		return apperrors.ErrInvalidIpFormat
	}

	// Check if user already has an existing whitelisted IP and not banned
	if exists, oldIP, status := database.UserExists(userID); exists {

		if status == "deny" {
			logger.Info("User is banned", "oldIP", oldIP, "newIP", ip, "userID", userID)
			return apperrors.ErrBannedUser
		}

		logger.Info("User already has a whitelisted IP, replacing it", "oldIP", oldIP, "newIP", ip, "userID", userID)

		// Remove the old firewall rule before applying a new one
		err := firewall.DeleteUFWRule(oldIP)
		if err != nil {
			logger.Error("Failed to remove existing UFW rule before whitelisting new IP", "oldIP", oldIP, "userID", userID, "error", err)
			return err
		}
	}

	// Apply the new firewall rule
	err := firewall.AllowUFWRule(ip)
	if err != nil {
		logger.Error("Failed to whitelist IP in UFW", "ip", ip, "userID", userID, "error", err)
		return err
	}

	// Store the new rule in the database after successful firewall application
	if err := database.AddRule(ip, userID, "allow"); err != nil {
		logger.Error("Failed to store new whitelisted IP in database", "ip", ip, "userID", userID, "error", err)
		return err
	}

	logger.Info("IP successfully whitelisted", "ip", ip, "userID", userID)
	return nil
}
