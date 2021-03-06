package cangallo

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type Image struct {
	Time        time.Time
	TotalSize   int64 `yaml:total_size`
	Size        int64
	SHA1        string
	Description string
	Parent      string
	Tag         string
	ParentTag   string `yaml:parent_tag`
}

type Index struct {
	Version int64
	Images  map[string]Image
	Tags    map[string]string
}

type Repo struct {
	Index Index
	Path  string
}

var BasicImageText string = `
name: ""
description: ""
parent: ""
`

func (repo *Repo) LoadIndex(file_name string) {
	file, err := ioutil.ReadFile(file_name)

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	err = yaml.Unmarshal(file, &repo.Index)

	if err != nil {
		log.Fatalf("Error parsing yaml: %v", err)
	}
}

func (repo *Repo) WriteIndex(file_name string) {
	text, err := yaml.Marshal(&repo.Index)

	if err != nil {
		log.Fatalf("Error generating yaml: %v", err)
	}

	err = ioutil.WriteFile(file_name, text, 0644)

	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}
}

func (repo *Repo) SaveIndex() {
	repo.WriteIndex(repo.IndexPath())
}

func (repo *Repo) IndexPath() (file_name string) {
	if repo.Path != "" {
		file_name = fmt.Sprintf("%v/index.yaml", repo.Path)
	} else {
		file_name = "index.yaml"
	}

	return file_name
}

func (repo *Repo) Init() {
	repo.LoadIndex(repo.IndexPath())
}

func (repo *Repo) AddImage(name string, image Image) {
	repo.Index.Images[name] = image
}

func (repo *Repo) Marshal() {
	text, err := yaml.Marshal(&repo.Index)

	if err != nil {
		log.Fatalf("Can not marshal index")
	}

	fmt.Printf("%s\n", text)
}
