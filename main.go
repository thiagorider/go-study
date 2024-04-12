package main

import "fmt"

// Pessoa é uma struct que representa uma pessoa.
// Ela tem dois campos: Nome e Idade.
type Pessoa struct {
    Nome  string
    Idade int
}

// DentroDaFaixaEtaria retorna true se a idade da pessoa estiver entre min e max, inclusive.
func (p Pessoa) DentroDaFaixaEtaria(min, max int) bool {
    return p.Idade >= min && p.Idade <= max
}

// main é a função principal do programa.
// Ela cria uma instância da struct Pessoa e a imprime no console.
func main() {
	P := Pessoa{"João", 30}
	fmt.Println(P.DentroDaFaixaEtaria(20, 40)) // true
	fmt.Println(P.DentroDaFaixaEtaria(41, 50)) // false
}