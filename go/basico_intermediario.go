package main

import (
	"fmt"
)

// Classe não tradicional

type Pessoa struct {
    Nome string
    Idade int
}

// Funções

func anonima() {
    multiplicacao := func(x float32, y float32) float32 {
        return x * y
    }
    fmt.Println(multiplicacao(10, 10))
}

func anonima2() {
    soma := func(a, z int) int {
        return a + z
    }
    fmt.Println(soma(5, 5))
}

func pessoa(nome string) string {
    return fmt.Sprintf("Seu nome: %s", nome)
}

func humano(idade int) string {
    return fmt.Sprintf("Sua idade: %d", idade )
}

func valor(numero float32) string {
    return fmt.Sprintf("Decimal: %f", numero)
}

func portugues(palavra string) string {
    return fmt.Sprintf("Palavra: %s", palavra)
}

func cadastro(nome0 string) string {
    return fmt.Sprintf("Nome: %s", nome0)
}

func calculadora_anonima(a, b float64) float64 {
    soma := func(a, b float64) float64 {
        return a + b
    }
    subtracao := func(a, b float64) float64 {
        return a - b
    }
    multiplicacao := func(a, b float64) float64 {
        return a * b
    }
    divisao := func(a, b float64) float64 {
        return a / b
    }
    fmt.Println(soma(a, b))
    fmt.Println(subtracao(a, b))
    fmt.Println(multiplicacao(a, b))
    fmt.Println(divisao(a, b))
    
    return soma(a, b)
    // return subtracao(c, d)
    // return multiplicacao(e, f)
    // return divisao(g, h)
}

// Função central

func main() {

    // numero, err := // valor com parametro
    // if err != nil {
        // fmt.Println("Oi")
    // }

    // Estrutura de controle

    resultado := calculadora_anonima(5, 5)
    fmt.Println("Resultado:", resultado)

    fmt.Println("Go funcionando!")
    
    fmt.Println(10 + 5)
    
    venda := 1000
    custo := 500
    lucro := venda - custo
    fmt.Println(lucro)

    nota := 7
    if nota >= 7 {
        fmt.Println("Passou")
    } else {
        fmt.Println("Não passou")
    }

    sim_nao := 0
    if sim_nao > 0 {
        fmt.Println(true)
    } else {
        fmt.Println(false)
    }

    for i := 3; i > 0; i-- {
        fmt.Println(i)
    }
    
    for i := 1; i < 3; i++ {
        fmt.Println(i)
    }

    // while

    contador := 1
    for contador < 3 {
        fmt.Printf("Contagem: %d\n", contador)
        contador += 1
    }

    // Lógica Booleana
    
    numero1 := true
    numero2 := false
        fmt.Println(!numero1) // not (true)
        fmt.Println(numero1 && numero2) // and (igual)
        fmt.Println(numero1 || numero2) // or (true)

    // Estrutura de dados

    list_slice := []int{1, 2, 3, 4}
    list_array := [4]int{1, 2, 3, 4}
    dict_texto := map[string]string{"Nome": "José"}
    dict_inteiro := map[string]int{"Idade": 25}

    fmt.Println(list_slice)
    fmt.Println(list_slice[0])
    fmt.Println(list_array)
    fmt.Println(list_array[1])
    fmt.Println(dict_texto)
    fmt.Println(dict_inteiro)

    anonima()
 
    anonima2()

    mensagem := pessoa("José")
    fmt.Println(mensagem)

    numero := humano(25)
    fmt.Println(numero)

    decimal := valor(24.75)
    fmt.Println(decimal)

    termo := portugues("Olá")
    fmt.Println(termo)

    pessoa := Pessoa{Nome: "José", Idade: 25}
    
    fmt.Println(pessoa.Nome)
    fmt.Println(pessoa.Idade)

    cadastro0 := cadastro("João")
    fmt.Println(cadastro0)

}