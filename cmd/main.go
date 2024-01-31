package main

import (
	"fmt"
	"gists"
	"log"
	"os"
	"path/filepath"
)

func main() {
	pageno := 0
	for gistpage := range gists.GetGistPages() {
		pageno++
		filename := filepath.Join("testdata", fmt.Sprintf("gistpage%d.json", pageno))
		err := writePage(filename, gistpage)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Wrote page %d\n", pageno)
	}
}

func writePage(filename string, gistpage string) error {
	err := os.WriteFile(filename, []byte(gistpage), 0644)
	return err
}