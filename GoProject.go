/*
In this application we are going to create an investment calculator that calculates the future value of an investment based on the principal amount, the interest rate, and the time period.
Also it calculates the real future value according to the inflation rate of the year...
*/
package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	// fmt.Print("The investment amount is : ")
	outputText("The investment amount is :")
	fmt.Scan(&investmentAmount)

	//fmt.Print("The expected return rate is : ")
	outputText("The expected return rate is :")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("The number of years is : ")
	outputText("The number of years is :")
	fmt.Scan(&years)
	futureValue, futureRealValue := CalculateFutureValues(investmentAmount, expectedReturnRate, years)
	//futureValue := investmentAmount * math.Pow((1 + expectedReturnRate), years)
	//futureRealValue := futureValue / math.Pow((1 + inflationRate), years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (Adjusted for Inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func CalculateFutureValues(investmentAmount, expectedReturnRate, years float64) (FV float64, RFV float64) {
	FV = investmentAmount * math.Pow((1+expectedReturnRate), years)
	RFV = FV / math.Pow((1+inflationRate), years)
	return FV, RFV
	// return
}
