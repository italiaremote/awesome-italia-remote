package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"text/template"
)

type Companies map[string][]Company

type Company struct {
	Name          string   `json:"name"`
	URL           string   `json:"url"`
	CareerPageURL string   `json:"career_page_url"`
	Type          string   `json:"type,omitempty"`
	Tags          []string `json:"tags,omitempty"`
}

func (c Company) GetTagsString() string {
	return strings.Join(c.Tags, " - ")
}

func main() {
	var companies Companies

	file, err := os.ReadFile("./data.json")
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(file, &companies)
	if err != nil {
		log.Fatalln(err)
	}

	templ, err := template.ParseFiles("./template.md")
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create("../README.md")
	if err != nil {
		log.Fatalln(err)
	}

	err = templ.Execute(f, companies)
	if err != nil {
		log.Fatalln(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
