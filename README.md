# FreelaIF

## Visão Geral

**FreelaIF** é uma plataforma de freelance criada para conectar estudantes do Instituto Federal a clientes da cidade e região, oferecendo oportunidades reais de trabalho, geração de renda e experiência profissional. A iniciativa visa também fortalecer o empreendedorismo estudantil e aproximar o ambiente acadêmico do mercado local.

## Objetivos

O projeto tem como principais objetivos:

- Proporcionar aos estudantes a divulgação de suas habilidades e a realização de serviços freelance.
- Permitir que empresas e pessoas físicas encontrem mão de obra qualificada e acessível.
- Incentivar a integração da comunidade acadêmica com o mercado regional.

Além disso, a plataforma contará com recursos de:

- **Recomendações inteligentes** de serviços para freelancers, baseadas em palavras-chave extraídas dos perfis e habilidades cadastradas.
- **Sugestões automáticas** de freelancers mais adequados para clientes, de acordo com os requisitos dos projetos.
  
## Funcionamento

1. **Cadastro**: Estudantes e clientes realizam seu registro e criam perfis personalizados.
2. **Publicação de Projetos**: Clientes anunciam tarefas, especificando prazos e valores.
3. **Candidatura**: Estudantes se candidatam às oportunidades disponíveis.
4. **Seleção e Execução**: Clientes escolhem o freelancer ideal para o projeto, que é então executado conforme as especificações.
5. **Entrega e Validação**: Após a entrega do serviço, o cliente valida o trabalho.
6. **Pagamento**: O pagamento é liberado ao freelancer, com uma taxa de intermediação retida pela plataforma.

## Funcionalidades Principais

- Cadastro e autenticação de estudantes e clientes.
- Publicação, edição e acompanhamento de tarefas/projetos.
- Candidatura a tarefas por parte dos estudantes.
- Sistema de notificações em tempo real para atualizações de status e mensagens.
- Avaliação de freelancers após a conclusão dos projetos.
- Gestão de pagamentos integrados, com retenção automática da comissão da plataforma.
- Painéis personalizados para estudantes, clientes e administradores.
- Recomendações automatizadas de freelancers e projetos baseadas em palavras-chave.

## Monetização

O FreelaIF será sustentado pela retenção de uma porcentagem sobre o valor de cada serviço concluído, garantindo a operação e expansão contínua da plataforma.

## Público-Alvo

- **Estudantes** do Instituto Federal que desejam atuar como freelancers.
- **Empresas** e **pessoas físicas** da região que buscam mão de obra qualificada para projetos de curta ou média duração.
- **Administradores da plataforma** que gerenciam usuários, tarefas e taxas.

## Requisitos do Sistema

### Funcionais

| ID | Descrição |
|----|-----------|
| RF01 | Cadastro e autenticação de estudantes e clientes. |
| RF02 | Publicação e gestão de projetos. |
| RF03 | Sistema de candidaturas para freelancers. |
| RF04 | Processo de seleção e contratação por parte do cliente. |
| RF05 | Controle de status de execução dos projetos. |
| RF06 | Envio de notificações de atualizações relevantes. |
| RF07 | Integração de pagamento e retenção de comissão. |
| RF08 | Sistema de avaliação de freelancers. |
| RF09 | Painel administrativo para gestão da plataforma. |
| RF10 | Módulo de recomendações por palavras-chave. |

### Não Funcionais

| ID | Descrição |
|----|-----------|
| RNF01 | Responsividade para dispositivos móveis. |
| RNF02 | Autenticação segura utilizando JWT e OAuth2. |
| RNF03 | Notificações em tempo real utilizando MQTT. |
| RNF04 | Tempo de resposta inferior a 2 segundos por operação. |
| RNF05 | Alta disponibilidade do sistema (>99% de uptime). |
| RNF06 | Criptografia de dados sensíveis. |

## Restrições

- A plataforma será disponibilizada inicialmente apenas para estudantes do Instituto Federal local.
- O sistema deverá ter suporte prioritário a navegadores modernos (Google Chrome, Mozilla Firefox, Microsoft Edge).
- Integração de pagamento será feita utilizando soluções consolidadas no mercado, como MercadoPago ou API Pix.

## Roadmap Futuro

- Desenvolvimento de um aplicativo móvel para Android e iOS.
- Expansão para atender estudantes de outras instituições.
- Implementação de sistema de análise automatizada da qualidade dos projetos antes da entrega final.
- Melhoria contínua dos algoritmos de recomendação.

## Contribuidores
- [Gessyca Moreira](https://github.com/Geessyca)
- [Marcus Vinicius](https://github.com/PyMarcus)
