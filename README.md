# < Credentials >
<fig>
<img src="https://nordicapis.com/wp-content/uploads/10-Login-APIs.png" alt="tela de login">
</fig>

## Inicialização
Para executar o projeto, utilize as ferramentas descritas na sessão *Ferramentas*.

## Ferramentas
* [Neovim](https://neovim.io/) / [Vscode](https://code.visualstudio.com/) - Editor de texto para desenvolvimento.
* [Docker](https://www.docker.com/) - Ferramenta para criar conteiners e facilitar o build da aplicação.
* [Make](https://embarcados.com.br/introducao-ao-makefile/) -  Ferramenta para scriptar comandos

## Links importantes
* [Go](https://go.dev/) -  Linguagem usada no projeto.
* [Gin](https://gin-gonic.com/) -  Package usado para criação da API.
* [Jwt](github.com/golang-jwt/jwt) - Package para gerar jwt para rotas 
* [Godotenv](https://github.com/joho/godotenv) - Package para acessar variavies de ambiente
* [Crypto](https://golang.org/x/crypto) - Package para cryptografia
* [Redis](https://redis.io/try-free/?utm_campaign=gg_s_brand_bam_acq_amert2-en&utm_source=google&utm_medium=cpc&utm_content=&utm_term=&gad_source=1&gclid=CjwKCAjw34qzBhBmEiwAOUQcF38U_Ub03TluL2jgYftiKBAjk7npGtUOaPkOHIS2xXnEIjggf19DJRoCDYUQAvD_BwE) - Banco de dados
  
# < Credentials >

## Introdução

> Sistema desenvolvidor para atuar como atuar como um middleware em outros projetos com foco em authenticação do usuário. 

Este projeto possui o objetivo principal validar o acesso a quaisquer projeto.  
Com os objetivos gerais de realizar a inserção de credenciais com password cryptografado, rotas com aplicação de jwt e observabilidade aplicada na solução. 

## Análise técnica

### Descrição do ambiente técnico

O sistema é composto por um banco de dados e uma api. Funcionalidades principais:

* **F1** - Manipular credencial de acesso.
* **F2** - Validar acesso por token.
* **F3** - Criar e verificar cryptografia de dados sensíveis.
* **F4** - Persistir dados.

As ferramentas utilizadas para o desenvolvimento incluem **Golang** que é uma linguagem de programação utilizada para o Back-end. **Redis** atuando como sistema gerenciador de banco de dados e **Docker** para utilizar o ambiente em container.

### Requisitos Funcionais
Respeitando a proposta, o sistema deverá atender os seguintes requisitos:

* **RF1** - Autenticação e Autorização com JWT.
* **RF2** - Cryptografia de dados sensíves.
* **RF3** - Logs de acesso e error.


### Mensagens internas

Rotas utilizadas pela aplicação web para executar metodos de **POST** e **GET** no banco de dados. Onde o retorno de cada uma das funções estara contido em uma sessão para renderização de páginas web.

| Nome | Funcionalidade|
|------|--------------|
|```GET``` /api/v1/read|Retorna todas as credencials cadastradas.|
|```POST``` /api/v1/login|Verifica a autenticidade da credencial e caso esteja de acordo retorna token de acesso.|
|```POST``` /api/v1/create|Insere uma nova credencial.|
|```DELETE``` /api/v1/:email|Deleta uma credencial pelo email.|
|```PUT``` /api/v1/updatePassword|Altera o password de uma credencial.|
