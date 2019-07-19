package proptests

import "strings"

var arabicToRomanMap = []struct {
	Arabic int
	Roman  string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

var romanToArabic = map[rune]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func ConvertToRoman(num int) string {
	var result strings.Builder
	for _, amap := range arabicToRomanMap {
		for num >= amap.Arabic {
			result.WriteString(amap.Roman)
			num -= amap.Arabic
		}
	}
	return result.String()
}

func ConvertToArabic(roman string) (arabic int) {
	var prev int
	for _, c := range roman {
		val := romanToArabic[c]
		if prev != 0 && prev < val {
			arabic -= prev
		} else {
			arabic += prev
		}
		prev = val
	}
	return arabic + prev
}
