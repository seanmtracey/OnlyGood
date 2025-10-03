package database

import(
	"os"
	"log"
	"path/filepath"
	"database/sql"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func InitDB() {

	homeDir, homeDirErr := os.UserHomeDir()
	if homeDirErr != nil {
		log.Fatalf("Failed to get home directory: %v", homeDirErr)
	}

	appDataDir := filepath.Join(homeDir, "Library", "Application Support", "OnlyGood")
	createSupportDirErr := os.MkdirAll(appDataDir, 0755)
	if createSupportDirErr != nil {
		log.Fatalf("Failed to create app data directory: %v", createSupportDirErr)
	}

	// Database path
	dbPath := filepath.Join(appDataDir, "onlygood.db")
	
	db, openDBErr := sql.Open("sqlite", dbPath)
	if openDBErr != nil {
		log.Fatalf("Failed to connect to database: %v", openDBErr)
	}

	createFeedsTableQuery := `
	CREATE TABLE IF NOT EXISTS feeds (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL,
		url TEXT NOT NULL,
		icon TEXT NOT NULL,
		hash TEXT NOT NULL
	);
	`

	_, createFeedsTableErr := db.Exec(createFeedsTableQuery)
	if createFeedsTableErr != nil {
		log.Fatalf("Failed to create feeds table: %v", createFeedsTableErr)
	}

	createArticlesTableQuery := `
		CREATE TABLE IF NOT EXISTS articles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			url TEXT UNIQUE NOT NULL,
			hash TEXT NOT NULL,
			sentiment REAL NOT NULL,
			alreadyRead BOOLEAN DEFAULT 0
		);
	`

	_, createArticlesTableErr := db.Exec(createArticlesTableQuery)
	if createArticlesTableErr != nil {
		log.Fatalf("Failed to create articles table: %v", createArticlesTableErr)
	}

}
