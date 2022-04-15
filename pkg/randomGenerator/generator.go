package randomgenerator

func metodoCongruenteLinear(semente, a, c, m, tamanho int) []int {
	result := make([]int, 0, tamanho)
	result[0] = semente

	for i := range result[:1] {
		result[i] = ((result[i-1] * a) + c) % m
	}

	return result
}

func exemploAula(tamanho int) []int {
	return metodoCongruenteLinear(7, 4, 4, 9, tamanho)
}
