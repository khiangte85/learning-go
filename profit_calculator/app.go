package main

import (
	"fmt"
	"khiangte/profit_calculator/utils"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Println(utils.GetProfit())
	fmt.Println("--------------")

	fmt.Print("Revenue: ")
	_, err := fmt.Scan(&revenue)

	if err != nil {
		panic(err)
	}
	
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Printf("\nEBT: %.2f\n", revenue)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Profit Ratio: %.2f\n", ratio)

	utils.StoreProfit(ebt, profit, ratio)
}
