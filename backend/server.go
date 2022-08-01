package backend

import (
	"io"
	"log"
	"net/http"
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

func getData(w http.ResponseWriter, r *http.Request) {
	dataFile, err := os.Open(dataFilePath)
	if err != nil {
		log.Fatal("file open on get", err.Error())
		w.WriteHeader(500)
		io.WriteString(w, "error reading file")
		return
	}

	defer dataFile.Close()

	_, err = io.Copy(dataFile, r.Body)
	if err != nil {
		log.Print("copy from request: ", err.Error())
		w.WriteHeader(500)
		return
	}
}
