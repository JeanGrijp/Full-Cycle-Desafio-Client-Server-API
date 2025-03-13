# Guia de Execução do Server

Este documento explica como configurar e executar o projeto utilizando **Docker** e **Docker Compose** para rodar a API em Go integrada ao banco de dados **PostgreSQL**.

## ✅ Pré-requisitos

Antes de iniciar, certifique-se de ter instalado:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Git](https://git-scm.com/downloads) (opcional, mas recomendado)

---

## 🚀 Passo 1: Clonar o Repositório

Se você ainda não clonou o projeto, faça isso com o seguinte comando:

```bash
git clone https://github.com/JeanGrijp/Full-Cycle-Desafio-Client-Server-API.git
cd Full-Cycle-Desafio-Client-Server-API/server
```

---

## 📦 Passo 2: Criar e Configurar as Imagens Docker

O projeto já possui um `Dockerfile` e um `docker-compose.yml` configurados. Para construir as imagens e iniciar os containers, execute:

```bash
docker-compose up --build
```

Isso fará com que:

- O PostgreSQL seja iniciado como um serviço no Docker.
- A API em Go seja compilada e iniciada.
- O banco de dados seja criado automaticamente com a estrutura necessária.

Caso queira rodar os containers em **modo detached** (em segundo plano), utilize:

```bash
docker-compose up --build -d
```

---

## 🛠️ Passo 3: Verificar os Containers em Execução

Para checar se os containers estão rodando corretamente, use:

```bash
docker ps
```

Você deve ver algo assim:

```plaintext
CONTAINER ID   IMAGE           STATUS       PORTS                    NAMES
123456789abc   go-api-server   Up 2 minutes 0.0.0.0:8080->8080/tcp   go-api-server
987654321def   postgres:16     Up 2 minutes 0.0.0.0:5432->5432/tcp   postgres-db
```

Se apenas o banco estiver rodando e a API não subir, verifique os logs com:

```bash
docker-compose logs api
```

---

## 🌍 Passo 4: Testar a API

Com a aplicação rodando, teste o endpoint da cotação do dólar:

```bash
curl http://localhost:8080/cotacao
```

Ou abra no navegador:

👉 [http://localhost:8080/cotacao](http://localhost:8080/cotacao)

Se tudo estiver certo, você verá um JSON semelhante a este:

```json
{
  "dolar": 5.8099
}
```

---

## ⛔ Passo 5: Parar os Containers

Para parar os containers sem removê-los:

```bash
docker-compose stop
```

Se quiser **parar e remover** os containers, execute:

```bash
docker-compose down
```

Caso precise limpar volumes do banco de dados e reconstruir tudo do zero:

```bash
docker-compose down -v
```

---

## 🔥 Conclusão

Agora você tem o projeto rodando completamente em containers Docker. Sempre que precisar iniciar o ambiente novamente, basta rodar:

```bash
docker-compose up -d
```

Se precisar de ajustes ou encontrar problemas, verifique os logs da API:

```bash
docker-compose logs api
```

Ou do banco de dados:

```bash
docker-compose logs db
```

Beba água! 💧
