package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите выражение (например, 1 + 2 или VI / III):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	aStr := parts[0]
	operator := parts[1]
	bStr := parts[2]

	romanToArabicMap := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	arabicToRomanMap := map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	}

	a, aIsRoman := romanToArabicMap[aStr]
	b, bIsRoman := romanToArabicMap[bStr]

	if aIsRoman != bIsRoman {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}

	if aIsRoman {
		result := doOperation(a, b, operator)
		if result < 1 {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
		fmt.Println("Результат:", arabicToRomanMap[result])
	} else {
		aInt, err := strconv.Atoi(aStr)
		if err != nil || aInt < 1 || aInt > 10 {
			panic("Выдача паники, введено неподходящее число.")
		}

		bInt, err := strconv.Atoi(bStr)
		if err != nil || bInt < 1 || bInt > 10 {
			panic("Выдача паники, введено неподходящее число.")
		}

		result := doOperation(aInt, bInt, operator)
		fmt.Println("Результат:", result)
	}
}

func doOperation(a int, b int, operator string) int {
	if operator == "+" {
		return a + b
	} else if operator == "-" {
		return a - b
	} else if operator == "*" {
		return a * b
	} else if operator == "/" {
		if b == 0 {
			panic("Выдача паники, деление на ноль.")
		}
		return a / b
	} else {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
}
