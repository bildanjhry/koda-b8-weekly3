package services

import (
	"encoding/json"
	"fmt"
	"mcd-clone/models"
	"os"
)

func GetDataCart() (models.Cart, error) {
	cart := models.Cart{}
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	file, err := os.ReadFile("./data/cart.json")
	if err != nil {
		return cart, fmt.Errorf("File tidak ditemukan")
	}
	errMar := json.Unmarshal([]byte(string(file)), &cart)
	if errMar != nil {
		panic(errMar.Error())
	}

	return cart, nil
}
