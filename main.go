package main

import "fmt"

func main() {
	input := 100

	generateFibonacci(input)
}

func romanToInt(input string) int {
	numericMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	result := 0
	i := 0

	for i < len(input) {
		current := numericMap[string(input[i])]

		if i < len(input)-1 {
			next := numericMap[string(input[i+1])]

			if next > current {
				result += next - current
				i += 2
				continue
			}
		}

		result += current
		i++
	}

	return result
}

func generateFibonacci(input int) {
	memo := map[int]int{}
	for i := 0; i <= input; i++ {
		fib := fibonacci(i, memo)

		fmt.Printf("%v ", fib)
	}
}

func fibonacci(input int, memo map[int]int) int {
	if stored, ok := memo[input]; ok {
		return stored
	}

	if input == 0 {
		memo[input] = 0
		return 0
	}

	if input == 1 {
		memo[input] = 1
		return 1
	}

	res := fibonacci(input-1, memo) + fibonacci(input-2, memo)
	memo[input] = res

	return res
}
