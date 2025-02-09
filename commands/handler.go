package commands

import "github.com/bwmarrin/discordgo"

// CommandHandlers stores the mapping of command names to their handler functions
var CommandHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))

// RegisterHandler adds a new command handler
func RegisterHandler(name string, handler func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	CommandHandlers[name] = handler
}
