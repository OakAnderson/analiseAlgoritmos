# Analise de Algoritmos

## Algoritmos de ordenação:

```go
package main

import (
    "fmt"

    "github.com/OakAnderson/analiseAlgoritmos/cronometer"
)

// Algoritmo do bubble sort para exemplo
func bubbleSort(arr []int) {
    swap := func(i, j int) {
        arr[i], arr[j] = arr[j], arr[i]
    }

    n := len(arr)
    swapped := true

    for x := 0; swapped; x++ {
        swapped = false
        for i := 1; i < n-x; i++ {
            if arr[i-1] > arr[i] {
                swap(i-1, i)
                swapped = true
            }
        }
    }
}

func main() {
    var crono cronometer.Sort

    n := 10000
    crono.SetArrSize(n) // Uma lista com valores aleatórios será criada com tamanho n
    crono.SetFunction(bubbleSort) // Atribui a função bubbleSort para efetuar o próximo teste

    singleTest := crono.SingleTest() // Executa um único teste e retorna o tempo que o algoritmo levou para concluir a execução
    mean, _ := crono.MultipleTestsMean(100) // Executa 100 testes para arrays diferentes de mesmo tamanho e retorna a média dos resultados

    fmt.Printf("Teste único: %v\tMédia para 100 testes: %v\n", singleTest, mean)
}

```
