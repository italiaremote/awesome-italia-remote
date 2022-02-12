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
	TagsString    string   `json:"-"`
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

	for typeName, companiesType := range companies {
		for k, company := range companiesType {
			tagsString := ""
			for _, tag := range company.Tags {
				tagsString += tag + " - "
			}
			tagsString = strings.Trim(tagsString, " - ")
			companies[typeName][k].TagsString = tagsString
		}
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
