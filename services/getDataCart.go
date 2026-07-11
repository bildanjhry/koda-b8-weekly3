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

	file, _ := os.ReadFile("./cart.json")
	json.Unmarshal([]byte(string(file)), &cart)

	return cart, nil
}
