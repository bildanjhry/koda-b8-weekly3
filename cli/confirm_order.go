package cli

import (
	"encoding/json"
	"fmt"
	"mcd-clone/models"
	"mcd-clone/services"
	"mcd-clone/ui"
	"mcd-clone/utils"
	"os"
)

func ConfirmOrder(items *models.Items, ItemList *[]models.Items) {

	myCart, err := services.GetDataCart()
	dis := ui.Display{}
	if err != nil {
		fmt.Println(err)
	}

	form := models.ItemsCheckout{
		Id:       items.Id,
		Name:     items.Name,
		Price:    items.Price,
		Category: items.Category,
		Qty:      1,
		Size:     items.Size,
		IsPromo:  items.IsPromo,
	}

	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	// oldVal := myCart.Products[:0]
	// for _, val := range myCart.Products {
	// 	if val.Id == items.Id {
	// 		val.Qty += 1
	// 		myCart.Products = append(oldVal, val)
	// 		break
	// 	}
	// }

	myCart.Products = append(myCart.Products, form)
	myCart.Total += items.Price

	data, err := json.MarshalIndent(myCart, "", " ")
	if err != nil {
		panic(err.Error())
	}

	err = os.WriteFile("./data/cart.json", data, 0644)
	if err != nil {
		panic(err.Error())
	}

	dis.OrderList(&myCart.Products)
	res, _ := utils.Io("\nMasukan Input: ")
	switch res {
	case "y":
		utils.ClearTerm(0, "")
		AskingItems(ItemList)
	case "n":
		utils.ClearTerm(0, "")
		HomeMenu(1)
	default:
		utils.ClearTerm(1, "* INVALID_INPUT *")
		ConfirmOrder(items, ItemList)
	}
}
