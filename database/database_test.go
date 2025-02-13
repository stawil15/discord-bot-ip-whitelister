package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) {
	var err error
	db, err = sql.Open("sqlite3", ":memory:") // Use in-memory DB for testing
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	createTableQuery := `CREATE TABLE IF NOT EXISTS rules (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ip TEXT NOT NULL,
		discord_user TEXT UNIQUE NOT NULL,
		status TEXT NOT NULL
	)`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		t.Fatalf("Failed to create test table: %v", err)
	}
}

func TestInitDB(t *testing.T) {
	setupTestDB(t)

	// Check if table exists
	var count int
	err := db.QueryRow("SELECT count(*) FROM sqlite_master WHERE type='table' AND name='rules'").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count, "Table 'rules' should exist")
}

func TestAddRule(t *testing.T) {
	setupTestDB(t)

	err := AddRule("192.168.1.1", "testUser", "active")
	assert.NoError(t, err, "Adding rule should not fail")

	// Verify insertion
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM rules WHERE discord_user = ?", "testUser").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count, "Rule should be added")
}

func TestUserExists(t *testing.T) {
	setupTestDB(t)

	// Insert test data
	_ = AddRule("192.168.1.2", "testUser2", "inactive")

	// Check if user exists
	exists, ip, status := UserExists("testUser2")
	assert.True(t, exists, "User should exist")
	assert.Equal(t, "192.168.1.2", ip, "IP should match")
	assert.Equal(t, "inactive", status, "Status should match")

	// Check for non-existing user
	exists, ip, status = UserExists("nonExistentUser")
	assert.False(t, exists, "User should not exist")
	assert.Empty(t, ip, "IP should be empty for non-existing user")
	assert.Empty(t, status, "Status should be empty for non-existing user")
}

func TestRemoveRule(t *testing.T) {
	setupTestDB(t)

	_ = AddRule("192.168.1.3", "testUser3", "active")
	err := RemoveRule("192.168.1.3")
	assert.NoError(t, err, "Removing rule should not fail")

	// Verify deletion
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM rules WHERE ip = ?", "192.168.1.3").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count, "Rule should be removed")
}
