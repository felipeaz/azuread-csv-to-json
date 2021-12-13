package service

import (
	"azuread-csv-to-json/internal/app/model"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type UserReader struct {
	FilePath string
}

func NewUserReader(path string) UserReader {
	return UserReader{
		FilePath: path,
	}
}

func (r UserReader) CreateJSONUsers() {
	users := r.getUserFromCSV()
	r.writeJSONFile(users)
	return
}

func (r UserReader) getUserFromCSV() []model.User {
	f, err := os.Open("./usr.csv")
	if err != nil {
		log.Fatal("Unable to open file:", err.Error())
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+r.FilePath, err)
	}

	var users []model.User
	records = records[1:]
	for _, data := range records {
		usr := model.User{
			Id:         data[5],
			FirstName:  data[4],
			LastName:   data[2],
			Email:      data[0],
			Phone:      data[18],
			AdProvider: "Azure Active Directory",
			Access:     nil,
		}
		users = append(users, usr)
	}
	return users
}

func (r UserReader) getExportFile(path string) (f *os.File) {
	filePath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal("failed to retrieve export file", err.Error())
		return
	}

	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		err = os.Mkdir(filePath, 0755)
		if err != nil {
			log.Fatal("failed to create export folder", err.Error())
			return
		}
	}

	fileName := "users.json"
	fullPath := fmt.Sprintf("%s/%s", filePath, fileName)

	f, err = os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("failed to create export file", err.Error())
		return nil
	}

	return
}

func (r UserReader) writeJSONFile(users []model.User) {
	f := r.getExportFile("./")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("failed to get export file")
		}
	}(f)

	b, e := json.Marshal(users)
	if e != nil {
		log.Fatal("failed to marshal users", e.Error())
	}

	_, e = f.Write(b)
	if e != nil {
		log.Fatal("failed to write users", e.Error())
	}
}
