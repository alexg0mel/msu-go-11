// Go рутины, одновременно работающие с общими данными сами собой не могу синхронизироваться
// Как пример защиты данных от небезопасных изменений, можно использовать структуру Mutex
// У нее также нет публичных полей, но, есть два публичных метода
// Lock(), который даст только одной подпрограмме продолжить выполнение блока, остальне заблокируются в ожидании
// Unlock(), который снимает лок, захваченный ранее
package main

import (
	"fmt"
	"log"
	"sync"
)

// Пусть у нас есть Счет
// Встроим в него объект Mutex
// теперь наш объект может использовать его публичные методы
type AccountProtected struct {
	sync.Mutex
	balance float64
}

func (a *AccountProtected) Balance() float64 {
	a.Lock()
	defer a.Unlock()
	return a.balance
}

func (a *AccountProtected) Deposit(amount float64) {
	a.Lock()
	defer a.Unlock()
	log.Printf("depositing: %f", amount)
	a.balance += amount
}

func (a *AccountProtected) Withdraw(amount float64) {
	a.Lock()
	defer a.Unlock()
	if amount > a.balance {
		return
	}
	log.Printf("withdrawing: %f", amount)
	a.balance -= amount
}

func main() {
	acc := AccountProtected{}

	// Стартуем 10 go рутин
	for i := 0; i < 10; i++ {
		go func() {
			// Каждая из которых, производит операции с аккаунтом
			for j := 0; j < 10; j++ {
				// Иногда снимает деньги
				if j%2 == 1 {
					acc.Withdraw(50)
					continue
				}
				// иногда кладет
				acc.Deposit(50)
			}
		}()
	}
	fmt.Scanln()
	// Теперь баланс всегда будет сходиться в 0
	fmt.Println(acc.Balance())

}
