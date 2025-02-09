package bot

import "github.com/bwmarrin/discordgo"

var Commands []*discordgo.ApplicationCommand

func RegisterCommand(cmd *discordgo.ApplicationCommand) {
	Commands = append(Commands, cmd)
}
