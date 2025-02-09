package database

import (
	"database/sql"
	"log"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "rules.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	createTableQuery := `CREATE TABLE IF NOT EXISTS rules (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ip TEXT NOT NULL,
		discord_user TEXT UNIQUE NOT NULL,
		status TEXT NOT NULL
	)`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func AddRule(ip, discordUser, status string) error {
	_, err := db.Exec("INSERT INTO rules (ip, discord_user, status) VALUES (?, ?, ?) ON CONFLICT(discord_user) DO UPDATE SET status=excluded.status, ip=excluded.ip", ip, discordUser, status)
	return err
}

func RemoveRule(ip string) error {
	_, err := db.Exec("DELETE FROM rules WHERE ip = ?", ip)
	return err
}

func UserIpExists(discordUser string) (bool, string) {
	var ip string
	err := db.QueryRow("SELECT ip FROM rules WHERE discord_user = ?", discordUser).Scan(&ip)
	if err == sql.ErrNoRows {
		return false, ""
	} else if err != nil {
		log.Fatalf("Failed to check if user exists: %v", err)
	}
	return true, ip
}

func DumpAllRules() {
	slog.Info("Dump all DB rules")
	slog.Info("=================")

	rows, err := db.Query("SELECT id, ip, discord_user, status FROM rules")
	if err != nil {
		log.Fatalf("Failed to query rules: %v", err)
		return
	}
	defer rows.Close()

	i := 1
	for rows.Next() {
		var id int
		var ip, discordUser, status string
		if err := rows.Scan(&id, &ip, &discordUser, &status); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		slog.Info("RULE ", "ID", i, "Ip", ip, "User", discordUser, "Status", status)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating rows: %v", err)
	}

	slog.Info("=================")
}
