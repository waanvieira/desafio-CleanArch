# Apelido do comando
createmigration:
	migrate create -ext=sql -dir=internal/infra/database/migrations -seq init

migrate:
	migrate -path=internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

migratedown:
	migrate -path=internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down
# PHONY serve para eliminar a possibilidade de executar outro comando que esteja na raiz do projeto, criando assim uma lista para executar os comandos apenas aqui
.PHONY: migrate migratedown createmigration