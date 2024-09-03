# rate-limiter
projeto para o rate limiter do full cycle

# Projeto Go - Arquitetura Hexagonal

## Visão Geral

Este projeto implementa uma aplicação Go seguindo o padrão de arquitetura hexagonal. Essa abordagem visa criar um sistema bem estruturado, com alta coesão e baixo acoplamento, facilitando a manutenção e evolução do software.

## Estrutura do Projeto
![Estrutura do Projeto](imagens/estrutura_do_projeto.png)


* **cmd:** Ponto de entrada da aplicação.
* **internal:** Contém a lógica interna:
  * **core:** Núcleo da aplicação, com regras de negócio e entidades.
  * **adapters:** Adaptadores para diferentes tecnologias (HTTP, banco de dados, etc.).
* **config:** Arquivos de configuração.
* **go.mod:** Gerenciamento de dependências.

## Conceitos-chave

* **Hexagonal Architecture:** Separa a lógica de negócio da infraestrutura, permitindo testes isolados e maior flexibilidade.
* **Ports & Adapters:** Definem a interface entre o core e o mundo externo.
* **Use Cases:** Representam as ações que o sistema pode realizar.
* **Entities:** Representam os objetos do domínio.
