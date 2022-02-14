package main

import (
	"bytes"
	"encoding/json"
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
			var buf bytes.Buffer
			err = json.Indent(&buf, file, "", "  ")
			if err != nil {
				log.Fatalln(err)
			}
			err = os.WriteFile(fileName, buf.Bytes(), 0644)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
