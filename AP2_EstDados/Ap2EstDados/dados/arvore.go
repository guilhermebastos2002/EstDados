package dados

import "fmt"

type NoABB struct {
	Produto  Produto
	Esquerda *NoABB
	Direita  *NoABB
}

type ArvoreABB struct {
	Raiz *NoABB
}

func (a *ArvoreABB) Adicionar(nome, descricao string, valor float64) {
	produto := Produto{
		Nome:      nome,
		Descricao: descricao,
		Valor:     valor,
	}

	a.Raiz = adicionarNo(a.Raiz, produto)
}

func adicionarNo(no *NoABB, produto Produto) *NoABB {
	if no == nil {
		return &NoABB{Produto: produto}
	}

	if produto.Nome < no.Produto.Nome {
		no.Esquerda = adicionarNo(no.Esquerda, produto)
	} else {
		no.Direita = adicionarNo(no.Direita, produto)
	}

	return no
}

func (a *ArvoreABB) Buscar(nome string) (*Produto, error) {
	no := buscarNo(a.Raiz, nome)
	if no == nil {
		return nil, fmt.Errorf("produto não encontrado")
	}
	return &no.Produto, nil
}

func buscarNo(no *NoABB, nome string) *NoABB {
	if no == nil || no.Produto.Nome == nome {
		return no
	}

	if nome < no.Produto.Nome {
		return buscarNo(no.Esquerda, nome)
	}
	return buscarNo(no.Direita, nome)
}

func (a *ArvoreABB) Remover(nome string) error {
	a.Raiz = removerNo(a.Raiz, nome)
	return nil
}

func removerNo(no *NoABB, nome string) *NoABB {
	if no == nil {
		return nil
	}

	if nome < no.Produto.Nome {
		no.Esquerda = removerNo(no.Esquerda, nome)
	} else if nome > no.Produto.Nome {
		no.Direita = removerNo(no.Direita, nome)
	} else {
		if no.Esquerda == nil {
			return no.Direita
		} else if no.Direita == nil {
			return no.Esquerda
		}

		sucessor := encontrarMenorNo(no.Direita)
		no.Produto = sucessor.Produto

		no.Direita = removerNo(no.Direita, sucessor.Produto.Nome)
	}

	return no
}

func encontrarMenorNo(no *NoABB) *NoABB {
	atual := no
	for atual.Esquerda != nil {
		atual = atual.Esquerda
	}
	return atual
}

func (a *ArvoreABB) Listar() []Produto {
	var lista []Produto
	listarNos(a.Raiz, &lista)
	return lista
}

func listarNos(no *NoABB, lista *[]Produto) {
	if no != nil {
		listarNos(no.Esquerda, lista)
		*lista = append(*lista, no.Produto)
		listarNos(no.Direita, lista)
	}
}