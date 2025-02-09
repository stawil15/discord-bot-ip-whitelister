package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"

	"github.com/geekloper/discord-bot-ip-whitelister/bot"
	_ "github.com/geekloper/discord-bot-ip-whitelister/commands"
	"github.com/geekloper/discord-bot-ip-whitelister/config"
	"github.com/geekloper/discord-bot-ip-whitelister/database"
	"github.com/geekloper/discord-bot-ip-whitelister/firewall"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Required environment variables
	botGuildID := config.GetEnv("BOT_GUILD_ID", true)
	servicePorts := config.GetEnv("SERVICE_PORTS", true)
	serviceProtocols := config.GetEnv("SERVICE_PROTOCOLS", true)
	deleteCommands := config.GetEnv("DELETE_COMMADS", false) == "true"

	// Initialize modules
	database.InitDB()
	firewall.InitFirewall(servicePorts, serviceProtocols)
	bot.InitBot()

	// Log all rules in debug level mode
	database.DumpAllRules()

	err := bot.OpenSession()
	if err != nil {
		log.Fatalf("Cannot open discord session: %v", err)
	}

	if deleteCommands {
		bot.RemoveCommands(botGuildID)
	}
	bot.RegisterCommands(botGuildID)

	// Graceful Shutdown Handling
	defer bot.CloseSession()
	waitForShutdown()
}

func waitForShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	slog.Info("Press Ctrl+C to exit")
	<-stop
	slog.Info("Gracefully shutting down.")
}
