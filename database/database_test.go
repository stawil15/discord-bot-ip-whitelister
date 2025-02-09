package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) {
	var err error
	db, err = sql.Open("sqlite3", ":memory:") // Use in-memory DB for testing
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	createTableQuery := `CREATE TABLE IF NOT EXISTS rules (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ip TEXT UNIQUE NOT NULL,
		discord_user TEXT NOT NULL,
		status TEXT NOT NULL
	)`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}
}

func TestAddRule(t *testing.T) {
	setupTestDB(t)

	err := AddRule("192.168.1.1", "user1", "active")
	if err != nil {
		t.Fatalf("Failed to add rule: %v", err)
	}

	dumpAllRules(t)
}

func TestUpdateRule(t *testing.T) {
	setupTestDB(t)

	_ = AddRule("192.168.1.1", "user1", "active")
	err := AddRule("192.168.1.1", "user2", "inactive") // Should update existing rule
	if err != nil {
		t.Fatalf("Failed to update rule: %v", err)
	}

	dumpAllRules(t)
}

func TestRemoveRule(t *testing.T) {
	setupTestDB(t)

	_ = AddRule("192.168.1.1", "user1", "active")
	err := RemoveRule("192.168.1.1")
	if err != nil {
		t.Fatalf("Failed to remove rule: %v", err)
	}

	// Verify deletion
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM rules WHERE ip = ?", "192.168.1.1").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query database: %v", err)
	}

	if count != 0 {
		t.Fatalf("Rule was not deleted")
	}

	dumpAllRules(t)
}

func dumpAllRules(t *testing.T) {
	rows, err := db.Query("SELECT id, ip, discord_user, status FROM rules")
	if err != nil {
		t.Fatalf("Failed to query rules: %v", err)
	}
	defer rows.Close()

	t.Log("Dumping all rules:")
	for rows.Next() {
		var id int
		var ip, discordUser, status string
		if err := rows.Scan(&id, &ip, &discordUser, &status); err != nil {
			t.Fatalf("Failed to scan row: %v", err)
		}
		t.Logf("ID: %d, IP: %s, User: %s, Status: %s", id, ip, discordUser, status)
	}

	if err := rows.Err(); err != nil {
		t.Fatalf("Error iterating rows: %v", err)
	}
}
