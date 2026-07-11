package cli

import (
	"fmt"
	"mcd-clone/models"
	"mcd-clone/services"
	"mcd-clone/ui"
	"mcd-clone/utils"
	"os"
	"sync"
)

func SuccessOrder(value *string) {
	dis := ui.Display{}
	wg := sync.WaitGroup{}
	data := models.Cart{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		res, _ := services.GetDataCart()
		data = res
	}()

	wg.Wait()
	if *value == "1" {
		dis.Struct(&data)
	} else {
		dis.Success(&data)
	}

	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	err := os.Remove("./data/cart.json")
	if err != nil {
		panic(err.Error())
	}

	for {
		val, _ := utils.Io("\nMasukan Input: ")
		switch val {
		case "1":
			utils.ClearTerm(0, "")
			HomeMenu(0)
		case "2":
			utils.ClearTerm(0, "")
			os.Exit(0)
		default:
			utils.ClearTerm(1, "* INVALID_INPUT *")
		}
	}
}
