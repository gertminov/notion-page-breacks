package unpack

import (
	"errors"
	"github.com/artdarek/go-unzip"
	"log"
	"os"
	"path/filepath"
)

// Extracts zip file and returns path to extracted folder
func ExtractZip(filename string) string {
	outDir := filepath.Dir(filename) + "/extract"
	uz := unzip.New(filename, outDir)
	err := uz.Extract()
	if err != nil {
		log.Fatal(err)
	}

	return outDir
}

func GetHTMLFile(dirname string) (string, error) {
	htmlInDir := filepath.Join(dirname, "*.html")
	matches, err := filepath.Glob(htmlInDir)
	if err != nil {
		log.Fatal(err)
	}
	if len(matches) > 1 {
		return "", errors.New("more than one html file found.\n" +
			"It is not implemented to unpack and convert muliplce html files," +
			" but you can run this programm again in the extracted folder with the *.html input pattern",
		)
	} else if len(matches) == 0 {
		return "", errors.New("No .html files foud in .zip package.\nAre you sure you opened a notion page")
	}

	return matches[0], nil
}

func CopyResourceFiles(extractedDir string, outDir string) {
	filesExtractedDir, err := filepath.Glob(extractedDir + "/*")
	if err != nil {
		log.Fatal("Problem getting resource files from extracted zip folder: " + extractedDir + "\n" + err.Error())
	}
	for _, match := range filesExtractedDir {
		ext := filepath.Ext(match)
		if ext != ".html" {
			outName := filepath.Join(outDir, filepath.Base(match))
			err = os.Rename(match, outName)
			if err != nil {
				log.Fatal("Error while moving resource files from extraced folder to fixed folder:\n" + err.Error())
			}
		}
	}
}
