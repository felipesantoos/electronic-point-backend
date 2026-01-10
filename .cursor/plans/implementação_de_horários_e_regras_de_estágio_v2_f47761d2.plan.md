---
name: Implementação de Horários e Regras de Estágio v2
overview: Implementar horários de entrada/saída nos estágios, validar tolerância de 30min e limitar o total de horas trabalhadas a 5h por dia (somando todos os registros do dia), além de melhorias na API e segurança.
todos:
  - id: db-migration-schedule-v2
    content: Criar migração SQL para adicionar horários na tabela internship
    status: pending
  - id: domain-update-internship-v2
    content: Atualizar domínio de Internship com scheduleEntryTime e scheduleExitTime
    status: pending
    dependencies:
      - db-migration-schedule-v2
  - id: repo-update-internship-v2
    content: Atualizar repositório Postgres para persistir e recuperar os horários do estágio
    status: pending
    dependencies:
      - domain-update-internship-v2
  - id: service-logic-validation-v2
    content: Implementar validação de tolerância (30min) e limite diário de 5h no serviço de TimeRecord
    status: pending
    dependencies:
      - repo-update-internship-v2
  - id: dto-response-enrichment-v2
    content: Atualizar DTOs de resposta para incluir informações de estudante/estágio no TimeRecord
    status: pending
    dependencies:
      - service-logic-validation-v2
  - id: security-cleanup-docs-v2
    content: Remover exemplos de senhas e dados sensíveis da documentação Swagger (DTOs)
    status: pending
---

# Plano de Implementação: Horários de Estágio e Validações (v2)

Este plano visa atender às solicitações de associar horários aos estágios, implementar validações de ponto e melhorar a segurança/eficiência da API.

## 1. Banco de Dados

- Criar migração para adicionar `schedule_entry_time` e `schedule_exit_time` (tipo `TIME`) na tabela `internship`.
- Arquivo: `config/database/postgres/migrations/000004_add_schedule_to_internship.up.sql`

## 2. Domínio (Internship)

- Atualizar interface e struct de `internship` com os novos campos.
- Atualizar builder de `internship`.
- Arquivos: `src/core/domain/internship/interface.go`, `src/core/domain/internship/internship.go`, `src/core/domain/internship/builder.go`

## 3. Repositório (Postgres)

- Atualizar queries SQL (Insert, Update, Select) e mapeamento (queryObject) para incluir os novos campos de horário.
- Arquivos: `src/infra/repository/postgres/query/internship.go`, `src/infra/repository/postgres/queryObject/internship.go`, `src/infra/repository/postgres/internship.go`

## 4. Lógica de Serviço (TimeRecord)

- Implementar validações no `Create` do `timeRecordServices`:
    - Buscar estágio atual do estudante para obter o horário previsto.
    - **Tolerância**: Validar se o `entry_time` está dentro da tolerância de 30min do horário agendado.
    - **Limite Diário de 5h**: 
        - Buscar todos os registros do estudante na mesma data.
        - Somar a duração dos registros existentes.
        - Se o total já atingiu 5h, retornar erro.
        - Se o novo registro fizer o total exceder 5h, ajustar o `exit_time` do novo registro para "cortar o excesso" e cravar o total do dia em 5h.
- Arquivo: `src/core/services/timeRecord.go`

## 5. API e DTOs

- Atualizar DTOs de requisição e resposta de `internship` e `timeRecord`.
- Incluir dados do estudante/estágio na resposta de `timeRecord` para evitar múltiplas requisições.
- Segurança: Remover `example` de senhas nos DTOs (Swagger).
- Arquivos: `src/apps/api/handlers/dto/request/internship.go`, `src/apps/api/handlers/dto/response/timeRecord.go`, `src/apps/api/handlers/dto/request/credentials.go`