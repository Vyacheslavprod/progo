package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

type Stocklevel struct {
	Product
	Alternate Product
	count int
}

func main() {
	slice := []Stocklevel {
		{
			Product: Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count: 100,
		},
	}
	fmt.Println(slice[0].count)

	kvp := map[string]Stocklevel {
		"kayak": {
			Product: Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count: 100,
		},
	}
	fmt.Println(kvp["kayak"].Product.name)

}