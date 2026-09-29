package main

import (
	"fmt"
	"time"
)

// Fungsi yang mensimulasikan proses async (misal: panggil API / query DB)
func FetchAPI(source string, ch chan string) {
	fmt.Printf("[START] Mengambil data dari %s...\n", source)
	
	// Simulasi delay 2 detik
	time.Sleep(4 * time.Second)

	// Kirim hasil data ke Channel menggunakan operator '<-'
	ch <- fmt.Sprintf("Data dari %s BERHASIL diambil!", source)
}

func main() {
	startTime := time.Now()

	// 1. BUAT CHANNEL (Pipa pengirim string)
	ch := make(chan string)

	// 2. JALANKAN 3 PROCESS SECARA PARALEL (GOROUTINES)
	// Dengan kata kunci 'go', 3 fungsi ini jalan berbarengan di background!
	go FetchAPI("API Payment", ch)
	go FetchAPI("API User Profile", ch)
	go FetchAPI("API Product Catalog", ch)

	// 3. TERIMA DATA DARI CHANNEL
	// Menerima data dari channel bersifat BLOCKING (menunggu sampai ada data masuk)
	msg1 := <-ch
	msg2 := <-ch
	msg3 := <-ch

	fmt.Println("\n=== HASIL DITERIMA ===")
	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)

	// Cek total waktu eksekusi
	fmt.Printf("\nTotal waktu eksekusi: %v\n", time.Since(startTime))
}