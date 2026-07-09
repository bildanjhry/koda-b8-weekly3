package cli

import (
	"encoding/json"
	"fmt"
	"mcd-clone/models"
	"mcd-clone/services"
	"mcd-clone/utils"
	"os"
)

func ConfirmOrder(items *models.Items, ItemList *[]models.Items) {

	myCart, err := services.GetDataCart()
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

	oldVal := myCart.Products[:0]
	for _, val := range myCart.Products {
		if val.Id == items.Id {
			val.Qty += 1
			myCart.Products = append(oldVal, val)
			break
		}
	}

	myCart.Products = append(myCart.Products, form)
	myCart.Total += items.Price

	data, err := json.MarshalIndent(myCart, "", " ")
	if err != nil {
		panic(err.Error())
	}

	err = os.WriteFile("data/cart.json", data, 0644)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Pesanan Anda:\n\n")
	for x, val := range myCart.Products {
		fmt.Printf("%d. %s", x+1, val.Name)
		if val.Size != "" {
			fmt.Printf(" (%s)", val.Size)
		}
		fmt.Printf(" %dx\n", val.Qty)
	}
	fmt.Printf("\n================================\n")
	fmt.Println("(Y/N) Ada lagi?")
	res, _ := utils.Io("\nMasukan Input: ")
	switch res {
	case "y":
		utils.ClearTerm(0, "")
		AskingItems(ItemList)
	case "n":
		utils.ClearTerm(0, "")
		HomeMenu(1)
	}
}
