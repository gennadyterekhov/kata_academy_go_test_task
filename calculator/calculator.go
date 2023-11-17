package calculator

import "errors"

type Calculator struct {
}

type Operation struct {
}

func CreateCalculator() *Calculator {
	return &Calculator{}
}

func Calculate(operandA int, operandB int, operation string) (int, error) {

	if operation == "+" {
		return operandA + operandB, nil
	}
	if operation == "-" {
		return operandA - operandB, nil
	}
	if operation == "*" {
		return operandA * operandB, nil
	}
	if operation == "/" {
		return operandA / operandB, nil
	}
	return -1, errors.New("unknown operation")
}
