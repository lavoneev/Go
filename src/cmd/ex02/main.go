package main

import (
	FScomparer "Day-01/pkg/compareFS"
	"flag"
)

var (
	oldFilename string
	newFilename string
)

func init() {
	flag.StringVar(&oldFilename, "old", "", "Read old FS file")
	flag.StringVar(&newFilename, "new", "", "Read new FS file")
}

func main() {
	flag.Parse()

	FScomparer.CompareFiles(oldFilename, newFilename)
}
