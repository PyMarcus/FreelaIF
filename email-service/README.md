# email-service

Consome mensagens da fila `email.send` do RabbitMQ e simula o envio de e-mails.

## Exemplo de mensagem consumida em `email.send`
```json
{
  "subject": "Assunto", 
  "email": "usuario@exemplo.com",
  "nome": "Usuario",
  "conteudo_html": "<></>"
}
```
