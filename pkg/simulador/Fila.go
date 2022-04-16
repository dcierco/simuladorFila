package simulador

type Fila struct {
	tempoChegada,
	tempoAtendimento,
	numeroServidores,
	capacidade,
	qtdeNaFila int
	fila []float32
}

func NewFila(tempoChegada, tempoAtendimento, numeroServidores, capacidade int) *Fila {
	p := new(Fila)
	p.tempoChegada = tempoChegada
	p.tempoAtendimento = tempoAtendimento
	p.numeroServidores = numeroServidores
	p.capacidade = capacidade
	p.fila = make([]float32, 0, capacidade)
	p.qtdeNaFila = 0
	return p
}

func (fila Fila) estaVazia() bool {
	if fila.qtdeNaFila == 0 {
		return true
	}
	return false
}

func (fila Fila) estaCheia() bool {
	if fila.qtdeNaFila == fila.capacidade {
		return true
	}
	return false
}

func (receiver Fila) adicionaFila() {

}

func (receiver Fila) removeFila() {

}
