package main

import (
	"database/sql"
	"gists"
	"log"
	"log/slog"
	"os"
	"path/filepath"
)

var (
	DBFILE = filepath.Join(os.TempDir(), "gists.db")
)

func init() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)
}
func main() {

	// Create an empty database
	filename := DBFILE
	err := gists.CreateDatabase(filename)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Created empty database", "file name", filename)
	
	// Open the database for output
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Download gists and write them to database
	pageno := 0
	gistno := 0
	for gistpage := range gists.GetGistPages() {
		slog.Info("Retrieved from GitHub", "page number", pageno)
		for gist := range gists.GetGists(gistpage) {
			gists.LoadDatabase(db, gist)
			slog.Info("Wrote to database", "gist number", gistno, "description", gist.Description)
			gistno++
		}
		pageno++
	}

	// Done
	slog.Info("Loaded database", "number of gists", gistno, "db file name", filename)
}
