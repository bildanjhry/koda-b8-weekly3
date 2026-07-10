package services

import (
	"encoding/json"
	"mcd-clone/models"
	"os"
)

func GetDataFood() []models.Items {
	foods := []models.Items{}
	defer func() {
		if x := recover(); x != nil {
			os.Stdout.Close()
		}
	}()

	file, err := os.ReadFile("./data/foods.json")
	if err != nil {
		panic(err.Error())
	}
	errMar := json.Unmarshal([]byte(string(file)), &foods)
	if errMar != nil {
		panic(errMar.Error())
	}

	return foods
}
