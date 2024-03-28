package utils

import (
	"fmt"
	"os"
)

func GetProfit() string {
	content, _ := os.ReadFile("profit")

	return string(content)
}


func StoreProfit(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f", ebt, profit, ratio)

	os.WriteFile("profit", []byte(results), 0644);
}
