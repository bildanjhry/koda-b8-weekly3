package models

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Items struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Stocks   int      `json:"stocks"`
	Category Category `json:"cat"`
	Size     string   `json:"size"`
	IsPromo  int      `json:"isPromo"`
}

type ItemsCheckout struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Category Category `json:"cat"`
	Qty      int      `json:"qty"`
	Size     string   `json:"size"`
	IsPromo  int      `json:"isPromo"`
}

type Cart struct {
	Products []ItemsCheckout `json:"products"`
	Total    int             `json:"total"`
}

type Checkout struct {
	Total         int     `json:"total"`
	PaymentStatus string  `json:"paymentStatus"`
	PaymentMethod string  `json:"paymentMethod"`
	Products      []Items `json:"products"`
}
