package main

import "fmt"

type Product struct {
	Name string
	Price float64
}

func ApplyDiscount(p *Product, discountPercentage float64) {
	p.Price = p.Price - (p.Price * (discountPercentage / 100))
}

func main() {
	products := []Product{
		{Name: "Laptop", Price: 1000.00},
		{Name: "Smartphone", Price: 500.00},
		{Name: "Tablet", Price: 300.00},
	}

	products = append(products, Product{Name: "Headphones", Price: 150.00})

	fmt.Println("====== DAFTAR PRODUK ======")

	for index, item := range products {
		fmt.Printf("%d. %s - $%.2f\n", index+1, item.Name, item.Price)
	}

	fmt.Println("====== DiSKON 10 % ======")

	ApplyDiscount(&products[0], 10)

	fmt.Printf("Harga baru %s seletah diskon: %.2f\n", products[0].Name, products[0].Price)

	fmt.Println("\n======== Looping tanpa index ========")

	for _, item := range products {
		fmt.Printf("%s - $%.2f\n", item.Name, item.Price)
	}
}