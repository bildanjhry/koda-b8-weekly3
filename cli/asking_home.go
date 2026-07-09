package cli

import (
	"mcd-clone/services"
	"mcd-clone/ui"
	"mcd-clone/utils"
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
	}
}
