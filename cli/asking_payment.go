package cli

import (
	"fmt"
	"mcd-clone/utils"
)

func AskingPayment() {

	for {
		fmt.Printf("Pilih Metode Pembayaran:\n\n")
		fmt.Printf("1. Cash\n2. QRIS")
		val, _ := utils.Io("\nMasukan Input: ")
		switch val {
		case "1":
			utils.ClearTerm(0, "")
			SuccessOrder(&val)
		case "2":
			utils.ClearTerm(0, "")
			SuccessOrder(&val)
		default:
			utils.ClearTerm(1, "* INVALID_INPUT *")
		}
	}

}
