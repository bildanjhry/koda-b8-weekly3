package cli

import (
	"encoding/json"
	"mcd-clone/models"
	"mcd-clone/ui"
	"mcd-clone/utils"
	"os"
	"strconv"
)

func DecreastItem(Items *models.Cart) {
	dis := ui.Display{}

	for {
		utils.ClearTerm(0, "")
		if len(Items.Products) < 1 {
			HomeMenu(0)
			return
		}
		dis.DecItem(&Items.Products)
		res, _ := utils.Io("\nMasukan Input :")
		if res == "0" {
			utils.ClearTerm(0, "")
			return
		}
		for x, val := range Items.Products {
			if it := x + 1; strconv.Itoa(it) == res {
				if val.Qty > 1 {
					Items.Products[x].Qty--
				} else {
					Items.Products = append(Items.Products[:x], Items.Products[x+1:]...)
				}
				Items.Total -= val.Price
			}
		}
		data, err := json.MarshalIndent(Items, "", " ")
		if err != nil {
			panic(err.Error())
		}

		err = os.WriteFile("./assets/data/cart.json", data, 0644)
		if err != nil {
			panic(err.Error())
		}
	}

}
