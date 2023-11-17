package converter

import (
	"testing"
)

func TestCanConvertToRoman(t *testing.T) {
	var roman string
	var err error

	roman, err = ConvertToRoman(1)
	if roman != "I" || err != nil {
		t.Errorf("incorrect")
	}

	roman, err = ConvertToRoman(4)
	if roman != "IV" || err != nil {
		t.Errorf("incorrect")
	}

	roman, err = ConvertToRoman(9)
	if roman != "IX" || err != nil {
		t.Errorf("incorrect")
	}

	roman, err = ConvertToRoman(10)
	if roman != "X" || err != nil {
		t.Errorf("incorrect")
	}

	roman, err = ConvertToRoman(44)
	if roman != "XLIV" || err != nil {
		t.Errorf("incorrect")
	}

	roman, err = ConvertToRoman(49)
	if roman != "XLIX" || err != nil {
		t.Errorf("incorrect")
	}

	roman, err = ConvertToRoman(81)
	if roman != "LXXXI" || err != nil {
		t.Errorf("incorrect")
	}

	roman, err = ConvertToRoman(100)
	if roman != "C" || err != nil {
		t.Errorf("incorrect")
	}
}

func TestCanConvertToRomanAndNumbersHigherThan100AreNotSupported(t *testing.T) {
	var roman string
	var err error

	roman, err = ConvertToRoman(200)
	if roman != "" || err != nil {
		t.Errorf("incorrect")
	}
}

func TestErrorWhenConvertingNonPositiveToRoman(t *testing.T) {
	var roman string
	var err error

	roman, err = ConvertToRoman(0)
	if roman != "" || err == nil {
		t.Errorf("incorrect")
	}

	roman, err = ConvertToRoman(-1)
	if roman != "" || err == nil {
		t.Errorf("incorrect")
	}
}

func TestCanConvertToArabic(t *testing.T) {
	var arabic int
	var err error

	arabic, err = ConvertToArabic("I")
	if arabic != 1 || err != nil {
		t.Errorf("incorrect")
	}

	arabic, err = ConvertToArabic("IV")
	if arabic != 4 || err != nil {
		t.Errorf("incorrect")
	}

	arabic, err = ConvertToArabic("IX")
	if arabic != 9 || err != nil {
		t.Errorf("incorrect")
	}

	arabic, err = ConvertToArabic("X")
	if arabic != 10 || err != nil {
		t.Errorf("incorrect")
	}
}

func TestCanConvertToArabicAndDontSupportNumbersLargerThan10(t *testing.T) {
	var arabic int
	var err error

	arabic, err = ConvertToArabic("XLIV")
	if arabic == 44 || err == nil {
		t.Errorf("incorrect")
	}

	arabic, err = ConvertToArabic("XLIX")
	if arabic == 49 || err == nil {
		t.Errorf("incorrect")
	}

	arabic, err = ConvertToArabic("LXXXI")
	if arabic == 81 || err == nil {
		t.Errorf("incorrect")
	}

	arabic, err = ConvertToArabic("C")
	if arabic == 100 || err == nil {
		t.Errorf("incorrect")
	}
}
