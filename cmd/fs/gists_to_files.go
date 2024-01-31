package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/philhanna/gists"
)

const (
	usage = `usage: gists_to_files [DBFILE] [FSROOT]
Reads gists from a database and writes their contents to the file system

positional arguments:
  DBFILE         database file name
  FSROOT         directory in which to write files

options:
  -f, --force    recreate directory if it exists
`
)

var (
	err      error
	optForce bool
	DBFILE   string
	FSROOT   string
)

func main() {

	// Set log style to include file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Override flag.Usage so that it prints this program's usage
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}

	// Parse the flags
	flag.BoolVar(&optForce, "f", false, "recreate directory if it exists")
	flag.BoolVar(&optForce, "force", false, "recreate directory if it exists")
	flag.Parse()

	// Verify command line arguments
	switch flag.NArg() {
	case 0:
		fmt.Fprintln(os.Stderr, "No command line arguments specified")
		fmt.Fprintln(os.Stderr, "Try --help for help")
		return
	case 1:
		fmt.Fprintln(os.Stderr, "No output directory specified")
		fmt.Fprintln(os.Stderr, "Try --help for help")
		return
	case 2:
		DBFILE = flag.Arg(0)
		if !gists.FileExists(DBFILE) {
			log.Fatalf("%q database file does not exist", DBFILE)
		}
		FSROOT = flag.Arg(1)
		_, err = os.Stat(FSROOT)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Run the download loop
	err = Run()
	if err != nil {
		log.Fatal(err)
	}
}

// forceCreateDirectory prompts the user to indicate whether to delete
// and recreate the chosen output director.
func forceCreateDirectory() error {
	fmt.Printf("Are you sure you want to delete directory %s? (y/N) ", FSROOT)
	reader := bufio.NewReader(os.Stdin)
	yesno, _ := reader.ReadString('\n')
	yesno = strings.TrimSpace(yesno)
	switch yesno {
	case "Y", "y":
		err = os.RemoveAll(FSROOT)
		if err != nil {
			return err
		}
		err = os.Mkdir(FSROOT, 0755)
		if err != nil {
			return err
		}
		fmt.Println()
	default:
		// Get out
		err = fmt.Errorf("exit, not deleting %s", FSROOT)
		return err
	}
	// OK
	return nil
}

// Run downloads all the gists and writes them to the database
func Run() error {

	// Create output directory
	if optForce {
		err = forceCreateDirectory()
		if err != nil {
			return err
		}
	}

	// Open the database for reading
	log.Printf("Opening %s for input", DBFILE)
	db, err := sql.Open("sqlite3", DBFILE)
	if err != nil {
		return err
	}
	defer db.Close()

	// Read each gist file from database
	sql := `
SELECT		f.id, f.filename, g.description, g.created_at, f.language, f.size, f.contents
FROM		gistfiles f
LEFT JOIN	gists g
ON			f.id = g.id
ORDER BY	g.created_at
;
`
	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	var (
		id          string
		filename    string
		description string
		created_at  string
		language    string
		size        int
		contents    []byte
	)

	for rows.Next() {

		// Read a gist file record
		err = rows.Scan(&id, &filename, &description, &created_at, &language, &size, &contents)
		if err != nil {
			return err
		}

		// Create a directory to contain it
		dirPath := filepath.Join(FSROOT, created_at)
		err = os.Mkdir(dirPath, 0755)
		if err != nil {
			return err
		}

		// Write the file contents to the directory
		gistFileName := filepath.Join(dirPath, filename)
		log.Printf("Writing %q to %q\n", filename, dirPath)
		err = os.WriteFile(gistFileName, contents, 0644)
		if err != nil {
			return err
		}

		// Write the metadata to the directory in JSON format
		mdFileName := filepath.Join(dirPath, "metadata.json")
		fp, err := os.Create(mdFileName)
		if err != nil {
			return err
		}
		defer fp.Close()
		fmt.Fprintf(fp, "{\n")
		fmt.Fprintf(fp, "  %q : %q,\n", "ID", id)
		fmt.Fprintf(fp, "  %q : %q,\n", "Description", description)
		fmt.Fprintf(fp, "  %q : %q,\n", "Language", language)
		fmt.Fprintf(fp, "  %q : %d\n", "Size", size)
		fmt.Fprintf(fp, "}\n")
	}

	return nil
}
