package ui

import "mcd-clone/models"

type Printing interface {
	HomeList(Status int)
	ItemList(Items *[]models.Items)
	CartList(Cart *[]models.Cart)
	Chekout(Items *models.Cart)
	OrderList(Items *[]models.ItemsCheckout)
	Success()
}
