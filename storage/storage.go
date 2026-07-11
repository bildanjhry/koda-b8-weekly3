package storage

import "os"

func init() {
	if _, err := os.Stat("cart.json"); os.IsNotExist(err) {
		os.WriteFile("cart.json", []byte("[]"), 0644)
	}
}
