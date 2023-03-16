package files

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

func ReadFiles(pattern string) (htmlFiles []string, zipFiles []string, err error) {
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, nil, errors.New("Error: Could not find any files matching the input pattern: " + pattern)
	}

	for _, file := range files {
		if strings.Contains(file, ".html") {
			htmlFiles = append(htmlFiles, file)
		} else if strings.Contains(file, ".zip") {
			zipFiles = append(zipFiles, file)
		}
	}

	if len(htmlFiles) == 0 && len(zipFiles) == 0 {
		return nil, nil, errors.New("Error: No files to process.\nCan only process .html and .zip files.")
	}

	fmt.Println("Following files will be processed:")
	for _, file := range htmlFiles {
		fmt.Println(" " + file)
	}
	for _, file := range zipFiles {
		fmt.Println(" " + file)
	}
	fmt.Println()

	return htmlFiles, zipFiles, nil
}
