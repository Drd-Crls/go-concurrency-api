# Imagem base Go Alpine
FROM golang:1.24-alpine

WORKDIR /app

# Instalar ferramentas básicas
RUN apk add --no-cache git bash

# Copiar todo o código primeiro
COPY . .

# Baixar dependências e gerar go.sum se não existir
RUN go mod tidy

# Compilar a aplicação
RUN go build -o server cmd/server/main.go

EXPOSE 8080

CMD ["./server"]
