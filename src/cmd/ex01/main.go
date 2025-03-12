package main

import (
	DBcomparer "Day-01/pkg/compareDB"
	"flag"
)

var (
	oldFilename string
	newFilename string
)

func init() {
	flag.StringVar(&oldFilename, "old", "", "Read old database")
	flag.StringVar(&newFilename, "new", "", "Read new database")
}

func main() {
	flag.Parse()

	DBcomparer.Run(oldFilename, newFilename)
}
