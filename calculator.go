package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	calculator()
}

func customPanic(text string) {
	if text == "" {
		text = "не верные операнды"
	}

	panic(fmt.Sprintf("Ошибка: %s\n", text))
}

func getOperator(userInput string) string {
	var (
		operators       = "+-*/"
		currentOperator string
	)

	for _, char := range userInput {
		if strings.Contains(operators, string(char)) {
			currentOperator = string(char)
			break
		}
	}

	if currentOperator == "" {
		customPanic("")
	}

	return currentOperator
}

func arabicToRoman(num int) string {
	values := [9]int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := [9]string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var result string

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			result += numerals[i]
		}
	}

	return result
}

func calculator() {
	reader := bufio.NewReader(os.Stdin)
	validInputRomanNumerals := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}

	fmt.Println("Введите выражение...")
	userInput, _ := reader.ReadString('\n')

	currentOperator := getOperator(userInput)
	operands := strings.Split(userInput, currentOperator)

	if len(operands) > 2 {
		customPanic("")
	}

	var (
		romanNumeralsCount int
		arabicNumbersCount int
	)

	for i, operand := range operands {
		operands[i] = strings.TrimSpace(operand)

		if num, err := strconv.Atoi(operands[i]); err == nil && num <= 10 {
			arabicNumbersCount++
		}

		if _, ok := validInputRomanNumerals[strings.ToUpper(operands[i])]; ok {
			operands[i] = strings.ToUpper(operands[i])
			romanNumeralsCount++
		}
	}

	isArabicNumbers := arabicNumbersCount == 2
	isRomanNumerals := romanNumeralsCount == 2

	fmt.Printf("romanNumeralsCount %d arabicNumbersCount %d\n", romanNumeralsCount, arabicNumbersCount)

	if !isArabicNumbers && !isRomanNumerals {
		customPanic("")
	}

	fmt.Printf("Type: %T Value: %#v\n", operands, operands)

	var (
		firstOperand  int
		secondOperand int
	)

	if isArabicNumbers {
		firstOperand, _ = strconv.Atoi(operands[0])
		secondOperand, _ = strconv.Atoi(operands[1])
	}

	if isRomanNumerals {
		firstOperand = validInputRomanNumerals[operands[0]]
		secondOperand = validInputRomanNumerals[operands[1]]
	}

	var output int

	switch currentOperator {
	case "+":
		output = firstOperand + secondOperand
	case "-":
		output = firstOperand - secondOperand

		if isRomanNumerals && output < 1 {
			customPanic("римскими могут быть только положительные числа")
		}
	case "*":
		output = firstOperand * secondOperand
	case "/":
		if secondOperand == 0 {
			customPanic("на 0 не делим")
		}

		output = firstOperand / secondOperand
	}

	if isRomanNumerals {
		fmt.Printf("Output: %s", arabicToRoman(output))
	} else {
		fmt.Printf("Output: %d", output)
	}
}
