package randomGenerator

import "fmt"

func MetodoCongruenteLinear(semente, a, c, m, tamanho int) []int {
	result := make([]int, tamanho)
	result[0] = semente
	fmt.Printf("Assinalou a semente %q", result[0])
	for i := 1; i < len(result); i++ {
		result[i] = ((result[i-1] * a) + c) % m
	}

	return result
}

func ExemploAula(tamanho int) []int {
	return MetodoCongruenteLinear(7, 4, 4, 9, tamanho)
}
