package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var returnRate float64

	fmt.Print("Enter Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter years: ")
	fmt.Scan(&years)
	fmt.Print("Enter rate: ")
	fmt.Scan(&returnRate)

	futureAmount := investmentAmount * math.Pow(1+returnRate/100, years)
	futureRealAmount := futureAmount / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Return Amount: %.2f\n", futureAmount)
	fmt.Printf("Real return Amount: %.2f\n", futureRealAmount)
}
