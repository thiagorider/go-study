package main

import (
    "fmt"
    "strings"
)

func main() {
    texto := "Um algoritmo é uma sequência de instruções bem definidas, normalmente usadas para resolver problemas de matemática específicos, executar tarefas, ou para realizar cálculos e equações. A origem da palavra “algoritmo” remete a Al Khowarizmi, famoso matemático árabe do século IX"

	// Declara uma lista de sinais de pontuação
    sinaisPontuacao := []string{".", ",", "!", "?", ":", ";", "-", "_", "(", ")", "[", "]", "{", "}", "<", ">", `"`}

    // Cria um mapa para armazenar a contagem de palavras
    //contagemPalavras := make(map[string]int)

    // Divide o texto em palavras
    palavras := strings.Fields(texto)
	

	fmt.Println(len(palavras))
	fmt.Println("---------------------")

	for _, palavra := range palavras {
        fmt.Println("Slice:", palavra)
        for _, sinal := range sinaisPontuacao {
            palavra[i] = strings.TrimRight(palavra, sinal)
        }
        fmt.Println("Limpando slice:", palavra)
    }
}