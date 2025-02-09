package bot

import (
	"log"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	"github.com/geekloper/discord-bot-ip-whitelister/config"
)

var s *discordgo.Session

func InitBot() {
	var err error
	botToken := config.GetEnv("BOT_TOKEN", true)

	s, err = discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal("Invalid bot parameters: ", err)
	}

	s.AddHandler(HandleInteractions)
	s.AddHandler(HandleReady)
}

func OpenSession() error {
	return s.Open()
}

func CloseSession() {
	s.Close()
}

func RemoveCommands(botGuildID string) {
	slog.Info("Removing commands...")
	registeredCommands, err := s.ApplicationCommands(s.State.User.ID, botGuildID)
	if err != nil {
		log.Fatal("Could not fetch registered commands: ", err)
	}

	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, botGuildID, v.ID)
		if err != nil {
			log.Fatalf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}

func RegisterCommands(botGuildID string) {
	slog.Info("Adding commands...")
	for _, v := range Commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, botGuildID, v)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", v.Name, err)
		}
	}
}

func HandleInteractions(session *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := CommandHandlers[i.ApplicationCommandData().Name]; ok {
		h(session, i)
	}
}

func HandleReady(session *discordgo.Session, r *discordgo.Ready) {
	slog.Info("Logged in as", "username", session.State.User.Username, "discriminator", session.State.User.Discriminator)
}
