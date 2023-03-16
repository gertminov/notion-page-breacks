package medicine

import (
	"log"
	"strings"
)

func InjectMedicine(htmlBytes *[]byte) []byte {
	medicine := "\n hr {\n    page-break-after: always !important;\n    visibility: hidden !important;\n}\n"
	content := string(*htmlBytes)

	index := strings.Index(content, "<style>")
	if index < 0 {
		log.Fatal("no <style> tag in document")
	}

	index = index + 8

	start := content[0:index]
	rest := content[index:]

	newString := start + medicine + rest
	return []byte(newString)
}
