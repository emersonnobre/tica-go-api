# tica-go-api
## Pacotes necessários para rodar esse projeto:
* make
* docker
* docker-compose
## Como rodar:
* Na raíz do projeto, execute: `make run-container` se estiver usando Linux. Caso esteja usando Windows, execute: `docker-compose up -d`
* Acesse o swagger em http://localhost:3000/swagger e utilize a API
## Para o ambiente de desenvolvimento, algumas tabelas já vem com dados:
### Tabelas que possuem dados fictícios para desenvolvimento e testes:
* Categorias (categories)
* Funcionários (employees)
* Produtos (products)
* Clientes (cutomers)
* Endereços (addresses)
**Obs**: Para verificar tais dados, basta consultar os respectivos endpoints. 