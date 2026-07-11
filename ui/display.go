package ui

import (
	"fmt"
	"mcd-clone/models"
	"mcd-clone/services"
	"sync"
)

type Display struct{}

func (u Display) HomeList(Status int) {
	wg := sync.WaitGroup{}
	var thisCart = []models.ItemsCheckout{}
	var thisCategories = []models.Category{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		cart, _ := services.GetDataCart()
		thisCart = cart.Products
	}()

	go func() {
		defer wg.Done()
		thisCategories = services.GetDataCategory()
	}()

	wg.Wait()
	if x := Status; x == 0 {
		fmt.Println("=======================================")
		fmt.Println("     SELAMAT DATANG DI MCDONALDS")
		fmt.Println("=======================================")
	} else {
		fmt.Println("=======================================")
		fmt.Printf("Home Menu :\n")
	}
	for x, val := range thisCategories {
		fmt.Printf("\n(%d) %s", x+1, val.Name)
	}
	fmt.Printf("\n\n")
	fmt.Println("=======================================")
	if x := Status; x == 1 || len(thisCart) > 0 {
		fmt.Printf("(7) Checkout\n")
	} else {
		fmt.Printf("(X) Keluar\n")
	}

}

func (u Display) ItemList(Items *[]models.Items) {
	fmt.Printf("Silahkan Pilih Item :\n\n")
	for x, val := range *Items {
		fmt.Printf("(%d) %s", x+1, val.Name)
		if val.Size != "" {
			fmt.Printf(", %s ", val.Size)
		}
		fmt.Printf("\n    Rp%d\n\n", val.Price)
	}

	fmt.Println("=======================================")
	fmt.Printf("(0) Kembali\n")
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
		fmt.Printf("  %d. %s %dx", x+1, val.Name, val.Qty)
		if val.Size != "" {
			fmt.Printf(", %s ", val.Size)
		}
		fmt.Printf("\n     Rp%d\n", val.Price)
	}
	fmt.Printf("\nTotal: Rp.%d\n\n", Items.Total)

	fmt.Println("=======================================")
	fmt.Printf("(1) Checkout    (2) Hapus    (3) Kembali\n")
}

func (u Display) DecItem(Items *[]models.ItemsCheckout) {
	fmt.Printf("Keranjang :\n\n")
	for x, val := range *Items {
		fmt.Printf("  %d. %s %dx", x+1, val.Name, val.Qty)
		if val.Size != "" {
			fmt.Printf(", %s ", val.Size)
		}
		fmt.Printf("\n     Rp%d\n", val.Price)
	}
	fmt.Println("=======================================")
	fmt.Printf("Pilih yang ingin dikurangi\n")
}

func (U Display) Success(Items *models.Cart) {
	fmt.Printf("                McDonalds\n")
	fmt.Printf("       Bogor, Jawa Barat Indonesia\n")
	fmt.Printf("                021183182\n")
	fmt.Printf("---------------------------------------\n\n")
	for x, val := range Items.Products {
		fmt.Printf("%d. %s %dx", x+1, val.Name, val.Qty)
		if val.Size != "" {
			fmt.Printf(", %s ", val.Size)
		}
		fmt.Printf("\n   Rp%d\n", val.Price)
	}
	fmt.Printf("\n\nTotal: Rp.%d\n\n", Items.Total)
	fmt.Printf("---------------------------------------\n")
	fmt.Printf("Payment Method                   QRIS\n")
	fmt.Printf("Payment Status                   Paid\n")
	fmt.Printf("---------------------------------------\n")
	fmt.Printf("\n    Silahkan tunggu pesanan anda.\n")
	fmt.Printf("             Terima Kasih\n\n")
	fmt.Println("=======================================")
	fmt.Printf("\n(1) Home  (2) Keluar\n")

}

func (U Display) Struct(Items *models.Cart) {
	fmt.Printf("                McDonalds\n")
	fmt.Printf("       Bogor, Jawa Barat Indonesia\n")
	fmt.Printf("                021183182\n")
	fmt.Printf("---------------------------------------\n\n")
	for x, val := range Items.Products {
		fmt.Printf("%d. %s %dx", x+1, val.Name, val.Qty)
		if val.Size != "" {
			fmt.Printf(", %s ", val.Size)
		}
		fmt.Printf("\n   Rp%d\n", val.Price)
	}
	fmt.Printf("\n\nTotal: Rp.%d\n\n", Items.Total)
	fmt.Printf("---------------------------------------\n")
	fmt.Printf("Payment Method                   Cash\n")
	fmt.Printf("Payment Status                   Unpaid\n")
	fmt.Printf("---------------------------------------\n")
	fmt.Printf("\n    Berikan Struk ini pada kasir.\n")
	fmt.Printf("             Terima Kasih\n\n")
	fmt.Println("=======================================")
	fmt.Printf("\n(1) Home  (2) Keluar\n")
}

func (u Display) Payment() {
	fmt.Printf("Pilih Metode Pembayaran:\n\n")
	fmt.Printf("(1) Cash\n(2) QRIS\n")
	fmt.Printf("\n================================\n")
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
