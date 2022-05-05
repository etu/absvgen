package doc

import (
	"fmt"
	"strings"

	"github.com/etu/absvgen/src/spec"
)

func WrapDocObjects(specFile spec.SpecFile, docObjects []string) string {
	var objects []string

	objects = append(objects, "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?>")
	objects = append(objects, fmt.Sprintf("<svg width=\"%d\" height=\"%d\" xmlns=\"http://www.w3.org/2000/svg\" version=\"1.1\">", specFile.Width, specFile.Height))

	for _, v := range docObjects {
		objects = append(objects, v)
	}

	objects = append(objects, "</svg>")

	return strings.Join(objects, "\n")
}
