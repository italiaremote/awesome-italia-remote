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

const homePath = "../../../"
const dataPath = homePath + "data/"

func main() {
	var companies Companies = make(map[string][]Company)
	sections, err := os.ReadDir(dataPath)
	if err != nil {
		log.Fatalln(err)
	}
	for _, sectionFile := range sections {
		section := sectionFile.Name()
		companyFiles, err := os.ReadDir(dataPath + section)
		if err != nil {
			log.Fatalln(err)
		}

		for _, companyFiles := range companyFiles {
			var company Company
			file, err := os.ReadFile(dataPath + section + "/" + companyFiles.Name())
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

	f, err := os.Create(homePath + "README.md")
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
