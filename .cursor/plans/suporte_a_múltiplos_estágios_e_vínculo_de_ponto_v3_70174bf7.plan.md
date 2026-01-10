---
name: Suporte a Múltiplos Estágios e Vínculo de Ponto v3
overview: Adicionar suporte para que um estudante tenha múltiplos estágios simultâneos, vinculando cada registro de ponto a um ID de estágio específico e validando as regras de negócio com base nesse vínculo.
todos:
  - id: db-migration-multi-internship-v3
    content: Criar migração SQL e atualizar fixtures CSV para incluir internship_id no ponto
    status: completed
  - id: domain-update-time-record-v3
    content: Atualizar domínio e builder de TimeRecord com o campo internshipID
    status: completed
    dependencies:
      - db-migration-multi-internship-v3
  - id: dto-update-request-v3
    content: Adicionar internship_id como obrigatório no DTO de criação de ponto
    status: completed
    dependencies:
      - domain-update-time-record-v3
  - id: repo-update-persistence-v3
    content: Atualizar queries (Select com INNER JOIN) e repositório para persistir internship_id
    status: completed
    dependencies:
      - domain-update-time-record-v3
  - id: service-refactor-logic-v3
    content: Refatorar serviço TimeRecord com validação de posse do estágio e novas regras
    status: completed
    dependencies:
      - repo-update-persistence-v3
      - dto-update-request-v3
---

# Plano de Implementação Detalhado: Suporte Multi-Estágio (v3)

Este plano detalha a transição do sistema para suportar múltiplos estágios ativos por estudante, garantindo que cada batida de ponto esteja vinculada a um contrato de estágio específico.

## 1. Banco de Dados & Fixtures

- **Migração SQL**: 
    - Criar `config/database/postgres/migrations/000005_add_internship_id_to_time_record.up.sql`:

        1. Adicionar coluna `internship_id` (UUID).
        2. Executar script de atualização para vincular registros existentes ao estágio mais recente do aluno (evitando órfãos).
        3. Aplicar restrição `NOT NULL` e `FOREIGN KEY` referenciando `internship(id)`.

    - Criar o arquivo `.down.sql` correspondente.
- **Fixtures**: Atualizar `config/database/postgres/migrations/fixtures/000002/time_record.csv` adicionando uma coluna para o ID do estágio, garantindo que novos setups de ambiente funcionem.

## 2. Camada de Domínio (TimeRecord)

- **Entidade**: Incluir `InternshipID() uuid.UUID` e `SetInternshipID(uuid.UUID)` na interface e struct.
- **Builder**: Adicionar `WithInternshipID(uuid.UUID)` e validar se o campo foi preenchido no `Build()`.
- **Arquivos**:
    - `src/core/domain/timeRecord/interface.go`
    - `src/core/domain/timeRecord/timeRecord.go`
    - `src/core/domain/timeRecord/builder.go`

## 3. API e DTOs

- **DTO Request**: Atualizar `src/apps/api/handlers/dto/request/timeRecord.go` para incluir `InternshipID uuid.UUID` com a tag `json:"internship_id"`.
- **Swagger**: Garantir que o campo apareça como obrigatório na documentação.

## 4. Camada de Repositório (Postgres)

- **Queries SQL**:
    - `Insert/Update`: Incluir o campo `internship_id`.
    - `Select (All/ByID)`: Remover o `LEFT JOIN LATERAL` (que "adivinhava" o estágio) e substituir por um `INNER JOIN internship ON internship.id = time_record.internship_id`. Isso garante integridade histórica do dado.
- **Mapeamento**: Atualizar `queryObject/timeRecord.go` para popular o domínio corretamente.
- **Arquivos**:
    - `src/infra/repository/postgres/query/timeRecord.go`
    - `src/infra/repository/postgres/queryObject/timeRecord.go`
    - `src/infra/repository/postgres/timeRecord.go`

## 5. Lógica de Serviço (TimeRecord)

- **Validação de Posse**: No `Create`, antes de salvar, buscar o estágio pelo `internship_id` fornecido e validar se `internship.student_id == user_profile_id`.
- **Regras de Negócio**:
    - **Tolerância**: Validar a entrada contra o `schedule_entry_time` do estágio específico vinculado.
    - **Limite de 5h**: Manter o limite diário acumulado (soma de todos os `time_records` do aluno no dia, independente do estágio).
- **Arquivo**: `src/core/services/timeRecord.go`