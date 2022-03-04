package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/italiaremote/awesome-italia-remote/pkg/cmp"
)

const homePath = "../../../"
const dataPath = homePath + "data/"

func main() {
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
			fileName := dataPath + section + "/" + companyFiles.Name()
			file, err := os.ReadFile(fileName)
			if err != nil {
				log.Fatalln(err)
			}
			var company cmp.Company

			err = json.Unmarshal(file, &company)
			if err != nil {
				log.Fatalln(err)
			}

			newFileName := dataPath + companyFiles.Name()
			newFile, err := os.ReadFile(newFileName)
			if err != nil {
				log.Println("File not found: ", newFileName)
				company.Categories = []string{
					section,
				}
			} else {
				var newCompany cmp.Company
				err = json.Unmarshal(newFile, &newCompany)
				if err != nil {
					log.Fatalln(err)
				}
				company.Categories = append(newCompany.Categories, section)
				if newCompany.Tags != nil {
					company.Tags = append(company.Tags, newCompany.Tags...)
				}
				if newCompany.HiringPolicy != "" {
					company.HiringPolicy = newCompany.HiringPolicy
				}
				if newCompany.RemotePolicy != "" {
					company.RemotePolicy = newCompany.RemotePolicy
				}
				if newCompany.Type != "" {
					company.Type = newCompany.Type
				}
			}

			jsonByte, err := json.MarshalIndent(company, "", "  ")
			if err != nil {
				log.Fatalln(err)
			}

			err = os.WriteFile(dataPath+companyFiles.Name(), jsonByte, 0644)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
