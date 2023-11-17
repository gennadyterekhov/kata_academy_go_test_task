package main

import (
	"fmt"

	"github.com/gennadyterekhov/kata_academy_go_test_task/calculator"
	cliReader "github.com/gennadyterekhov/kata_academy_go_test_task/cli_reader"
	"github.com/gennadyterekhov/kata_academy_go_test_task/converter"
)

func main() {
	runCalculator()
}

func runCalculator() {
	operandA, operandB, operation, isRoman, err := cliReader.GetInput()
	if err != nil {
		panic(err)
	}

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
