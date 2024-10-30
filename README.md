# Base Project DDD with Golang

Este projeto serve como uma base para aplicações construídas com o padrão Domain-Driven Design (DDD) em Go.

## Descrição

A ideia deste projeto é ser uma base de um projeto que já separa o domínio da infraestrutura e implementa um sistema de ordem de compra com entidades `Customer`, `Product`, `Order`, e `OrderItem`.

O projeto está organizado da seguinte 

```base-project-ddd-with-golang/
│
├── cmd/                     # Ponto de entrada da aplicação
│
├── domain/                  # Camada de Domínio (entidades e lógica de negócio)
│   ├── checkout,
├   ├   customer,
├   ├   product/             # Contexto geral
│   │   ├── entity/          # Entidades do domínio (Order, Customer, Product)
│   │   ├── Factory/         # Função que cria novas instâncias de suas respectivas entidades
│   │   ├── repository/      # Interfaces para repositórios do domínio
│   │   └── service/         # Serviços e lógica de domínio
│
├── infrastructure/          # Camada de Infraestrutura (conexão com BD, APIs)
│   ├── repository/          # Implementações dos repositórios para BD
│   └── models/              # Modelos do GORM para BD
│
└── README.md                # Documentação ```


### Principais Diretórios e Arquivos

- `domain/`: Contém as entidades, repositórios e objetos de valor.
- `infrastructure/`: Contém os modelos ORM e implementações de repositórios.
- `main.go`: O ponto de entrada da aplicação.

## Configuração do Projeto

### Pré-requisitos

- [Go](https://golang.org/doc/install) 1.16 ou superior
- [GORM](https://gorm.io/)

### Instalação

1. Clone o repositório:

```git clone https://github.com/PierryMedeiros/ddd-go-aplication.git```

2. Vá até o diretório do projeto:

```cd base-project-ddd-with-golang```

3. Instale as dependências:

```go mod tidy```

4. Rode os testes

```go test ./...```

Contribuição
Faça um fork do projeto.

Crie um branch para sua feature (git checkout -b feature/nova-feature).

Commit suas mudanças (git commit -am 'Adiciona nova feature').

Faça push para o branch (git push origin feature/nova-feature).

Crie um novo Pull Request.