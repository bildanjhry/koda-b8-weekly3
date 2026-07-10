package cli

import (
	"mcd-clone/services"
	"mcd-clone/ui"
	"mcd-clone/utils"
	"os"
)

func AskingHome(past int) string {
	dis := ui.Display{}
	dis.HomeList(past)
	res, _ := utils.Io("\nMasukan Input: ")
	return res
}

func HomeMenu(past int) {
	switch input := AskingHome(past); input {
	case "1":
		utils.ClearTerm(0, "")
		foods := services.GetDataFood()
		AskingItems(&foods)
	case "7":
		utils.ClearTerm(0, "")
		Checkout()
	case "x":
		utils.ClearTerm(0, "Sampai Jumpa")
		os.Exit(0)
	default:
		utils.ClearTerm(1, "* WORKING_ON_PROGRESS *")
		HomeMenu(0)
	}
}
