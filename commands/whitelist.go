package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/geekloper/discord-bot-ip-whitelister/firewall"
)

const (
	cmdName        = "whitelist"
	cmdDescription = "whitelist your IP"
	cmdOptionName  = "ip"
	cmdOptionDesc  = "Your IP address"
)

func init() {
	RegisterCommand(&discordgo.ApplicationCommand{
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

	RegisterHandler(cmdName, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options

		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		msgformat := "Your IP have been added successfully 🥳"

		if option, ok := optionMap[cmdOptionName]; ok {
			ip := option.StringValue()

			err := firewall.AllowUFWRule(ip, i.Interaction.Member.User.ID)

			if err == nil {
				msgformat += "> ip : " + ip
				msgformat += "> discord id : " + i.Interaction.Member.User.ID
			} else {
				fmt.Printf("Error %s", err)
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
