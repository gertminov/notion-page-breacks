package main

import (
	"fmt"
	"github.com/artdarek/go-unzip"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

	fmt.Println("Conversion Finished. New file is at: " + outputName)

}

func handleArguments() (string, string) {
	var inputFile string
	amtArgs := len(os.Args)
	if amtArgs < 2 {
		log.Fatal("No input file given.\nRun with -h to see available arguments")
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

	outputName := ""
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

func getOutputName(outputName string, inputName string) string {
	if outputName == "" {
		newOutputName := TrimSuffix(inputName, ".html")
		newOutputName = newOutputName[:len(newOutputName)-33]
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

//Extracts zip file and returns path to extracted folder
func extractZip(filename string) string {
	outDir := filepath.Dir(filename) + "/extract"
	uz := unzip.New(filename, outDir)
	err := uz.Extract()
	if err != nil {
		log.Fatal(err)
	}

	return outDir
}

func getHTMLFile(dirname string) string {
	filename := ""
	dir, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range dir {
		if strings.HasSuffix(info.Name(), ".html") {
			filename = info.Name()
		}
	}

	if filename == "" {
		fmt.Println("Couldn't find an HTML file in the extracted directory, something must have gone wrong.\nAre you shoure you exportet the notion page as HTML?")
	}

	return filename
}
