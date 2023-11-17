package main

import (
	"fmt"
	"os"

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
		fmt.Printf("An error ocurred: %v\n", err.Error())
		os.Exit(1)
	}

	answer, err := calculator.Calculate(operandA, operandB, operation)
	if err != nil {
		fmt.Printf("An error ocurred: %v\n", err.Error())
		os.Exit(1)
	}

	if !isRoman {
		fmt.Println(answer)
		return
	}

	if answer < 1 {
		fmt.Printf("Cannot have non-positive result when using roman numerals\n")
		os.Exit(1)
	}
	answerAsString, err := converter.ConvertToRoman(answer)
	if err != nil {
		fmt.Printf("An error ocurred: %v\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(answerAsString)
}
