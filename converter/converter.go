package converter

import "errors"

func ConvertToRoman(arabicNumeral int) (string, error) {
	if arabicNumeral < 1 {
		return "", errors.New("cannot convert non-positive to roman")
	}
	units := arabicNumeral % 10
	tens := arabicNumeral / 10 % 10
	hundreds := arabicNumeral / 100 % 10

	hundredsString := getHundredsString(hundreds)
	tensString := getTensString(tens)
	unitsString := getUnitsString(units)

	return hundredsString + tensString + unitsString, nil
}

var arabicToUnitsMap = [...]string{
	"",
	"I",
	"II",
	"III",
	"IV",
	"V",
	"VI",
	"VII",
	"VIII",
	"IX",
}
var arabicToTensMap = [...]string{
	"",
	"X",
	"XX",
	"XXX",
	"XL",
	"L",
	"LX",
	"LXX",
	"LXXX",
	"XC",
}

// only need to support 1-10
func ConvertToArabic(romanNumeral string) (int, error) {

	if romanNumeral == "X" {
		return 10, nil
	}

	for k, v := range arabicToUnitsMap {
		if v == romanNumeral {
			return k, nil
		}
	}

	return -1, errors.New("could not convert " + romanNumeral)
}

func getHundredsString(arabicNumeral int) string {
	if arabicNumeral == 1 {
		return "C"
	}
	return ""
}

func getTensString(arabicNumeral int) string {
	return arabicToTensMap[arabicNumeral]
}

func getUnitsString(arabicNumeral int) string {
	return arabicToUnitsMap[arabicNumeral]
}
