# Go Concurrency API

Este projeto é uma API em Go que combina REST e GraphQL para gerenciar usuários e posts. Ele utiliza `gqlgen` para GraphQL e `resty` para requisições HTTP. O projeto é estruturado para testes unitários, testes de integração e é containerizável via Docker.

## Estrutura do Projeto

```
concurrency-go-api/
├── cmd/server          # Ponto de entrada da aplicação
├── graph               # Schema e resolvers do GraphQL
├── internal/
│   ├── api             # Cliente REST para usuários e posts
│   ├── handler         # Handlers HTTP
│   ├── model           # Modelos de dados
│   ├── router          # Configuração de rotas HTTP e GraphQL
│   └── service         # Lógica de agregação e serviços
├── template            # Templates HTML
├── go.mod
└── go.sum
```

## Tecnologias Utilizadas

* Go 1.24
* gqlgen
* Resty
* HTTP Server nativo do Go

## Configuração

1. Clone o repositório:

```bash
git clone https://github.com/seuusuario/go-concurrency-api.git
cd go-concurrency-api
```

2. Build e run:

```bash
docker build -t go-api .
docker run -p 8080:8080 go-api
```

3. O servidor vai rodar em `http://localhost:8080`.

## Endpoints

* `GET /` - Página inicial HTML
* `GET /user` - Lista de usuários (JSON)
* `GET /post` - Lista de posts (JSON)
* `POST /query` - Endpoint GraphQL
* `/playground` - Playground GraphQL

### GraphQL Query Exemplo

```graphql
query {
  userSummary(userID: 1) {
    name
    email
    postCount
  }
}
```

## Testes

O projeto possui testes unitários e de integração:

```bash
go test ./... -cover
```
