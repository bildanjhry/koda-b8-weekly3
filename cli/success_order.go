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
	fmt.Printf("\n1. Home\n2. Keluar\n")
	val, _ := utils.Io("\nMasukan Input: ")

	switch val {
	case "1":
		utils.ClearTerm(0, "")
		HomeMenu(0)
	case "2":
		utils.ClearTerm(1, "----------")
		os.Exit(0)
	}
}
