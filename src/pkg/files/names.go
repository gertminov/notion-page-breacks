package files

import (
	"path/filepath"
	"strconv"
	"strings"
)

func GetOutputName(outputName string, inputName string, idx int) string {
	if outputName == "" {
		newOutputName := filepath.Base(inputName)
		newOutputName = trimSuffix(newOutputName, ".html")
		if len(newOutputName) > 33 {
			newOutputName = newOutputName[:len(newOutputName)-33]
		}
		outputName = newOutputName + "-breaks.html"
	} else if filepath.Ext(outputName) != ".html" {
		outputName = outputName + ".html"
	}
	return strconv.Itoa(idx) + "-" + outputName
}

func trimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
