package services

import (
	"encoding/json"
	"fmt"
	"mcd-clone/models"
	"os"
)

func GetDataDessert() []models.Items {
	desserts := []models.Items{}
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
			os.Stdout.Close()
		}
	}()

	file, err := os.ReadFile("./data/desserts.json")
	if err != nil {
		panic(err.Error())
	}
	errMar := json.Unmarshal([]byte(string(file)), &desserts)
	if errMar != nil {
		panic(errMar.Error())
	}

	return desserts
}
