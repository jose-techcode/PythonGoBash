package main

import "fmt"

func main() {

	// Normal

	// Nome
	
	var nome string
	fmt.Print("Digite seu nome:")
	fmt.Scanln(&nome) // Equivalente ao input do python
	fmt.Println("Olá,", nome)

	// Idade

    var idade int
	fmt.Print("Digite sua idade:")
	fmt.Scanln(&idade)
	fmt.Println("Sua idade é:", idade)

	// Tratamento de erros

    // Nome

	var nome1 string
	fmt.Println("Digite seu nome:")

	if _, err := fmt.Scanln(&nome1); err != nil {
		fmt.Println("Erro na entrada:", err)
		return
	}
    
	fmt.Println("Seu nome é:", nome1)

	// Idade

	var idade0 int
	fmt.Println("Digite sua idade:")

	if _, err := fmt.Scanln(&idade0); err != nil {
		fmt.Println("Erro na entrada:", err)
		return
	}

    fmt.Println("Sua idade é:", idade0)

	// While true (for infinito) com tratamento de erros incluído

	var texto string
    fmt.Println("Digite um texto (ou 'sair' para sair):")

    for {
        fmt.Print("Saída:")
        _, err := fmt.Scanln(&texto)

        if err != nil {
            fmt.Println("Erro na digitação do texto:", err)
            continue // continue ou break
        }

        if texto == "sair" {
            fmt.Println("Encerrado.")
            break
        }

        fmt.Println("Texto:", texto)
    }
}