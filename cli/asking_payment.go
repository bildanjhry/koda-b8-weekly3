package cli

import (
	"mcd-clone/ui"
	"mcd-clone/utils"
	"time"
)

func AskingPayment() {
	for {
		dis := ui.Display{}
		dis.Payment()
		val, _ := utils.Io("\nMasukan Input: ")
		if val == "1" || val == "2" {
			utils.ClearTerm(1, "Prosess ..")
			time.Sleep(2 * time.Second)
			utils.ClearTerm(0, "")
			SuccessOrder(&val)
		} else {
			utils.ClearTerm(1, "* INVALID_INPUT *")
		}
	}

}
