package main

import (
	DBreader "Day-01/pkg/readDB"
	"flag"
	"fmt"
)

var filename string

func init() {
	flag.StringVar(&filename, "f", "", "Read database")
}

func main() {
	flag.Parse()

	output, err := DBreader.ReadDB(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(output))
}
