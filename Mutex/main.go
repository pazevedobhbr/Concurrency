package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// Variable for bank balance
	var bankBalance int
	var balance sync.Mutex

	//print out starting values
	fmt.Printf("Starting bank balance: %d.00", bankBalance)
	fmt.Println()
	//define weekly revenue
	incomes := []Income{
		{"Main Job", 500},
		{"Gifts", 10},
		{"Par Time Job", 50},
		{"Investments", 100},
	}

	wg.Add(len(incomes))
	//loop through 52 weeks and print how much is made; keep a running total
	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("Week %d: you earned $$%d.00 from %s  \n", week, income.Amount, income.Source)

			}
		}(i, income)
	}
	wg.Wait()
	//print out final balance
	fmt.Printf("Final bank balance: %d.00", bankBalance)
	fmt.Println()
}
