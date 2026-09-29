package main

import (
	"errors"
	"fmt"
)

// 1. FUNGSI DENGAN ERROR HANDLING
// Mengembalikan (float64, error)
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		// Mengembalikan error menggunakan errors.New()
		return 0, errors.New("tidak bisa membagi angka dengan nol (0)")
	}

	// Jika sukses, error diisi nil (artinya tidak ada error)
	return a / b, nil
}

// 2. DEMO DEFER
func ProcessDatabase() {
	fmt.Println("[1] Membuka koneksi ke Database...")

	// 'defer' akan ditahan dan DIJALANKAN PALING AKHIR
	// tepat sebelum fungsi ProcessDatabase() ini selesai.
	defer fmt.Println("[4] DEFER: Menutup koneksi Database (Cleanup!)")

	fmt.Println("[2] Membaca data dari Database...")
	fmt.Println("[3] Selesai memproses data!")
}

func main() {
	fmt.Println("=== DEMO DEFER ===")
	ProcessDatabase()

	fmt.Println("\n=== DEMO ERROR HANDLING ===")

	// Kasus 1: Pembagian Sukses
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error terjadi:", err)
	} else {
		fmt.Println("Hasil 10 / 2 =", result)
	}

	// Kasus 2: Pembagian Gagal (Bagi dengan 0)
	result2, err2 := Divide(10, 0)
	if err2 != nil {
		// Pola "if err != nil" ini adalah pola paling umum di Go!
		fmt.Println("Error terjadi:", err2)
	} else {
		fmt.Println("Hasil 10 / 0 =", result2)
	}
}