package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// 1. Struct Produk PPOB (Katalog)
type Product struct {
	Code  string  `json:"code"`  // Contoh: "PULSA10K", "PLN20K"
	Name  string  `json:"name"`  // Contoh: "Pulsa Telkomsel 10rb"
	Price float64 `json:"price"` // Contoh: 10500
}

// 2. Struct Transaksi PPOB
type Transaction struct {
	RefID        string `json:"ref_id"`        // Unique Reference ID (misal: "TRX-1001")
	ProductCode  string `json:"product_code"`  // Kode produk yang dibeli
	CustomerNo   string `json:"customer_no"`   // No HP / ID Pelanggan PLN
	Status       string `json:"status"`        // "PENDING", "SUCCESS", "FAILED"
	ResponseCode string `json:"response_code"` // "00" = Success, "68" = Pending, "99" = Failed
	SN           string `json:"sn,omitempty"`  // Serial Number dari vendor (misal Token PLN / SN Pulsa)
}

// 3. Struct Format Response JSON Standar
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// --- IN-MEMORY STORAGE & MUTEX (Step 2) ---
var (
	// Data Katalog Produk
	products = []Product{
		{Code: "PULSA10K", Name: "Pulsa Telkomsel 10rb", Price: 10500},
		{Code: "PLN20K", Name: "Token PLN 20rb", Price: 20500},
	}

	// Map penampung transaksi menggunakan RefID sebagai Key
	// Contoh: transactions["TRX-1001"] = Transaction{...}
	transactions = make(map[string]Transaction)

	// RWMutex untuk mengamankan pembacaan & penulisan memori
	storageMutex sync.RWMutex
)

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, `{"success": false, "message": "Method tidak diizinkan! Gunakan GET."}`)
		return
	}

	storageMutex.RLock() // Lock untuk membaca data produk
	defer storageMutex.RUnlock()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "Berhasil mengambil data produk",
		Data:    products,
	})
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprint(w, `{"success": false, "message": "Method tidak diizinkan! Gunakan POST."}`)
        return
    }

    // Decode request ke struct Product
    var newProd Product
    if err := json.NewDecoder(r.Body).Decode(&newProd); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(Response{Success: false, Message: "Format Request JSON tidak valid!"})
        return
    }

    if newProd.Code == "" || newProd.Name == "" || newProd.Price <= 0 {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(Response{Success: false, Message: "Code, Name, dan Price wajib diisi!"})
        return
    }

    // SIMPAN KE SLICE products (BUKAN transactions!)
    storageMutex.Lock()
    products = append(products, newProd)
    storageMutex.Unlock()

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(Response{
        Success: true,
        Message: "Produk berhasil ditambahkan!",
        Data:    newProd,
    })
}

func createTransactionHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprint(w, `{"success": false, "message": "Method tidak diizinkan! Gunakan POST."}`)
        return
    }

    var req struct {
        ProductCode string `json:"product_code"`
        CustomerNo  string `json:"customer_no"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(Response{Success: false, Message: "Format Request JSON tidak valid!"})
        return
    }

    if req.ProductCode == "" || req.CustomerNo == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(Response{Success: false, Message: "Product Code dan Customer No wajib diisi!"})
        return
    }

    refID := fmt.Sprintf("TRX-%d", len(transactions)+1001)

    newTRX := Transaction{
        RefID:        refID,
        ProductCode:  req.ProductCode,
        CustomerNo:   req.CustomerNo,
        Status:       "PENDING",
        ResponseCode: "68",
    }

    storageMutex.Lock()
    transactions[refID] = newTRX
    storageMutex.Unlock()

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(Response{
        Success: true,
        Message: "Transaksi berhasil dibuat! Status: PENDING",
        Data:    newTRX,
    })
}

func main() {
    // Endpoint Produk
    http.HandleFunc("/api/products", getProductsHandler)
    http.HandleFunc("/api/products/create", createProductHandler)

    // Endpoint Transaksi
    http.HandleFunc("/api/transactions/create", createTransactionHandler)

    fmt.Println("Server berjalan di http://localhost:8080...")
    http.ListenAndServe(":8080", nil)
}