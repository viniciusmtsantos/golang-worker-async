# Message Broker em Golang

Este projeto implementa um **Message Broker** utilizando o pacote **async** em Golang. O sistema é desacoplado, onde as interações são feitas por meio de um fluxo de mensageria **pub/sub** via Redis, com a integração de **gRPC** para invocar o processo de trabalho (workers).

## Arquitetura

O fluxo de mensagens é o seguinte:
1. **gRPC API**: A API recebe chamadas para iniciar o processamento.
2. **Message Broker**: A comunicação é realizada através de um sistema pub/sub usando **Redis**.
3. **Workers**: Os workers escutam os tópicos de mensagens no Redis, processando-as conforme necessário.

## Funcionalidades
- **Desacoplamento**: Cada componente opera de forma independente, promovendo escalabilidade e flexibilidade.
- **Mensageria Pub/Sub**: Redis como canal para garantir a entrega eficiente das mensagens entre os componentes.
- **gRPC**: Usado para comunicação rápida e eficiente entre a API e os workers.

## Como rodar o projeto

1. Instalar dependências: 
    ```bash
    go mod tidy
    ```
2. Subir o servidor Redis.
3. Executar o servidor:
    ```bash
    go run main.go
    ```

---

### Diagrama da Arquitetura

O diagrama abaixo ilustra a estrutura de como os componentes interagem:

```
 +------------+       +------------+        +------------+
 |  gRPC API  | <---> |  Redis Pub | <--->  |   Workers  |
 +------------+       +------------+        +------------+
       ^                     |
       |                     v
 Chama o fluxo       Processa mensagens
```

Neste diagrama, a **gRPC API** inicia o fluxo de mensagens, o **Redis Pub/Sub** distribui as mensagens para os **Workers**, que as processam conforme o necessário. O sistema é escalável e facilmente desacoplado.
