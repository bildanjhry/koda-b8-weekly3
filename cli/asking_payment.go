package cli

import (
	"fmt"
	"mcd-clone/utils"
)

func AskingPayment() {
	fmt.Printf("Pilih Metode Pembayaran:\n\n")
	fmt.Printf("1. Cash\n")
	val, _ := utils.Io("\nMasukan Input: ")
	if val == "1" {
		utils.ClearTerm(0, "")
		SuccessOrder()
	}
}
