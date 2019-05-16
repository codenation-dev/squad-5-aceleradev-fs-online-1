# Gestão de clientes Banco Uati

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
