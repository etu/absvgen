package spec

import (
	"io/ioutil"
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
	filename, _ := filepath.Abs(specFile)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var spec SpecFile

	err = yaml.Unmarshal(yamlFile, &spec)

	if err != nil {
		panic(err)
	}

	return spec
}
