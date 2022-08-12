package main

import (
	"fmt"
	"github.com/artdarek/go-unzip"
	"github.com/pkg/browser"
	"github.com/sqweek/dialog"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inputName, outputName := HandleArguments()

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

	if UiMode {
		yes := dialog.Message("Conversion completed successfully\n want to open the file in the browser?").Title("Completed").YesNo()
		if yes {
			browser.OpenURL(outputName)
		}
	} else {
		fmt.Println("Conversion Finished. New file is at: " + outputName)
	}

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
	curDir := filepath.Dir(filename)
	err := os.RemoveAll(curDir + "/extract")
	if err != nil {
		log.Fatal(err)
	}

	outDir := filepath.Dir(filename) + "/extract"
	uz := unzip.New(filename, outDir)
	err = uz.Extract()
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
