package ui

import (
	"fmt"
	"mcd-clone/models"
)

type Display struct{}

func (u Display) HomeList(Status int) {
	if x := Status; x == 0 {
		fmt.Println("=======================================")
		fmt.Println("SELAMAT DATANG DI MC DONALDS")
		fmt.Println("=======================================")
	} else {
		fmt.Printf("Home Menu :\n")
	}
	fmt.Printf("\n1. Foods\n2. Drinks\n3. Snacks\n4. Dessert\n5. Paket\n6. Happy Meal\n\n")
	fmt.Println("=======================================")
	if x := Status; x == 1 {
		fmt.Printf("7. Checkout\n")
	} else {
		fmt.Printf("(X) Keluar\n")
	}
}

func (u Display) ItemList(Items *[]models.Items) {
	for x, val := range *Items {
		fmt.Printf("%d. %s", x+1, val.Name)
		if val.Size != "" {
			fmt.Printf(", %s ", val.Size)
		}
		fmt.Printf("\n   Rp%d\n\n", val.Price)
	}

	fmt.Println("=======================================")
	fmt.Printf("0. Kembali\n")
}

func (u Display) CartList(Carts *[]models.Cart) {
	fmt.Println("=======================================")
	fmt.Print("Pilihan anda:")
	for _, val := range *Carts {
		fmt.Print(val.Products)
	}

	fmt.Println("=======================================")

}

func (u Display) Checkout(Items *models.Cart) {
	fmt.Printf("Keranjang :\n\n")
	for x, val := range Items.Products {
		fmt.Printf("%d. %s %dx", x+1, val.Name, val.Qty)
		if val.Size != "" {
			fmt.Printf(", %s ", val.Size)
		}
		fmt.Printf("\n   Rp%d\n", val.Price)
	}
	fmt.Printf("\nTotal: Rp.%d\n\n", Items.Total)

	fmt.Println("=======================================")
}

func (U Display) Success() {
	fmt.Println("=======================================")
	fmt.Printf("\nTERIMAKASIH\n\n")
	fmt.Println("=======================================")
}

func (U Display) Struct(Items *models.Cart) {
	fmt.Println("=======================================")
	fmt.Printf("             McDonalds\n")
	fmt.Printf("=======================================\n\n")
	for x, val := range Items.Products {
		fmt.Printf("%d. %s %dx", x+1, val.Name, val.Qty)
		if val.Size != "" {
			fmt.Printf(", %s ", val.Size)
		}
		fmt.Printf("\n   Rp%d\n", val.Price)
	}
	fmt.Printf("\nTotal: Rp.%d\n\n", Items.Total)
	fmt.Println("=======================================")
	fmt.Printf("            TERIMAKASIH\n")
	fmt.Printf("=======================================\n\n")

}

func (u Display) OrderList(Items *[]models.ItemsCheckout) {
	fmt.Printf("Pesanan Anda:\n\n")
	for x, val := range *Items {
		fmt.Printf("%d. %s", x+1, val.Name)
		if val.Size != "" {
			fmt.Printf(", %s ", val.Size)
		}
		fmt.Printf(" %dx\n", val.Qty)
	}
	fmt.Printf("\n================================\n")
	fmt.Println("(Y/N) Ada lagi?")
}
