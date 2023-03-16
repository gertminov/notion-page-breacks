package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"notion-page-breacks/src/pkg/files"
	"notion-page-breacks/src/pkg/medicine"
	"notion-page-breacks/src/pkg/unpack"
	"os"
	"path/filepath"
)

const (
	OUTPUT_NAME = "out"
)

func main() {
	var outName string
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal("could not get Working directory... wtf??")
	}
	app := &cobra.Command{
		Use:   "notion",
		Short: "create pagebreaks for notion",
		Long:  "Insert CSS to start a new page at every --- in the notion document.\nYou can either input a single .html file, or a whole .zip package",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				color.Red("Error:  No input name specified.")
				fmt.Println("Please specify an input name like so:")
				color.Green("\n\tnotion-pb mydocument.html")
				return
			}
			inputName := args[0]
			htmlFiles, zipFiles, err := files.ReadFiles(inputName)
			if err != nil {
				color.Red(err.Error())
				return
			}
			if len(zipFiles) > 0 {
				for _, file := range zipFiles {
					extracedDirPath := unpack.ExtractZip(file)
					htmlFile, err := unpack.GetHTMLFile(extracedDirPath)
					if err != nil {
						color.Red(err.Error())
						return
					}
					htmlFiles = append(htmlFiles, htmlFile)
				}
			}

			outDir := filepath.Join(workingDir, "fixed")
			err = os.MkdirAll(outDir, os.FileMode(666))
			if err != nil {
				log.Fatal("could not create 'fixed' folder")
			}

			for i, file := range htmlFiles {
				readFile, err := os.ReadFile(file)
				if err != nil {
					log.Fatal(err)
				}
				patchedHTML := medicine.InjectMedicine(&readFile)
				outputName := files.GetOutputName(outName, file, i)

				err = os.WriteFile(filepath.Join(outDir, outputName), patchedHTML, os.FileMode(666))
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Conversion Finished. New file is at: " + outputName)
			}
			fmt.Println("flags: " + outName)
		},
	}
	app.Flags().StringVarP(&outName, "out", "o", "", "set output name")
	if err := app.Execute(); err != nil {
		log.Fatal(err)
	}
}
