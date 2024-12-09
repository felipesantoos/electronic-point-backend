# Preparando o ambiente de testes

Como sabemos, testes desempenham um papel fundamental no desenvolvimento de software, pois garantem a qualidade, confiabilidade e desempenho do produto final.

Logo abaixo, segue o passo a passo de como configurar o ambiente de testes no `Pedagoga Backend`.

## Passo 1° - Subir o contêiner com o banco de dados de teste

1. Antes de tudo, utilizaremos o comando `source ./src/utils/tests/envs/.env.test` para definir as variáveis de ambiente de testes.

2. E para subir o contêiner com o banco de dados postgres, use o comando: `docker compose -f ./src/utils/tests/docker/docker-compose.test.yml up database_test --build -d`

Logo após, para garantirmos que está tudo funcionando, utilize o comando: `docker ps` para listar todos os contêineres em execução. Haverá uma coluna com o nome `CONTAINER ID`. Copie o ID e utilize o comando: `docker exec -it <CONTAINER ID> bash`, para acessar o terminal shell do contêiner.

Com o terminal do contêiner disponível, use: `psql -U pedagoga_test -p 5432` para acessar o terminal do PostgreSQL e logo após utilize: `\dt` para listar as tabelas da base de dados. Provavelmente, você receberá uma mensagem como esta: `Did not find any relations.` Resolveremos isso no próximo passo.

Para sair dos terminais: `\q` para o PostgreSQL e `exit` para o contêiner

Se tiver problemas como este: `Error response from daemon: driver failed programming external connectivity on endpoint pedagoga_database (2c3dc61f5275053c4d88c8116de683c55d88a00c83636431e11d375ce032da0c): Error starting userland proxy: listen tcp4 0.0.0.0:5432: bind: address already in use`.

Uma possível solução é utilizar o `sudo netstat -pan | grep 5432` para listar os processos que estão rodando na porta `5432`. Copie o ID do processo e finalize com o `sudo kill <ID>`

## Passo 2° - Rodar as migrations

1. Para carregar os dados das fixtures: `migrate -source file://config/database/postgres/migrations -database postgres://pedagoga_test:12345678_test@localhost:5433/pedagoga_test?sslmode=disable up`

Caso tenha problemas como este: `Dirty database version 1. Fix and force version`. Repita o mesmo comando acima, substituindo o `up` por `force 1`

## Passo 3° - Criar os testes

Para cada funcionalidade que você for desenvolver, é essencial ter pelo menos 2 casos de testes em cada camada, um para sucesso e outro para falha.

Se a sua funcionalidade se comunicar com o banco de dados e não existirem testes para determinada funcionalidade, crie o teste de integração na camada de `repositório` em `./src/infra/repository`.

Caso seja necessário criar uma nova entidade, ela também deverá estar aliada aos testes unitários, que deverão ficar na camada de `domínio` em `./src/core/domain`.

Para os testes de ponta a ponta, você terá que usar um comando adicional para gerar os mocks na camada de API com: `mockgen -source=./src/core/interfaces/primary/<NOME_DO_ARQUIVO_DE_ENTRADA>.go -destination=./src/apps/api/handlers/mocks/<NOME_DO_ARQUIVO_DE_SAÍDA>.go -package=mocks`

Para a camada de serviço: `mockgen -source=./src/core/interfaces/secondary/<NOME_DO_ARQUIVO_DE_ENTRADA>.go -destination=./src/core/services/mocks/<NOME_DO_ARQUIVO_DE_SAÍDA>.go -package=mocks`.

Caso não tenha o mockgen instalado, utilize: `go install go.uber.org/mock/mockgen@latest`.

## Passo 4° - Rodar os testes

Se quiser rodar todos os testes de uma só vez: `go test ./...`

Para pacotes específicos: `go test ./<URL>`

Para funções específicas: `go test -run <NOME_DA_FUNÇÃO> ./<URL>`
