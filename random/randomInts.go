package random

import "math/rand"


// Ints retorna um slice de n valores aleatórios
func Ints(n int) []int {
    arr := make([]int, n)

    for i := range arr {
        arr[i] = rand.Int()
    }
    
    return arr
}

