package spec

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type SpecFile struct {
	Width  uint16      `yaml:"width"`
	Height uint16      `yaml:"height"`
	Layers []SpecLayer `yaml:"layers"`
}

type SpecLayer struct {
	Module string   `yaml:"module"`
	Size   uint8    `yaml:"size"`
	Colors []string `yaml:"colors"`
}

func New(specFile string) SpecFile {
	var spec SpecFile
	var yamlFile []byte
	var err error

	// If we get special filename "-", read from stdin instead of a file
	if specFile == "-" {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			yamlFile = append(yamlFile, byte(10))
			yamlFile = append(yamlFile, scanner.Bytes()...)
		}
	} else {
		filename, _ := filepath.Abs(specFile)
		yamlFile, err = ioutil.ReadFile(filename)

		if err != nil {
			panic(err)
		}
	}

	err = yaml.Unmarshal(yamlFile, &spec)

	if err != nil {
		panic(err)
	}

	return spec
}
