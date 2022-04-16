package randomGenerator

import (
	"math"
	"time"
)

func MetodoCongruenteLinear(semente, a, c, m, tamanho int) []float64 {
	result := make([]float64, tamanho)
	result[0] = float64(semente) / float64(m)
	lastInt := semente
	for i := 1; i < len(result); i++ {
		lastInt = ((lastInt * a) + c) % m
		result[i] = float64(lastInt) / float64(m)
	}

	return result
}

func ExemploAula(tamanho int) []int {
	// MetodoCongruenteLinear(7, 4, 4, 9, tamanho)
	result := make([]int, tamanho)
	result[0] = 7
	for i := 1; i < len(result); i++ {
		result[i] = ((result[i-1] * 4) + 4) % 9
	}

	return result
}

func numericalRecips(tamanho int) []float64 {
	return MetodoCongruenteLinear(int(time.Now().Unix()), 1664525, 1013904223, int(math.Exp2(32)), tamanho)
}

func glibc(tamanho int) []float64 {
	//Parecido, mas não pega os bits entre 30..0
	return MetodoCongruenteLinear(int(time.Now().Unix()), 1103515245, 12345, int(math.Exp2(32)), tamanho)
}

func microsoftVB(tamanho int) []float64 {
	return MetodoCongruenteLinear(int(time.Now().Unix()), 1140671485, 12820163, int(math.Exp2(24)), tamanho)
}

func javaUtilRandom(tamanho int) []float64 {
	return MetodoCongruenteLinear(int(time.Now().Unix()), 25214903917, 11, int(math.Exp2(48)), tamanho)
}

func javaUtilRandomSeed(tamanho, semente int) []float64 {
	return MetodoCongruenteLinear(semente, 25214903917, 11, int(math.Exp2(48)), tamanho)
}
