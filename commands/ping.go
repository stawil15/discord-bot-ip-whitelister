package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/geekloper/discord-bot-ip-whitelister/bot"
)

func init() {
	bot.RegisterCommand(&discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Returns 'Pong!'",
	})

	bot.RegisterHandler("ping", func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong! 🏓",
			},
		})
	})
}
