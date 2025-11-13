package main // **TODO** arquivo Go é iniciado declarando ao pacote que ele pertence. Quando se é dito que ele pertence ao pacote main, o programa indica ao compilador que o código em questão deve se tornar um **programa executável**. Se o package possuísse qualquer outro nome, seria apenas uma biblioteca para ser utilizada por outros programas.

import "fmt" // Importação da biblioteca padrão do Go. Fmt é a abreviação para **format**, e esta lib contém funções de formatação e impressão de texto.

func main() { // func -> forma como se são declaradas as funções. A função main é a função principal que o compilador buscará para executar todo o código contido nela.

	// STRING ---------------
	// Usada para texto
	var nome string = "Mundo"
	otherName := "Tifany"

	// INT ---------------
	// Usado para números inteiros (sem casa decimal)
	var number int = 10
	otherNumber := 20
	var idade int32 = 45             // int de 32 bits (ocupa menos memória)
	var populacao int64 = 8000000000 // int de 64 bits (para números gigantes)

	// FLOAT (Números Decimais) ---------------
	var preco float64 = 19.99
	pi := 3.14159
	var taxa float32 = 0.5

	// BOOLEAN (Verdadeiro/Falso) ---------------
	var logado bool = true
	ativo := false

	// OUTROS TIPOS BÁSICOS ---------------

	// RUNE (para caracteres individuais)
	var letra rune = 'A'
	simbolo := 'ç'

	// BYTE
	var umByte byte = 255
	outroByte := byte(10) // Também podemos "converter" um número para byte

	// VALORES "ZERO" (Padrão) ---------------

	var s string  // Valor zero: "" (string vazia)
	var i int     // Valor zero: 0
	var f float64 // Valor zero: 0.0
	var b bool    // Valor zero: false

	fmt.Println("--- Strings ---")
	fmt.Println(nome, otherName, s)

	fmt.Println("--- Inteiros ---")
	fmt.Println(number, otherNumber, idade, populacao, i)

	fmt.Println("--- Floats ---")
	fmt.Println(preco, pi, taxa, f)

	fmt.Println("--- Booleans ---")
	fmt.Println(logado, ativo, b)

	fmt.Println("--- Outros ---")
	fmt.Printf("Rune: %c, %c\n", letra, simbolo)    // %c formata para caractere
	fmt.Printf("Byte: %d, %d\n", umByte, outroByte) // %d formata para número
}
