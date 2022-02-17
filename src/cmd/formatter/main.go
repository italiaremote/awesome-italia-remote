package main

import (
	"encoding/json"
	"github.com/italiaremote/awesome-italia-remote/pkg/cmp"
	"log"
	"os"
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

			jsonByte, err := json.MarshalIndent(company, "", "  ")
			if err != nil {
				log.Fatalln(err)
			}

			err = os.WriteFile(fileName, jsonByte, 0644)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
