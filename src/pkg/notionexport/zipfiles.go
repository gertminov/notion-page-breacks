package notionexport

import (
	"github.com/fatih/color"
	"notion-page-breacks/src/pkg/unpack"
)

func ExtractZipFiles(zipFiles []string, outDir string) []string {
	var htmlFiles []string

	for _, file := range zipFiles {
		extracedDirPath := unpack.ExtractZip(file)
		htmlFile, err := unpack.GetHTMLFile(extracedDirPath)
		unpack.CopyResourceFiles(extracedDirPath, outDir)
		if err != nil {
			color.Red(err.Error())
			return nil
		}
		htmlFiles = append(htmlFiles, htmlFile)
	}
	return htmlFiles
}
