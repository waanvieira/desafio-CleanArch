# Sistema simplificado de transferências 
<p>
<a href="https://github.com/waanvieira/desafio-CleanArch/blob/main/LICENSE"><img src="https://img.shields.io/packagist/l/laravel/framework" alt="License"></a>
</p>

# Sobre o projeto
Projeto para criar e listar orders usando clean architecture utiliznado REST API/GRAPHQL/GRPC
- Endpoint REST (GET /order), POST (POST/ order)
- Service CreateOrder, ListOrders com GRPC
- Query CreateOrder, ListOrders GraphQL

# Tecnologias utilizadas
- 1.22.5
- MYSQL 5.7
- RabbitMQ

# Como executar o projeto

## Pré-requisitos
Docker
https://www.docker.com/get-started/

```bash
# clonar repositório
git clone https://github.com/waanvieira/desafio-CleanArch.git

# entrar na pasta do projeto back end
cd desafio-CleanArch

# executar o projeto
docker-compose up -d

#Criar migrations
make migrate

#Excutar o GO

cd cmd/orderSystem

go run main.go wire_gen.go 

```

# Uso do sistema

REST API
POST http://localhost:8000/order HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
    "price": 100.5,
    "tax": 0.5
}

GET http://localhost:8000/orders HTTP/1.1
Host: localhost:8000
Content-Type: application/json

GRAPHQL
http://localhost:8080/

mutation createTodo {
  createOrder(
    input: {
      Price: 10,
      Tax: 10 }
  	) {
    id
    Price
    Tax
    FinalPrice
  }
}

query listOrders {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}

GRPC

evans -r repl

Mostrar nosso pacotes, temos que escolher um pacore para utilizar
show package

Mostra nossos service
show service

Selecionar o package
package pb

Para executarmos uma chamada
service OrderService

Depois de selecionar o service, chamar o método

Criar Order - call CreateOrder
Listar Orders - call ListOrders



# Autor

Wanderson Alves Vieira

https://www.linkedin.com/in/wanderson-alves-vieira-59b832148
