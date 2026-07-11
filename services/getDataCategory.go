package services

import (
	"encoding/json"
	"fmt"
	"mcd-clone/models"
	"os"
)

func GetDataCategory() []models.Category {
	categories := []models.Category{}
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
			os.Stdout.Close()
		}
	}()

	file, err := os.ReadFile("./data/categories.json")
	if err != nil {
		panic(err.Error())
	}
	errMar := json.Unmarshal([]byte(string(file)), &categories)
	if errMar != nil {
		panic(errMar.Error())
	}

	return categories
}
