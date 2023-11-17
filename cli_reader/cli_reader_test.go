package cli_reader

import (
	"testing"
)

func TestRejectLongLines(t *testing.T) {
	// cannot test it  - need to mock fmt

}

func TestCanReadArabicLines(t *testing.T) {
	var err error

	operandA, operandB, isRoman, err := checkLine("1", "2", "+")
	if err != nil {
		t.Errorf(err.Error())
	}
	if operandA != 1 {
		t.Errorf("expected 1, got %v", operandA)
	}
	if operandB != 2 {
		t.Errorf("expected 2, got %v", operandB)
	}
	if isRoman == true {
		t.Errorf("expected false, got %v", isRoman)
	}

}

func TestCanReadRomanLines(t *testing.T) {
	var err error

	operandA, operandB, isRoman, err := checkLine("III", "IV", "+")
	if err != nil {
		t.Errorf(err.Error())
	}
	if operandA != 3 {
		t.Errorf("expected 3, got %v", operandA)
	}
	if operandB != 4 {
		t.Errorf("expected 4, got %v", operandB)
	}
	if isRoman == false {
		t.Errorf("expected true, got %v", isRoman)
	}

}

func TestRejectMixedNumerals(t *testing.T) {
	var err error
	_, _, _, err = checkLine("1", "V", "+")
	if err == nil {
		t.Errorf("expected mixed numerals error, got nil")
	}

	_, _, _, err = checkLine("1", "I", "+")
	if err == nil {
		t.Errorf("expected mixed numerals error, got nil")
	}

	_, _, _, err = checkLine("IV", "5", "+")
	if err == nil {
		t.Errorf("expected mixed numerals error, got nil")
	}
}

func TestRejectFloats(t *testing.T) {
	_, _, _, err := checkLine("1.5", "2", "+")
	if err == nil {
		t.Errorf("expected float error, got nil")
	}
}

func TestRejectNegative(t *testing.T) {

	_, _, _, err := checkLine("-1", "2", "+")
	if err == nil {
		t.Errorf("expected negative error, got nil")
	}
}
