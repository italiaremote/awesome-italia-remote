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
	companiesFile, err := os.ReadDir(dataPath)
	if err != nil {
		log.Fatalln(err)
	}

	for _, companyFiles := range companiesFile {
		fileName := dataPath + companyFiles.Name()
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
