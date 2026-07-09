package cli

import (
	"fmt"
	"mcd-clone/ui"
	"mcd-clone/utils"
	"os"
)

func SuccessOrder() {
	dis := ui.Display{}
	dis.Success()

	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	err := os.Remove("./data/cart.json")
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("\n1. Home\n2. Keluar\n")
	val, _ := utils.Io("\nMasukan Input: ")

	switch val {
	case "1":
		utils.ClearTerm(0, "")
		HomeMenu(0)
	case "2":
		utils.ClearTerm(0, "")
		os.Exit(0)
	}
}
