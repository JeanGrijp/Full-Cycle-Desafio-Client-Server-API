# Usa uma imagem oficial leve do Go
FROM golang:1.23-alpine

# Define diretório de trabalho dentro do container
WORKDIR /app

# Define variáveis para melhorar compatibilidade e evitar erros com SQLite/Postgres
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Copia apenas os arquivos de dependências primeiro para cache eficiente
COPY go.mod go.sum ./
RUN go mod download

# Agora copia todo o código do projeto
COPY . .

# Compila o binário final
RUN go build -o main .

# Expõe a porta que a aplicação usa
EXPOSE 8080

# Comando para rodar o servidor
CMD ["./main"]
