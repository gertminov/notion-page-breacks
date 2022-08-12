package main

import (
	"fmt"
	"github.com/sqweek/dialog"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var UiMode = false

func HandleArguments() (string, string) {
	var inputFile string
	outputName := ""
	amtArgs := len(os.Args)
	if amtArgs < 2 {
		inputFile = useFilePicker()
		fmt.Println("No input file given.\nRun with -h to see available arguments")
		return inputFile, outputName
	}
	firstArg := os.Args[1]

	if firstArg == "-h" {
		fmt.Println("Usage: ")
		fmt.Println("       notion-breaks input.html     or     notion-breaks input.zip")
		fmt.Println("Arguments:")
		fmt.Println("-o [output]    custom output name")
		os.Exit(0)
	} else if !(strings.Contains(firstArg, ".html") || strings.Contains(firstArg, ".zip")) {
		log.Fatal("only .html files and .zip archives can be processed")
	} else if strings.Contains(firstArg, ".zip") {
		dirPath := extractZip(firstArg)
		inputFile = dirPath + "/" + getHTMLFile(dirPath)

	} else if strings.Contains(firstArg, ".html") {
		inputFile = firstArg
	}

	if amtArgs > 2 {
		secondArg := os.Args[2]
		if secondArg == "-o" {
			outputName = os.Args[3]
			dirPath := func() string {
				if filepath.Dir(inputFile) == "." {
					return ""
				} else {
					return filepath.Dir(inputFile)
				}
			}()
			outputName = dirPath + "/" + outputName
			if !strings.HasSuffix(outputName, ".html") {
				outputName = outputName + ".html"
			}
		}

	}

	return inputFile, outputName
}

func useFilePicker() string {
	inputFile := ""
	filename, err := dialog.File().Load()
	if err != nil {
		log.Fatal(err)
	}
	UiMode = true
	if strings.HasSuffix(filename, ".zip") {
		dirPath := extractZip(filename)
		inputFile = dirPath + "/" + getHTMLFile(dirPath)
	} else if strings.HasSuffix(filename, ".html") {
		inputFile = filename
	} else {
		fmt.Println("Only .zip or .html files can be picked")
	}

	return inputFile
}
