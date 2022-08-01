package backend

import (
	"log"
	"os"
)

const dataFilePath = "./data/crm-app-out.txt"

func ensureDataFileExists() {
	_, err := os.Stat(dataFilePath)
	if os.IsNotExist(err) {
		dataFile, err := os.OpenFile(dataFilePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer dataFile.Close()
		_, err = dataFile.Write([]byte("[]"))

		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	}
}
