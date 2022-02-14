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
	var companies Companies = make(map[string][]Company)
	sections, err := os.ReadDir("./data")
	if err != nil {
		log.Fatalln(err)
	}
	for _, section_file := range sections {
		section := section_file.Name()
		company_files, err := os.ReadDir("./data/" + section)
		if err != nil {
			log.Fatalln(err)
		}

		for _, company_file := range company_files {
			var company Company
			file, err := os.ReadFile("./data/" + section + "/" + company_file.Name())
			if err != nil {
				log.Fatalln(err)
			}

			err = json.Unmarshal(file, &company)
			if err != nil {
				log.Fatalln(err)
			}

			companies[section] = append(companies[section], company)
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
