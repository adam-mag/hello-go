package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 1. Struct untuk Data Model
// `json:"..."` disebut Struct Tag, berguna untuk mapping key di response JSON
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Struct untuk Response JSON standar
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Database dummy di memory
var products = []Product{
	{ID: 1, Name: "Kopi Hitam", Price: 15000},
	{ID: 2, Name: "Roti Bakar", Price: 20000},
}

// 2. Handler function untuk endpoint /products
func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Set Header agar response dibaca sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Guard Clause: Hanya izinkan method GET
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "Method tidak diizinkan! Gunakan GET.",
		})
		return
	}

	// Set status OK (200) dan return data JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "Berhasil mengambil data produk",
		Data:    products,
	})
}

func main() {
	// 3. Routing Endpoint
	http.HandleFunc("/api/products", getProductsHandler)

	fmt.Println("Server berjalan di http://localhost:8080...")

	// 4. Jalankan HTTP Server di port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}