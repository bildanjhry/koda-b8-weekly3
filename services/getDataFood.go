package services

import (
	"encoding/json"
	"fmt"
	"mcd-clone/models"
	"os"
)

func GetDataFood() []models.Items {
	foods := []models.Items{}
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
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
