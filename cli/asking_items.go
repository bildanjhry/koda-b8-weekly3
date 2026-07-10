package cli

import (
	"fmt"
	"mcd-clone/models"
	"mcd-clone/ui"
	"mcd-clone/utils"
	"strconv"
)

func AskingItems(Items *[]models.Items) {
	dis := ui.Display{}

	fmt.Printf("Silahkan Pilih Item :\n\n")
	dis.ItemList(Items)
	res, _ := utils.Io("\nMasukan Input: ")

	for x, val := range *Items {
		if it := x + 1; strconv.Itoa(it) == res {
			utils.ClearTerm(0, "")
			ConfirmOrder(&val, Items)
		}
	}

	switch res {
	case "x":
		utils.ClearTerm(0, "")
		HomeMenu(0)
	case "0":
		utils.ClearTerm(0, "")
		HomeMenu(0)
	default:
		utils.ClearTerm(1, "* INVALID_INPUT *")
		AskingItems(Items)
	}

}
