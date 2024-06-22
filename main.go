package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArabicMap = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRomanMap = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите выражение (например, 1 + 2 или VI / III):")
	scanner.Scan()
	input := scanner.Text()

	result, err := calculate(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}

func calculate(input string) (string, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return "", fmt.Errorf("выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}

	aStr, operator, bStr := parts[0], parts[1], parts[2]

	a, isRomanA := romanToArabicMap[aStr]
	b, isRomanB := romanToArabicMap[bStr]

	if isRomanA != isRomanB {
		return "", fmt.Errorf("выдача паники, так как используются одновременно разные системы счисления")
	}

	var aInt, bInt int
	var resultInt int

	if isRomanA {
		aInt = a
		bInt = b
	} else {
		var err error
		aInt, err = strconv.Atoi(aStr)
		if err != nil || aInt < 1 || aInt > 10 {
			return "", fmt.Errorf("выдача паники, прекращение работы")
		}
		bInt, err = strconv.Atoi(bStr)
		if err != nil || bInt < 1 || bInt > 10 {
			return "", fmt.Errorf("выдача паники, прекращение работы")
		}
	}

	switch operator {
	case "+":
		resultInt = aInt + bInt
	case "-":
		resultInt = aInt - bInt
	case "*":
		resultInt = aInt * bInt
	case "/":
		if bInt == 0 {
			return "", fmt.Errorf("выдача паники, деление на ноль")
		}
		resultInt = aInt / bInt
	default:
		return "", fmt.Errorf("выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}

	if isRomanA {
		if resultInt < 1 {
			return "", fmt.Errorf("выдача паники, так как в римской системе нет отрицательных чисел")
		}
		return arabicToRoman(resultInt), nil
	}

	return strconv.Itoa(resultInt), nil
}

func arabicToRoman(num int) string {
	var result strings.Builder

	for num >= 10 {
		result.WriteString("X")
		num -= 10
	}
	if num > 0 {
		result.WriteString(arabicToRomanMap[num])
	}

	return result.String()
}
