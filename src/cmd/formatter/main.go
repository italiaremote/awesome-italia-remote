package main

import (
	"encoding/json"
	"log"
	"os"
	"regexp"

	"github.com/italiaremote/awesome-italia-remote/pkg/cmp"
)

const homePath = "../../../"
const dataPath = homePath + "data/"

func main() {
	companiesFile, err := os.ReadDir(dataPath)
	if err != nil {
		log.Fatalln(err)
	}
	fileNameRegex := regexp.MustCompile(`([a-zA-Z0-9\_\-]+).json$`)

	for _, companyFiles := range companiesFile {
		fileName := dataPath + companyFiles.Name()
		file, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatalln(err)
		}
		var company cmp.Company

		if !fileNameRegex.MatchString(fileName) {
			log.Fatalln("Invalid filename, use only letters. numbers and - _")
		}

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
