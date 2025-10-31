package main // **TODO** arquivo Go é iniciado declarando ao pacote que ele pertence. Quando se é dito que ele pertence ao pacote main, o programa indica ao compilador que o código em questão deve se tornar um **programa executável**. Se o package possuísse qualquer outro nome, seria apenas uma biblioteca para ser utilizada por outros programas.

import "fmt" // Importação da biblioteca padrão do Go. Fmt é a abreviação para **format**, e esta lib contém funções de formatação e impressão de texto.

func main() { // func -> forma como se são declaradas as funções. A função main é a função principal que o compilador buscará para executar todo o código contido nela.
    fmt.Println("Olá, Mundo!") //Println é a função exportada da biblioteca fmt previamente importada.
}