package main

import "fmt"

// 1. DEFINISI INTERFACE (Kontrak Payment Gateway)
type PaymentProcessor interface {
	Pay(amount float64) string
}

// 2. STRUCT PERTAMA: Midtrans (Misal untuk e-wallet/QRIS)
type Midtrans struct {
	MerchantID string
}

// Method Pay() untuk Midtrans
func (m Midtrans) Pay(amount float64) string {
	return fmt.Sprintf("[Midtrans %s] Memproses pembayaran sebesar Rp%.2f", m.MerchantID, amount)
}

// 3. STRUCT KEDUA: Stripe (Misal untuk kartu kredit)
type Stripe struct {
	ApiKey string
}

// Method Pay() untuk Stripe
func (s Stripe) Pay(amount float64) string {
	return fmt.Sprintf("[Stripe API] Memproses pembayaran sebesar $%.2f", amount)
}

// 4. FUNGSI YANG MENERIMA INTERFACE
// Fungsi ini tidak peduli payment-nya pakai Midtrans atau Stripe, 
// asalkan punya method Pay(float64) string!
func ProcessOrder(p PaymentProcessor, price float64) {
	result := p.Pay(price)
	fmt.Println("LOG TRANSACTION:", result)
}

func main() {
	midtransGate := Midtrans{MerchantID: "M-12345"}
	stripeGate := Stripe{ApiKey: "sk_test_9999"}

	fmt.Println("=== DEMO INTERFACE DI GO ===")

	// Keduanya bisa dimasukkan ke fungsi ProcessOrder!
	ProcessOrder(midtransGate, 150000)
	ProcessOrder(stripeGate, 25.50)
}