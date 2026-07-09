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
	fmt.Printf("\n1. Checkout")
	val, _ := utils.Io("\nMasukan input: ")
	switch val {
	case "1":
		AskingPayment()
	}
}
