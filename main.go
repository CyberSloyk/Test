package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var romanNumerals = []RomanNumeral{
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

func toRoman(n int) string {
	result := ""
	for _, numeral := range romanNumerals {
		for n >= numeral.Value {
			result += numeral.Symbol
			n -= numeral.Value
		}
	}
	return result
}

func toArabic(s string) (int, error) {
	romanMap := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	var res int
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romanMap[string(s[i])] < romanMap[string(s[i+1])] {
			res += romanMap[string(s[i+1])] - romanMap[string(s[i])]
			i++
		} else {
			res += romanMap[string(s[i])]
		}
	}
	return res, nil
}

func calculate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("некорректная операция")
	}
}

func main() {
	var input string
	fmt.Println("Введите выражение (число операция число):")
	fmt.Scanln(&input)

	input = strings.ReplaceAll(input, " ", "")
	if len(input) < 3 {
		fmt.Println("Ошибка: формат математической операции не удовлетворяет заданию")
		os.Exit(1)
	}

	op := ""
	for _, char := range []string{"+", "-", "*", "/"} {
		if strings.Contains(input, char) {
			op = char
			break
		}
	}

	if op == "" {
		fmt.Println("Ошибка: некорректная операция")
		os.Exit(1)
	}

	parts := strings.Split(input, op)
	if len(parts) != 2 {
		fmt.Println("Ошибка: формат математической операции не удовлетворяет заданию — два операнда и один оператор")
		os.Exit(1)
	}

	aStr, bStr := parts[0], parts[1]

	a, aErr := strconv.Atoi(aStr)
	b, bErr := strconv.Atoi(bStr)

	if aErr != nil && bErr != nil {
		a, err := toArabic(aStr)
		if err != nil || a < 1 || a > 10 {
			fmt.Println("Ошибка: неверный формат римского числа или число вне диапазона (I - X)")
			os.Exit(1)
		}
		b, err := toArabic(bStr)
		if err != nil || b < 1 || b > 10 {
			fmt.Println("Ошибка: неверный формат римского числа или число вне диапазона (I - X)")
			os.Exit(1)
		}
		result, err := calculate(a, b, op)
		if err != nil {
			fmt.Println("Ошибка:", err)
			os.Exit(1)
		}
		if result <= 0 {
			fmt.Println("Ошибка: результат для римских чисел не может быть меньше или равен нулю")
			os.Exit(1)
		}
		fmt.Println("Результат:", toRoman(result))
	} else if aErr == nil && bErr == nil {
		if a < 1 || a > 10 || b < 1 || b > 10 {
			fmt.Println("Ошибка: числа должны быть от 1 до 10")
			os.Exit(1)
		}
		result, err := calculate(a, b, op)
		if err != nil {
			fmt.Println("Ошибка:", err)
			os.Exit(1)
		}
		fmt.Println("Результат:", result)
	} else {
		fmt.Println("Ошибка: нельзя смешивать арабские и римские числа")
		os.Exit(1)
	}
}
