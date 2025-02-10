package commands

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/geekloper/discord-bot-ip-whitelister/bot"
	apperror "github.com/geekloper/discord-bot-ip-whitelister/errors"
	"github.com/geekloper/discord-bot-ip-whitelister/services"
	"github.com/geekloper/discord-bot-ip-whitelister/utils"
)

const (
	cmdName        = "whitelist"
	cmdDescription = "whitelist your IP"
	cmdOptionName  = "ip"
	cmdOptionDesc  = "Your IP address"
)

func init() {
	bot.RegisterCommand(&discordgo.ApplicationCommand{
		Name:        cmdName,
		Description: cmdDescription,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        cmdOptionName,
				Description: cmdOptionDesc,
				Required:    true,
			},
		},
	})

	bot.RegisterHandler(cmdName, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options

		optionMap := utils.OptionsToMap(options)

		var msgformat string

		if option, ok := optionMap[cmdOptionName]; ok {
			ip := option.StringValue()

			err := services.WhitelistIP(ip, i.Interaction.Member.User.ID)

			// Handle errors and set response message
			if errors.Is(err, apperror.ErrInvalidIpFormat) {
				msgformat = "❌ Your IP is not valid, please provide a valid IP"
			} else if errors.Is(err, apperror.ErrBannedUser) {
				msgformat = "❌ Sorry you're banned, please contact a server admin"
			} else if err != nil {
				msgformat = "An unexpected error occurred"
			} else {
				msgformat = "Your IP has been added successfully 🥳"
			}
		}

		// Send confirmation or error
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: msgformat,
			},
		})
	})
}
