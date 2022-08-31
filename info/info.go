package info

import "fmt"

const mainTitle = "BMI Calculator"
const separator = "=============="

const WeightPrompt = "Please Enter Your Weight(kg): "
const HeightPrompt = "Please Enter Your Height(cm): "

func PrintWelcome() {
	fmt.Println(mainTitle)
	fmt.Println(separator)
}
