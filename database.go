package gists

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
	"strings"
)

// CreateDatabase reads gists from a channel and writes them
// to an SQLite3 database.
func CreateDatabase(filename string) error {

	// Delete the output database file, if it exists
	if FileExists(filename) {
		os.Remove(filename)
	}

	// Open the database for output
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return err
	}
	defer db.Close()

	// Begin a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	createTablesSQL := `
CREATE TABLE gists (
    id          TEXT,
    url         TEXT,
    description TEXT,
    created_at  TEXT
);
CREATE TABLE gistfiles (
    id          TEXT,
    filename    TEXT,
    language    TEXT,
    size        TEXT,
    contents    BLOB
);
CREATE VIEW joined AS
    SELECT      g.id, f.filename, f.language, g.description, f.contents
    FROM        gistfiles f
    LEFT JOIN   gists g
    ON          f.id = g.id
    ORDER BY    4
    ;
`
	// Create the tables
	_, err = tx.Exec(createTablesSQL)
	if err != nil {
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// FileExists returns true if the specified file exists
func FileExists(filename string) bool {
	if strings.HasPrefix(filename, "~/") {
		simpleName := strings.TrimPrefix(filename, "~/")
		dirname, _ := os.UserHomeDir()
		filename = filepath.Join(dirname, simpleName)
	}
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// LoadDatabase writes a gist to the database
func LoadDatabase(db *sql.DB, gist Gist) error {

	// Begin a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Write gists to tables
	insertSQL := `INSERT INTO gists (id, url, description, created_at) VALUES(?, ?, ?, ?)`
	_, err = tx.Exec(insertSQL, gist.ID, gist.URL, gist.Description, gist.CreatedAt)
	if err != nil {
		return err
	}
	for _, file := range gist.Files {
		blob, err := file.GetContents()
		if err != nil {
			return err
		}
		insertSQL = `INSERT INTO gistfiles (id, filename, language, size, contents) VALUES(?, ?, ?, ?, ?)`
		_, err = tx.Exec(insertSQL, gist.ID, file.Filename, file.Language, file.Size, blob)
		if err != nil {
			return err
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	// No error
	return nil
}
