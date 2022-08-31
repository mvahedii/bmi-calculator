package main

import (
	"BMI-calculator/info"
)

func main() {
	info.PrintWelcome()

	weight, height := getUserMetrics()

	bmi := bmiCalculator(weight, height)

	printBmi(bmi)

}

func bmiCalculator(weight float64, height float64) float64 {
	return weight * 100 / (height * height)
}
