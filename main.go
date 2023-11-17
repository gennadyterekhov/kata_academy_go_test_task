package main

import (
	"fmt"
)

func main() {
	runCalculator()
}

func runCalculator() {
	converter := CreateConverter()
	cliReader := CreateCliReader(&converter)
	operandA, operandB, operation, isRoman, err := cliReader.GetInput()
	if err != nil {
		panic(err)
	}

	calculator := CreateCalculator()
	answer, err := calculator.Calculate(operandA, operandB, operation)
	if err != nil {
		panic(err)
	}

	if !isRoman {
		fmt.Println(answer)
		return
	}

	if answer < 1 {
		panic("Cannot have non-positive result when using roman numerals")
	}
	answerAsString, err := converter.ConvertToRoman(answer)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(answerAsString)
}
