# Gestão de clientes Banco Uati

[![Go Report Card](https://goreportcard.com/badge/github.com/codenation-dev/squad-5-aceleradev-fs-online-1)](https://goreportcard.com/report/github.com/codenation-dev/squad-5-aceleradev-fs-online-1)

## Objetivo

O objetivo deste produto é monitorar e gerar alertas da captura de uma determinada fonte com base em uma determinada base do cliente e regra pré estabelecida.

## Contextualização

O Banco Uati gostaria de monitorar de forma contínua e automatizada caso um de seus clientes vire um funcionário público do estado de SP (http://www.transparencia.sp.gov.br/busca-agentes.html) ou seja um bom cliente com um salário maior que 20 mil reais.

A lista de clientes do banco Uati encontra-se no arquivo ``clientes.csv`` contido neste projeto.


## Requisitos técnicos obrigatórios

- Tela de login;
- Uma tela para cadastrar os usuários que devem receber os alertas;
- Uma tela para importação dos clientes do banco (Upload de CSV);
- Uma tela para controle do monitoramento/dashboard com nº de alertas e outras funcionalidades que o grupo julgar interessantes;
- Uma tela para listar e detalhar os alertas,  listar os envios de emails e para quem foi enviado, data, hora e outras funcionalidades que o grupo julgar interessantes;
- Enviar um alerta através de e-mail quando um cliente se tornar um funcionário do banco;
- Todas essas funcionalidades devem ser expostas para clientes que queiram integrar através de uma API.


## Pré-requisitos
- Instalar 
  - Docker https://docs.docker.com/install/
  - Docker-compose https://docs.docker.com/compose/install/
- Configurar para linux
  - https://docs.docker.com/install/linux/linux-postinstall/

## Ambiente de desenvolvimento

- Start do banco de dados
  - docker-compose up -d
- Stop do banco de dados
  - docker-compose down
- Start da documentação OpenAPI/Swagger
  - docker-compose up -d swagger-ui
  - Acessar a URL http://localhost:8082/
- Rodar os testes unitários 
  - go test ./...
- Rodar os testes unitários com o relatório do cobertura
  - go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## Servidor de desenvolvimento
- Frontend http://bancouati.ga/
  - Exemplo: http://bancouati.ga/dashboard
- Backend http://bancouati.ga/api/
  - Exemplo GET http://bancouati.ga/api/users
- Swagger http://bancouati.ga/docs/
