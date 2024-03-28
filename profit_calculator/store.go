package main

import (
	"fmt"
	"os"
)



func storeProfit(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f", ebt, profit, ratio)

	os.WriteFile("profit", []byte(results), 0644);
}
