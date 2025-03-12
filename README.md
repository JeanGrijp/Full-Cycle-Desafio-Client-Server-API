# Desafio - Módulo Banco de Dados | Pós-graduação Go Expert

Este repositório contém a resolução do desafio do módulo de Banco de Dados da pós-graduação Go Expert, oferecida pela Full Cycle Faculdade de Tecnologia.

## Descrição do Desafio

Olá dev, tudo bem?

Neste desafio vamos aplicar o que aprendemos sobre webserver HTTP, contextos, banco de dados e manipulação de arquivos com Go.

Você precisará entregar dois sistemas desenvolvidos em Go:
- `client.go`
- `server.go`

### Requisitos

- O `client.go` deverá realizar uma requisição HTTP ao `server.go`, solicitando a cotação do dólar.

- O `server.go` deverá consumir a API com o câmbio de Dólar e Real disponível no endereço:
  ```
  https://economia.awesomeapi.com.br/json/last/USD-BRL
  ```
  Em seguida, deverá retornar o resultado no formato JSON para o cliente.

- Usando o pacote `context`, o `server.go` deverá registrar no banco de dados SQLite cada cotação recebida, respeitando os seguintes timeouts:
  - Timeout máximo de **200ms** para a requisição na API externa.
  - Timeout máximo de **10ms** para persistir os dados no banco de dados.

- O `client.go` deverá receber do `server.go` apenas o valor atual do câmbio (campo `bid` do JSON). Utilizando o pacote `context`, o `client.go` terá um timeout máximo de **300ms** para receber a resposta do servidor.

- Todos os contextos deverão retornar erro nos logs caso o tempo de execução seja insuficiente.

- O `client.go` deverá salvar a cotação atual em um arquivo chamado `cotacao.txt`, no seguinte formato:
  ```
  Dólar: {valor}
  ```

- O endpoint fornecido pelo `server.go` deverá ser `/cotacao`, executando na porta **8080**.

## Execução do Projeto

### Servidor
```bash
go run server.go
```

### Cliente
```bash
go run client.go
```

## Tecnologias Utilizadas

- Golang
- SQLite
- HTTP e Context package do Go

---

Ao finalizar, envie o link deste repositório para a correção do desafio. Boa sorte!

