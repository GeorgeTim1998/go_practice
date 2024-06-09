// Паттерн "адаптер" (Adapter) используется для преобразования интерфейса одного класса в интерфейс,
// который ожидает клиент. Адаптер позволяет классам работать вместе, что невозможно из-за несовместимых интерфейсов.

package main

import "fmt"

// Интерфейс, ожидаемый клиентом (функцией makePaymentWithAdapter)
type PaymentProcessor interface {
	Pay(amount float64)
}

// Старая платежная система (Legacy) и ее метод
type LegacyPaymentSystem struct{}

func (*LegacyPaymentSystem) ProcessOldPayment(amount float64) {
	fmt.Printf("Processed payment of %.2f using Legacy Payment System\n", amount)
}

// Адаптер для старой платежной системы
type LegacyPaymentAdapter struct {
	legacySystem *LegacyPaymentSystem
}

func (adapter *LegacyPaymentAdapter) Pay(amount float64) {
	// Дополнительные действия, если нужны, чтобы заставить legacy работать так как нужно
	fmt.Println("Performing additional actions before processing payment...")

	// Вызов метода старой платежной системы
	adapter.legacySystem.ProcessOldPayment(amount)
}

// клиент, с помощью которого производится оплата legacy системой
func makePaymentWithAdapter(t PaymentProcessor, amount float64) {
	t.Pay(amount)
}

func main() {
	// Создание объекта legacy платежной системы
	legacySystem := &LegacyPaymentSystem{}

	// Создание адаптера для legacy платежной системы
	adapter := &LegacyPaymentAdapter{legacySystem: legacySystem}

	// Использование адаптера для оплаты старой платежной системы типа LegacyPaymentSystem
	makePaymentWithAdapter(adapter, 50.0)
}
