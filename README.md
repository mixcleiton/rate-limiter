## Rate Limiter em Go

Objetivo
Desenvolver um rate limiter em Go capaz de controlar o fluxo de requisições a um serviço web, limitando o número de requisições por segundo por IP ou por token de acesso.

Descrição
Este projeto implementa um mecanismo de rate limiting em Go, permitindo configurar limites de requisições por segundo para cada endereço IP ou token de acesso. O rate limiter atua como um middleware, interceptando as requisições e verificando se o limite foi excedido.

Características
Flexibilidade: Permite configurar limites por IP ou token de acesso.
Personalização: Permite configurar o tempo de bloqueio após exceder o limite.
Configuração: As configurações são feitas através de variáveis de ambiente ou arquivo .env.
Persistência: Utiliza Redis para armazenar as informações de rate limiting.
Escalabilidade: A arquitetura permite a troca do mecanismo de persistência.
Tratamento de erros: Retorna o código de status HTTP 429 e mensagem de erro apropriada quando o limite é excedido.
Testes: Possui testes unitários e de integração para garantir a qualidade do código.
Arquitetura
Middleware: Intercepta as requisições HTTP e verifica se o limite foi excedido.
Lógica de rate limiting: Implementa a lógica de verificação de limites e atualização do estado no Redis.
Persistência: Utiliza Redis para armazenar as informações de rate limiting.
Configuração: Carrega as configurações de um arquivo .env ou variáveis de ambiente.
Instalação e Execução
Pré-requisitos:
Go
Docker
Redis
Clonar o repositório:
Bash
git clone https://seu-repositorio.git


Criar um arquivo .env:
# Exemplo de arquivo .env

REDIS_HOST=localhost
REDIS_PORT=6379
LIMIT_PER_IP=10
LIMIT_PER_TOKEN=50
BLOCK_DURATION=5m

Construir e executar:
Bash
docker-compose up --build

Uso
Limitação por IP: O rate limiter verifica o endereço IP de cada requisição e compara com o limite configurado.
Limitação por token: O token de acesso deve ser enviado no header da requisição no formato API_KEY: <TOKEN>. O rate limiter verifica o token e compara com o limite configurado.
Testes
Para executar os testes, utilize o comando:

Bash
go test ./...