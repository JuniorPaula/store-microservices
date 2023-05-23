# Projeto de Microserviços em Golang

Este projeto consiste em uma arquitetura de microserviços desenvolvida em Golang. O sistema é composto por cinco microserviços interconectados: `product`, `catalog`, `checkout`, `order` e `payment`. Cada microserviço desempenha um papel específico na aplicação, facilitando a escalabilidade e manutenção do sistema.

## Microserviço `product`

O microserviço `product` é responsável por fornecer uma lista de produtos disponíveis por meio de uma API HTTP. Ele responde às solicitações com informações detalhadas sobre os produtos, como nome, descrição, preço, etc.

## Microserviço `catalog`

O microserviço `catalog` comunica-se com o microserviço `product` por meio de requisições HTTP para obter a lista dos produtos disponíveis. Ele atua como um catálogo centralizado, armazenando informações sobre os produtos e fornecendo uma visão consolidada dos mesmos.

## Microserviço `checkout`

O microserviço `checkout` é responsável por gerar o processo de checkout para um produto específico. Ele se comunica com o microserviço `product` para obter os detalhes do produto desejado e, em seguida, envia esses dados para uma fila do RabbitMQ, que será consumida posteriormente.

## Microserviço `order`

O microserviço `order` recebe os dados do pedido por meio da fila do RabbitMQ. Ao consumir as mensagens publicadas pelo `checkout`, ele salva as informações do pedido no banco de dados utilizando o Redis. Esse microserviço é responsável por gerenciar os pedidos realizados pelos usuários.

## Microserviço `payment`

O microserviço `payment` é responsável pelo processamento dos dados do pedido e pela publicação do status de aprovação em uma fila de pagamento. Ele recupera os dados da ordem da fila, realiza o processamento de pagamento e, em seguida, publica o status de aprovação na fila correspondente.

## Integração entre microserviços

Os microserviços são interconectados por meio de chamadas HTTP e utilizam a fila do RabbitMQ para comunicação assíncrona entre os serviços de `checkout`, `order` e `payment`. O Redis é utilizado para armazenar informações sobre os pedidos e status de pagamento, garantindo a consistência e recuperação de dados.

## Dependências

- Golang
- RabbitMQ
- Redis

Certifique-se de ter as dependências acima instaladas e configuradas corretamente antes de executar o projeto.

## Executando o projeto

1. Clone o repositório para sua máquina local.
2. Navegue para o diretório raiz do projeto.
3. Execute cada microserviço individualmente, seguindo as instruções de cada um deles.
4. Certifique-se de iniciar o RabbitMQ e o Redis antes de executar os microserviços.

## Considerações finais

Este projeto é apenas um exemplo de arquitetura de microserviços em Golang. Você pode expandi-lo, adicionar novas funcionalidades e melhorar sua robustez de acordo com os requisitos do seu sistema. Sinta-se à vontade para personalizar e adaptar o projeto conforme suas necessidades.
