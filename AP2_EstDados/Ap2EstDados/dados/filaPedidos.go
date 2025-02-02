package dados

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const VALOR_DELIVERY = 10.0

type Pedidos struct {
	Fila []Pedido
}

func (p *Pedidos) inicializar() {
	p.Fila = make([]Pedido, 0)
}

func (p *Pedidos) Adicionar(delivery bool, idProdutos []int) error {
	valorTotal := 0.0
	if delivery {
		valorTotal += VALOR_DELIVERY
	}

	for _, id := range idProdutos {
		prod, err := ListaProdutos.Buscar(id)
		if err != nil {
			return err
		}

		valorTotal += prod.Valor
	}

	pedido := Pedido{Delivery: delivery, IdProdutos: idProdutos, ValorTotal: valorTotal}
	pedido.RegistrarID()

	p.Fila = append(p.Fila, pedido)
	MetricasColetadas.PedidosEmAndamento++

	currentTime := time.Now()
	fmt.Printf("%s Pedido adicionado à fila\n", currentTime.Format("02/01/2006 15:04:05"))

	return nil
}

func (p *Pedidos) Expedir() error {
	if len(p.Fila) == 0 {
		return fmt.Errorf("fila de pedidos está vazia")
	}

	MetricasColetadas.PedidosEmAndamento--
	MetricasColetadas.PedidosEncerrados++
	MetricasColetadas.FaturamentoTotal += p.Fila[0].ValorTotal

	currentTime := time.Now()
	fmt.Printf("%s Expedido pedido %d, valor R$ %.2f\n", currentTime.Format("02/01/2006 15:04:05"), p.Fila[0].Id, p.Fila[0].ValorTotal)

	p.Fila = p.Fila[1:]

	return nil
}

func (p *Pedidos) Listar() []Pedido {
	return p.Fila
}
