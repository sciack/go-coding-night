package main

import "fmt"

func decodeChar(value byte) uint {
	switch value {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'M':
		return 1000
	case 'D':
		return 500
	default:
		return 0
	}
}

func arabicNumber(roman string) uint {
	var result uint = 0
	var inputLen = len(roman)
	for i := 0; i < inputLen; i++ {
		var currentChar = decodeChar(roman[i])
		if i < inputLen-1 && decodeChar(roman[i+1]) > currentChar {
			currentChar = -currentChar
		}
		result += currentChar

	}
	return result
}

func main() {
	fmt.Printf("Arabic number of MMXXI is %d", arabicNumber("MMXXI"))
}
