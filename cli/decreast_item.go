package cli

import (
	"encoding/json"
	"mcd-clone/models"
	"mcd-clone/services"
	"mcd-clone/ui"
	"mcd-clone/utils"
	"os"
	"strconv"
	"sync"
)

func DecreastItem() {
	dis := ui.Display{}
	wg := sync.WaitGroup{}
	cart := models.Cart{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		cart, _ = services.GetDataCart()
	}()
	wg.Wait()

	for {
		utils.ClearTerm(0, "")
		if len(cart.Products) < 1 {
			HomeMenu(0)
			return
		}
		dis.DecItem(&cart.Products)
		res, _ := utils.Io("\nMasukan Input :")
		for x, val := range cart.Products {
			if it := x + 1; strconv.Itoa(it) == res {
				if val.Qty > 1 {
					cart.Products[x].Qty--
				} else {
					cart.Products = append(cart.Products[:x], cart.Products[x+1:]...)
				}
				cart.Total -= val.Price
			}
		}
		data, err := json.MarshalIndent(cart, "", " ")
		if err != nil {
			panic(err.Error())
		}

		err = os.WriteFile("./data/cart.json", data, 0644)
		if err != nil {
			panic(err.Error())
		}
	}

}
