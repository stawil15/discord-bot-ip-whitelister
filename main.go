package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"

	"github.com/geekloper/discord-bot-ip-whitelister/commands"
	"github.com/geekloper/discord-bot-ip-whitelister/config"
	"github.com/geekloper/discord-bot-ip-whitelister/database"
	"github.com/geekloper/discord-bot-ip-whitelister/firewall"

	"github.com/bwmarrin/discordgo"
)

var s *discordgo.Session

func initBot() {
	var err error
	botToken := config.GetEnv("BOT_TOKEN", true)

	s, err = discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal("Invalid bot parameters: ", err)
	}

	s.AddHandler(handleInteractions)
	s.AddHandler(handleReady)
}

func handleInteractions(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commands.CommandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}

func handleReady(s *discordgo.Session, r *discordgo.Ready) {
	slog.Info("Logged in as", "username", s.State.User.Username, "discriminator", s.State.User.Discriminator)
}

func main() {
	// Load environment variables
	config.LoadEnv()

	// Required environment variables
	botGuildID := config.GetEnv("BOT_GUILD_ID", true)
	servicePorts := config.GetEnv("SERVICE_PORTS", true)
	deleteCommands := config.GetEnv("DELETE_COMMADS", false) == "true"

	// Initialize modules
	database.InitDB()
	firewall.InitFirewall(servicePorts)
	initBot()

	// Log all rules
	database.DumpAllRules()

	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open discord session: %v", err)
	}

	// Manage commands
	if deleteCommands {
		removeCommands(botGuildID)
	}
	registerCommands(botGuildID)

	// Graceful Shutdown Handling
	defer s.Close()
	waitForShutdown()
}

func removeCommands(botGuildID string) {
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

func registerCommands(botGuildID string) {
	slog.Info("Adding commands...")
	for _, v := range commands.Commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, botGuildID, v)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", v.Name, err)
		}
	}
}

func waitForShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	slog.Info("Press Ctrl+C to exit")
	<-stop
	slog.Info("Gracefully shutting down.")
}
