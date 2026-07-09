package cli

import (
	"fmt"
	"mcd-clone/services"
	"mcd-clone/ui"
	"mcd-clone/utils"
)

func Checkout() {
	res, _ := services.GetDataCart()
	dis := ui.Display{}
	dis.Checkout(&res)
	fmt.Printf("1. Checkout\n")
	val, _ := utils.Io("\nMasukan input: ")
	switch val {
	case "1":
		utils.ClearTerm(0, "")
		AskingPayment()
	}
}
