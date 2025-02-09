package commands

import "github.com/bwmarrin/discordgo"

func init() {
	RegisterCommand(&discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Returns 'Pong!'",
	})

	RegisterHandler("ping", func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong! 🏓",
			},
		})
	})
}
