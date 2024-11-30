package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

//TODO make a template code file and use go substite thing 
//TODO make a data file

func main() {
	log.SetPrefix("Create: ")
	log.SetFlags(0)

	// get the cmd arg
	dayNumber := flag.String("Day", "", "")
	flag.Parse()
	fmt.Printf("Day: \"%v\"\n", string(*dayNumber))
	day:= fmt.Sprintf("Day%s", *dayNumber)

	// create the dir
	os.Mkdir(fmt.Sprintf("../%s", day), 0750)
	// create the file in it 
	err := os.WriteFile(fmt.Sprintf("../%s/%s.go", day, day), []byte("Hello\n"), 0755)
    if err != nil {
        fmt.Printf("unable to write file: %w", err)
    }

}
