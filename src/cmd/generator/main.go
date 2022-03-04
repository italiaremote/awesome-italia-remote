package main

import (
	"encoding/json"
	"log"
	"os"
	"text/template"

	"github.com/italiaremote/awesome-italia-remote/pkg/cmp"
)

const homePath = "../../../"
const dataPath = homePath + "data/"

func main() {
	companies := make(map[string][]cmp.Company)
	companiesFile, err := os.ReadDir(dataPath)
	if err != nil {
		log.Fatalln(err)
	}

	for _, companyFiles := range companiesFile {
		var company cmp.Company
		file, err := os.ReadFile(dataPath + companyFiles.Name())
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

		for _, category := range company.Categories {
			companies[category] = append(companies[category], company)
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
