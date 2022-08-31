package main

import (
	"BMI-calculator/info"
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (weight float64, height float64) {

	weight = getUserInput(info.WeightPrompt)
	height = getUserInput(info.HeightPrompt)

	return
}

func getUserInput(prompt string) (value float64) {
	fmt.Print(prompt)
	userInput, _ := reader.ReadString('\n')

	if runtime.GOOS == "windows" {
		userInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		userInput = strings.Replace(userInput, "\n", "", -1)
	}

	value, _ = strconv.ParseFloat(userInput, 64)
	return
}
