package cli

import (
	"mcd-clone/models"
	"mcd-clone/services"
	"mcd-clone/ui"
	"mcd-clone/utils"
	"sync"
)

func Checkout() {
	wg := sync.WaitGroup{}
	cart := models.Cart{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		res, _ := services.GetDataCart()
		cart = res
	}()
	dis := ui.Display{}
	wg.Wait()
	for {
		dis.Checkout(&cart)
		val, _ := utils.Io("\nMasukan input: ")
		switch val {
		case "1":
			utils.ClearTerm(0, "")
			AskingPayment()
		case "2":
			utils.ClearTerm(0, "")
			DecreastItem(&cart)
		case "3":
			utils.ClearTerm(0, "")
			HomeMenu(1)
		default:
			utils.ClearTerm(1, "* INVALID_INPUT *")
		}

	}
}
