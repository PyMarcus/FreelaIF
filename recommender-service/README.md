# recommender-service

Consome mensagens da fila `user.created` do RabbitMQ, gera recomendações simuladas, salva no PostgreSQL e publica na fila `email.send`.

## Exemplo de mensagem recebida em `user.created`
```json
{
  "user_id": "123"
}
```

## Exemplo de mensagem publicada em `email.send`
```json
{
  "subject": "Assunto", 
  "email": "usuario@exemplo.com",
  "nome": "Usuario",
  "conteudo_html": "<></>"
}
```
