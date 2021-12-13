package main

import (
	"azuread-csv-to-json/internal/app/service"
)

func main() {
	reader := service.NewUserReader("./usr.csv")
	reader.CreateJSONUsers()
}
