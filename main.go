package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	inputName, outputName := handleArguments()

	fmt.Println(outputName)

	medicine := "\n hr {\n    page-break-after: always !important;\n    visibility: hidden !important;\n}\n"

	contentBytes, err := ioutil.ReadFile(inputName)
	if err != nil {
		log.Fatal(err)
	}

	content := string(contentBytes)

	index := strings.Index(content, "<style>")
	if index < 0 {
		log.Fatal("no <style> tag in document")
	}

	index = index + 8

	start := content[0:index]
	rest := content[index:]

	newString := start + medicine + rest
	data := []byte(newString)
	outputName = getOutputName(outputName, inputName)
	err = ioutil.WriteFile(outputName, data, 0)
	if err != nil {
		log.Fatal(err)
	}

}

func handleArguments() (string, string) {
	var inputFile string
	amtArgs := len(os.Args)
	if amtArgs < 2 {
		log.Fatal("No input file given.\nRun with -h to see available arguments")
	}
	firstArg := os.Args[1]
	outputName := ""
	if amtArgs > 2 {
		secondArg := os.Args[2]
		if secondArg == "-o" {
			outputName = os.Args[3]
			if !strings.HasSuffix(outputName, ".html") {
				outputName = outputName + ".html"
			}
		}

	}

	if firstArg == "-h" {
		fmt.Println("Usage: notion-breaks input.html")
		fmt.Println("Arguments:")
		fmt.Println("-o [output]    custom output name")
		os.Exit(0)
	} else if !strings.Contains(firstArg, ".html") {
		log.Fatal("only .html files can be processed")
	} else if strings.Contains(firstArg, ".html") {
		inputFile = firstArg
	}

	return inputFile, outputName
}

func getOutputName(outputName string, inputName string) string {
	if outputName == "" {
		newOutputName := TrimSuffix(inputName, ".html")
		outputName = newOutputName + "-breaks.html"
	}
	return outputName
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
