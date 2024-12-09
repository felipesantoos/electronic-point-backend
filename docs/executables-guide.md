# GUIA DE USO DOS EXECUTÁVEIS DO PROJETO
## COMO TESTAR OS EXECUTÁVEIS ENVIRONMENT
### TESTANDO EXECUTÁVEL `development.sh`
Esse executável é responsável por construir todo o ambiente do backend para subir o servidor e analisar as rotas do swagger.
1. Na pasta raiz do projeto, execute o comando `./execute.sh -environment -development`;
2. Aguarde a finalização dos comandos pelo arquivo executável;
3. Pronto, agora você tem a aplicação rodando.

### TESTANDO EXECUTÁVEL `test.sh`
Esse executável é responsável por construir todo o ambiente de testes e executa-los para facilitar o processo de revisões de código pelos devs.
1. Na pasta raiz do projeto, execute o comando `./execute.sh -environment -test`
2. Aguarde a finalização dos comandos pelo arquivo executável;
3. Pronto, agora você sabe se os testes estão passando.

### TESTANDO `development.sh` E `test.sh`
Para facilitar ainda mais o processo, é possível verificar os testes e logo em seguida construir o backend da aplicação para observar as funcionalidades no swagger.

**OBS.:** _Caso os testes sejam mal-sucedidos, o backend não é construído!_
1. Na pasta raiz do projeto, execute o comando `./execute.sh -environment -all`
2. Aguarde a finalização dos comandos pelo arquivo executável;
3. Pronto, agora você sabe se os testes estão passando e tem a aplicação rodando.

## COMO TESTAR EXECUTÁVEL MOCKGEN
Esse executável é responsável por gerar os mocks da camada de API e serviço da aplicação.
### TESTANDO EXECUTÁVEL `mockgen.sh`
1. Na pasta raiz do projeto, execute o comando `./execute.sh -mockgen -usecases`
2. Aguarde a finalização dos comandos pelo arquivo executável;
3. Pronto, agora você tem todos os mocks da camada de API.

**OBS.:** _Para gerar mocks específicos, use o comando `./execute.sh -mockgen -usecases -nome1 -nome2...`_

1. Na pasta raiz do projeto, execute o comando `./execute.sh -mockgen -adapters`
2. Aguarde a finalização dos comandos pelo arquivo executável;
3. Pronto, agora você tem todos os mocks da camada de serviço.

**OBS.:** _Para gerar mocks específicos, use o comando `./execute.sh -mockgen -adapters -nome1 -nome2...`_

1. Na pasta raiz do projeto, execute o comando `./execute.sh -mockgen -all`
2. Aguarde a finalização dos comandos pelo arquivo executável;
3. Pronto, agora você tem todos os mocks.
