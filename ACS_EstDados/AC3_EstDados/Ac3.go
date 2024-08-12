package main

import (
	"fmt"
)

type Contato struct {
	nome  string
	email string
}

func (c *Contato) alterarEmail(novoEmail string) {
	c.email = novoEmail
}

func adicionarContato(contato Contato, arrayContatos []Contato) []Contato {
	for i := 0; i < len(arrayContatos); i++ {
		if arrayContatos[i].nome == "" && arrayContatos[i].email == "" {
			arrayContatos[i] = contato
			return arrayContatos
		}
	}
	fmt.Println("Não foi possivel adicionar o contato, o array está cheio.")
	return arrayContatos
}

func excluirContato(arrayContatos []Contato) []Contato {
	for i := len(arrayContatos) - 1; i >= 0; i-- {
		if arrayContatos[i].nome != "" || arrayContatos[i].email != "" {
			arrayContatos[i] = Contato{}
			return arrayContatos
		}
	}
	fmt.Println("Não há contatos para excluir.")
	return arrayContatos
}

func main() {
	arrayContatos := make([]Contato, 5)

	var opcao int
	for {
		fmt.Println("Selecione uma opção")
		fmt.Println("1 - Adicionar contato")
		fmt.Println("2 - Excluir ultimo contato")
		fmt.Println("3 - Sair")

		fmt.Println("Opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			var nome, email string
			fmt.Println("Digite o nome de contato: ")
			fmt.Scan(&nome)
			fmt.Println("Digite o email do contato: ")
			fmt.Scan(&email)
			novoContato := Contato{nome, email}
			arrayContatos = adicionarContato(novoContato, arrayContatos)
			fmt.Println("Contatos adicionado com sucesso.")
		case 2:
			arrayContatos = excluirContato(arrayContatos)
			fmt.Println("Último contato excluido.")
		case 3:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}
