package commands

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/geekloper/discord-bot-ip-whitelister/bot"
	apperror "github.com/geekloper/discord-bot-ip-whitelister/errors"
	"github.com/geekloper/discord-bot-ip-whitelister/services"
	"github.com/geekloper/discord-bot-ip-whitelister/utils"
)

func init() {
	bot.RegisterCommand(&discordgo.ApplicationCommand{
		Name:        banCmd.Name,
		Description: banCmd.Description,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        banCmd.OptionName,
				Description: banCmd.OptionDesc,
				Required:    true,
			},
		},
	})

	bot.RegisterHandler(banCmd.Name, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options

		optionMap := utils.OptionsToMap(options)

		var msgContent string

		if option, ok := optionMap[banCmd.OptionName]; ok {
			userID := option.StringValue()

			err := services.BanUser(userID, i.Interaction.Member.User.ID)

			// Handle errors and set response message
			if errors.Is(err, apperror.ErrUserDBNotFound) {
				msgContent = "❌ User not found. Please ensure this user is whitelisted before proceeding."
			} else if errors.Is(err, apperror.ErrUserNotAdmin) {
				msgContent = "❌ Sorry you don't have permission to ban a user, please contact a server admin"
			} else if err != nil {
				msgContent = "An unexpected error occurred"
			} else {
				msgContent = "User " + userID + " has been banned successfully 🥳"
			}
		}

		// Send confirmation or error
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: msgContent,
			},
		})
	})
}
