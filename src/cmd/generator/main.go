package main

import (
	"encoding/json"
	"github.com/italiaremote/awesome-italia-remote/pkg/cmp"
	"log"
	"os"
	"text/template"
)

const homePath = "../../../"
const dataPath = homePath + "data/"

func main() {
	companies := make(map[string][]cmp.Company)
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
			var company cmp.Company
			file, err := os.ReadFile(dataPath + section + "/" + companyFiles.Name())
			if err != nil {
				log.Println(companyFiles.Name())
				log.Fatalln(err)
			}

			err = json.Unmarshal(file, &company)
			if err != nil {
				log.Println(companyFiles.Name())
				log.Fatalln(err)
			}
			err = company.Validate()
			if err != nil {
				log.Println(companyFiles.Name())
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
