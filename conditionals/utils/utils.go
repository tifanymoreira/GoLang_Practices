package utils

import "fmt"

func VerifyNumber(number int) {
	if number > 0 {
		fmt.Println(number, " é positivo.")
	} else if number < 0 {
		fmt.Println(number, "é negativo.")
	} else {
		fmt.Println("O número é zero.")
	}
}
