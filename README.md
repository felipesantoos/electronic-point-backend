# API de Ponto Eletrônico

## O que é o projeto?

Esse é um projeto que utiliza a arquitetura hexagonal e, até o momento, tem uma única porta de entrada: um projeto BackEnd que utiliza o Framework [Echo](https://echo.labstack.com/).

---

## Como executar?

Os passos abaixo assumem que você já fez/tem:
1. O clone (`git clone`) do projeto no seu computador e um terminal aberto na pasta baixada.
   
### **Para Desenvolvimento**
Se você é um membro do time de desenvolvimento desse projeto, siga os passos abaixo para executar as configurações apropriadas:

1. Certifique-se de que você possui a ferramenta CLI do Go instalada ([instruções de instalação](https://go.dev/learn/));
2. Certifique-se de que o Docker esteja instalado no seu computador ([instruções de instalação](https://www.docker.com/));
3. Copie todo o conteúdo do arquivo `.env.example` e cole em um novo arquivo chamado `.env` na raiz do projeto;
4. Execute o seguinte script para configurar e iniciar o projeto automaticamente ou apenas execute os comandos `chmod +x execute.sh` e `./execute.sh -environment -development` (o script `execute.sh` irá executar todos os comandos abaixo):

```bash
#!/bin/bash

# Load the environment variables
source .env
schema=$(echo $DATABASE_SCHEMA | sed "s/\r//")
user=$(echo $DATABASE_USER | sed "s/\r//")
password=$(echo $DATABASE_PASSWORD | sed "s/\r//")
host=$(echo $DATABASE_HOST | sed "s/\r//")
port=$(echo $DATABASE_PORT | sed "s/\r//")
name=$(echo $DATABASE_NAME | sed "s/\r//")
ssl_mode=$(echo $DATABASE_SSL_MODE | sed "s/\r//")
migrations_path=$(echo $DATABASE_MIGRATIONS_PATH | sed "s/\r//")
uri="$schema://$user:$password@$host:$port/$name?sslmode=$ssl_mode"

# Start the databases
docker compose -f docker-compose.dev.yml up database redis --build -d

# Download the project dependencies
go mod tidy

# Generate the API documentation
bash -c "cd src/apps/api && swag init -g ../main.go --output ./docs --dir ./handlers"

# Wait 5 seconds so that the database can initiate and then load the migrations
migrate -path $migrations_path -database $uri up

# Start the server
go run src/apps/api/main.go
```

### **Para Testes de Qualidade (QA)**
Se você é um membro do time de qualidade (QA), siga os passos abaixo para executar as configurações apropriadas:

1. Copie todo o conteúdo do arquivo `.env.example` e cole em um novo arquivo chamado `.env` na raiz do projeto;
2. Execute o seguinte comando para iniciar os serviços necessários:

```bash
docker compose -f docker-compose.dev.yml up --build
```

O projeto estará disponível no endereço [http://localhost:8000](http://localhost:8000).
