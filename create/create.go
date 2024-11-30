package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
)


//todo make mod

func main() {
	log.SetPrefix("Create: ")
	log.SetFlags(0)

	// get the cmd arg
	dryRun := flag.Bool("d", true, "")
	dayNumber := flag.String("day", "", "")
	flag.Parse()
	fmt.Printf("Day: \"%v\"\n", string(*dayNumber))
	day := fmt.Sprintf("Day%s", *dayNumber)

	// create the dir
	directory := fmt.Sprintf("../%s", day)
	if *dryRun {
		fmt.Println("mkdir %s", directory)
	} else {
		os.Mkdir(directory, 0750)
	}

	// create the data files in it
	dataNames := [2]string{"-example.txt", "-data.txt"}
	for _, dataName := range dataNames {
		if *dryRun {
			fmt.Println("make data file %s", dataName)
		} else {
			err := os.WriteFile(fmt.Sprintf("%s/%s%s", directory, day, dataName), []byte("\n\n\n"), 0755)
			if err != nil {
				fmt.Println("unable to write file: %w", err)
			}
		}
	}
	// make the mod
	mod := `module dummy.test/%s

	go 1.23.3
	`
	if *dryRun {
		fmt.Println("make mod file ")
	} else {
		modData := fmt.Sprintf(string(mod), day)
		err := os.WriteFile(fmt.Sprintf("%s/go.mod", directory), []byte(modData), 0755)
		if err != nil {
			fmt.Println("unable to write file: %w", err)
		}
	}
	type Day struct {
		Day string
	}
	// create the main file
	templateData := Day{
		Day: "Day1",
	}
	var tmplFile = "main.tmpl"
	outputName := "main.go"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if *dryRun {
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(os.Stdout, templateData)
		if err != nil {
			panic(err)
		}
	} else {
		var f *os.File
		f, err = os.Create(fmt.Sprintf("%s/%s%s", directory, day, outputName))
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(f, templateData)
		if err != nil {
			panic(err)
		}
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}

}
