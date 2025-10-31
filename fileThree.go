package main

import "fmt"

func main() {
    numero := -5 // A partir da forma simplificada, infere-se que a variável numero é do tipo signed int.

	verifyNumber(numero)

}

func verifyNumber(number int) {
	if number > 0 {
		fmt.Println(number, " é positivo.")
	} else if number < 0 {
        fmt.Println(number, "é negativo.")
    } else {
        fmt.Println("O número é zero.")
    }
}