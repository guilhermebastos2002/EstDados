# API de Gerenciamento de Loja

Esta é uma API de gerenciamento de loja desenvolvida em Go, que oferece funcionalidades para adicionar, remover e listar produtos, gerenciar pedidos e coletar métricas sobre a loja.

## Rotas Disponíveis

### Produtos

- `POST /produto`: Adiciona um novo produto ao estoque.
- `GET /produto?id={id}`: Busca um produto pelo seu ID.
- `DELETE /produto?id={id}`: Remove um produto do estoque.
- `GET /produtos`: Lista todos os produtos disponíveis no estoque.

### Pedidos

- `POST /pedido`: Adiciona um novo pedido à fila.
- `GET /pedidos?ordenar={algoritmo}`: Lista todos os pedidos na fila. O parâmetro opcional "ordenar" permite ordenar os pedidos pelo valor total em ordem crescente, utilizando os algoritmos "bubblesort" ou "quicksort".

### Métricas

- `GET /metricas`: Retorna as métricas coletadas sobre a loja.

### Loja

- `POST /abrir?intervalo={tempo}`: Abre a loja e inicia o processamento de pedidos com um intervalo especificado entre expedições (em segundos).
- `POST /fechar`: Fecha a loja.

## Uso

- Certifique-se de ter Go instalado em sua máquina.
- Clone este repositório: `git clone https://github.com/seu-usuario/api-loja.git`
- Navegue até o diretório do projeto: `cd api-loja`
- Execute o servidor: `go run main.go`
- A API estará acessível em `http://localhost:8080`

## Exemplo de Uso

```bash
# Adiciona um produto
curl -X POST -d '{"nome":"Produto A", "descricao":"Descrição do Produto A", "valor":10.5}' http://localhost:8080/produto

# Adiciona um pedido
curl -X POST -d '{"delivery":true, "id_produtos":[1,2]}' http://localhost:8080/pedido

# Lista os pedidos ordenados pelo valor total
curl http://localhost:8080/pedidos?ordenar=quicksort
