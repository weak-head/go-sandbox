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
