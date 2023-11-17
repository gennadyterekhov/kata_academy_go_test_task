package cli_reader

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gennadyterekhov/kata_academy_go_test_task/converter"
)

func GetInput() (int, int, string, bool, error) {

	operandA, operandB, operation, isRoman, err := getCorrectDataFromLine()
	if err != nil {
		return operandA, operandB, operation, isRoman, errors.New(err.Error())
	}

	return operandA, operandB, operation, isRoman, nil
}

func getCorrectDataFromLine() (int, int, string, bool, error) {
	var operandARaw string
	var operandBRaw string
	var operation string
	var err error

	operandARaw, operandBRaw, operation, err = getRawLine()
	if err != nil {
		return -1, -1, operation, false, errors.New(err.Error())
	}

	operandA, operandB, isRoman, err := checkLine(operandARaw, operandBRaw, operation)
	if err != nil {
		return -1, -1, operation, false, errors.New(err.Error())
	}

	return operandA, operandB, operation, isRoman, nil
}

func checkLine(operandARaw string, operandBRaw string, operation string) (int, int, bool, error) {
	var operandA int
	var operandB int
	var isARoman bool
	var isBRoman bool
	var isRoman bool = false
	var err error

	err = checkOperation(operation)
	if err != nil {
		return -1, -1, false, errors.New(err.Error())
	}
	operandA, isARoman, err = getOperand(operandARaw)
	if err != nil {
		return -1, -1, false, errors.New(err.Error())
	}
	operandB, isBRoman, err = getOperand(operandBRaw)
	if err != nil {
		return -1, -1, false, errors.New(err.Error())
	}

	if isARoman != isBRoman {
		return -1, -1, false, errors.New("mixed numerals")
	}
	if isARoman && isBRoman {
		isRoman = true
	}
	return operandA, operandB, isRoman, nil
}

func getRawLine() (string, string, string, error) {
	var operandA string
	var operandB string
	var operation string

	var count int
	var err error

	count, err = fmt.Scanln(&operandA, &operation, &operandB)
	if count != 3 || err != nil {
		return operandA, operandB, operation, errors.New(err.Error())
	}

	return operandA, operandB, operation, nil
}

func checkOperation(operation string) error {
	if operation == "+" {
		return nil
	}
	if operation == "-" {
		return nil
	}
	if operation == "*" {
		return nil
	}
	if operation == "/" {
		return nil
	}

	return errors.New("unknown operation")
}

func getOperand(rawOperand string) (int, bool, error) {

	intOperand, err := strconv.Atoi(rawOperand)
	if err == nil {
		return intOperand, false, checkOperand(intOperand)
	}

	if err != nil {
		converted, err := converter.ConvertToArabic(rawOperand)
		if err != nil {
			return -1, true, errors.New(err.Error())
		}

		return converted, true, checkOperand(converted)
	}

	return -1, false, errors.New(err.Error())
}

func checkOperand(operand int) error {

	if operand > 0 && operand < 11 {
		return nil
	}

	return errors.New("operand out of bounds")
}
