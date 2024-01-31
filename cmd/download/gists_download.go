package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"gists"
	"log"
	"log/slog"
	"os"
	"strings"
)

const (
	usage = `usage: gists_download [DBFILE]
Downloads this user's gists from Github into an SQLite3 database

positional arguments:
  DBFILE         database file name

options:
  -f, --force    create new database if old one exists
`
)

var (
	err      error
	optForce bool
	DBFILE   string
)

func init() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)
}

func main() {
	// Set log style to include file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Override flag.Usage so that it prints this program's usage
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}

	// Parse the flags
	flag.BoolVar(&optForce, "f", false, "Create new databse if old one exists")
	flag.BoolVar(&optForce, "force", false, "Create new databse if old one exists")
	flag.Parse()

	// Verify command line arguments
	switch flag.NArg() {
	case 0:
		log.Println("No command line arguments specified.")
		log.Fatalln("Try --help for help")
	default:
		DBFILE = flag.Arg(0)
		if optForce {
			if gists.FileExists(DBFILE) {
				fmt.Printf("Are you sure you want to delete database %s? (y/N) ", DBFILE)
				reader := bufio.NewReader(os.Stdin)
				yesno, _ := reader.ReadString('\n')
				yesno = strings.TrimSpace(yesno)
				switch yesno {
				case "Y", "y":
					err = os.Remove(DBFILE)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println()
				default:
					// Get out
					log.Fatalf("Exit, not deleting %s", DBFILE)
				}
			}
		}
	}

	// Die if database already exists
	if gists.FileExists(DBFILE) {
		log.Fatalf("Database file %s already exists\n", DBFILE)
	}

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
