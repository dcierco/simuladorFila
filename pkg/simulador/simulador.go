package simulador

type simulador struct {
	tempoAtual float64
}

func chegada(fila Fila) {
	contabilizaTempo()
	if fila.estaVazia() {
		fila.adicionaFila()
		if fila.qtdeNaFila <= fila.numeroServidores {
			agendaSaida()
		}
	} else {
		//perda++
	}
	agendaChegada()
}

func saida(fila Fila) {
	contabilizaTempo()
	fila.removeFila()

}

func contabilizaTempo() {
	panic("unimplemented")
}

func agendaChegada() {
	panic("unimplemented")
}

func agendaSaida() {
	panic("unimplemented")
}
