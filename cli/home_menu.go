package cli

import (
	"mcd-clone/models"
	"mcd-clone/services"
	"mcd-clone/ui"
	"mcd-clone/utils"
	"os"
	"sync"
)

func AskingHome(past int) string {
	dis := ui.Display{}
	dis.HomeList(past)
	res, _ := utils.Io("\nMasukan Input: ")
	return res
}

func HomeMenu(past int) {
	var (
		foods    []models.Items
		drinks   []models.Items
		snacks   []models.Items
		desserts []models.Items
	)
	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {
		defer wg.Done()
		foods = services.GetDataFood()
	}()

	go func() {
		defer wg.Done()
		drinks = services.GetDataDrink()
	}()

	go func() {
		defer wg.Done()
		snacks = services.GetDataSnack()
	}()

	go func() {
		defer wg.Done()
		desserts = services.GetDataDessert()
	}()

	for {
		input := AskingHome(past)
		wg.Wait()
		switch input {
		case "1":
			utils.ClearTerm(0, "")
			AskingItems(&foods)

		case "2":
			utils.ClearTerm(0, "")
			AskingItems(&drinks)

		case "3":
			utils.ClearTerm(0, "")
			AskingItems(&snacks)

		case "4":
			utils.ClearTerm(0, "")
			AskingItems(&desserts)

		case "5":
			utils.ClearTerm(1, "* WILL_AVAILABLE_SOON *")

		case "6":
			utils.ClearTerm(1, "* WILL_AVAILABLE_SOON *")

		case "7":
			utils.ClearTerm(0, "")
			Checkout()

		case "x":
			utils.ClearTerm(0, "Sampai Jumpa")
			os.Exit(0)

		default:
			utils.ClearTerm(1, "* INVALID_INPUT *")
		}

	}

}
