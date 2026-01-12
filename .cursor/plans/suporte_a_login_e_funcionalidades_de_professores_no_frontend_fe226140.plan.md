---
name: Suporte a Login e Funcionalidades de Professores no Frontend
overview: "Adicionar suporte completo para professores fazerem login no sistema e acessarem todas as funcionalidades permitidas: gerenciar estudantes, visualizar/aprovar registros de ponto, gerenciar estágios e locais de estágio."
todos:
  - id: update-context-helpers
    content: Adicionar campo IsTeacher ao UserInfo em context.go e atualizar NewPageData para incluir IsTeacher baseado em ctx.RoleName()
    status: completed
  - id: update-login-redirect
    content: Atualizar LoginPage e Login em auth_views.go para redirecionar professores após login (atualmente só redireciona admin)
    status: completed
  - id: update-sidebar
    content: Atualizar sidebar em base.html para ocultar seção Admin e mostrar apenas links permitidos baseado na role (usar isAdmin/isTeacher do renderer)
    status: completed
    dependencies:
      - update-context-helpers
  - id: update-time-record-approve
    content: Atualizar time-records/show.html para mostrar botões de aprovar/desaprovar para professores também (não apenas admin), e remover campo teacher_id para professores
    status: completed
    dependencies:
      - update-context-helpers
  - id: block-time-records-create
    content: Ocultar botão Novo Registro em time-records/list.html para professores (professores não podem criar time records, apenas GET e PATCH)
    status: completed
    dependencies:
      - update-context-helpers
  - id: verify-other-views
    content: Verificar se outras views (students, internships, internship_locations, dashboard) funcionam corretamente para professores - confirmado que já funcionam
    status: completed
---

# Plano: Suporte a Login e Funcionalidades de Professores no Frontend

## Análise das Funcionalidades de Professores

Baseado na análise do backend (policy.csv, handlers e services), professores podem:

### Funcionalidades Completas

1. **Estudantes** (`/api/students/*`)

- Criar estudantes (com seu próprio ID como responsible_teacher_id)
- Atualizar estudantes
- Deletar estudantes
- Listar estudantes (filtrado por TeacherID automaticamente)
- Ver detalhes de estudantes

2. **Registros de Ponto** (`/api/time-records/*`)

- Listar registros (filtrado por TeacherID automaticamente)
- Ver detalhes de registros
- **Aprovar registros** (`PATCH /api/time-records/{id}/approve`)
- **Desaprovar registros** (`PATCH /api/time-records/{id}/disapprove`)

3. **Estágios** (`/api/internships/*`)

- Criar estágios
- Atualizar estágios
- Deletar estágios
- Listar estágios
- Ver detalhes de estágios

4. **Locais de Estágio** (`/api/internship-locations/*`)

- Criar locais de estágio
- Atualizar locais de estágio
- Deletar locais de estágio
- Listar locais de estágio
- Ver detalhes de locais de estágio

### Funcionalidades Somente Leitura

5. **Status de Registro de Ponto** (`/api/time-record-status/*`)

- Listar status
- Ver detalhes de status

6. **Recursos** (apenas GET)

- Cursos (`/api/courses/*`)
- Campus (`/api/campus/*`)
- Instituições (`/api/institutions/*`)
- Arquivos (`/api/files/*`)

### Perfil

7. **Conta** (`/api/accounts/profile`, `/api/accounts/update-password`)

- Ver perfil
- Atualizar perfil
- Atualizar senha

## Problemas Identificados no Frontend

1. **Login**: Redireciona apenas admin após login, não professores
2. **Sidebar**: Mostra seção "Admin" e todos os links para todos os usuários. "Internship Locations" está na seção Admin mas professores podem acessar
3. **Templates**: Botões de aprovar/desaprovar verificam apenas `IsAdmin`, não `IsTeacher`
4. **Helpers de Context**: Não inclui `IsTeacher` no `UserInfo` (apenas `IsAdmin`)
5. **Time Records Create**: Professores NÃO podem criar time records (apenas GET e PATCH), mas a rota `/time-records/new` está acessível para todos
6. **Views**: Algumas views podem não verificar adequadamente permissões de professores

## Implementação

### 1. Atualizar Helpers de Context

**Arquivo**: `src/apps/api/handlers/views/helpers/context.go`

- Adicionar campo `IsTeacher` em `UserInfo`
- Atualizar `NewPageData` para incluir `IsTeacher` baseado em `ctx.RoleName()`

### 2. Atualizar Login

**Arquivo**: `src/apps/api/handlers/views/auth_views.go`

- Modificar `LoginPage` para redirecionar professores também (atualmente só redireciona admin)
- Modificar `Login` se necessário

### 3. Atualizar Sidebar

**Arquivo**: `src/apps/api/views/templates/layouts/base.html`

- Ocultar seção "Admin" para professores (apenas admin pode ver)
- Mover "Internship Locations" para seção "Gestão" (professores podem acessar)
- Ocultar link "Contas" (`/admin/accounts`) para professores
- Mostrar apenas links permitidos baseado na role do usuário
- Usar `.User.IsAdmin` e `.User.IsTeacher` (funções helper `isAdmin`/`isTeacher` já estão registradas)

### 4. Atualizar Templates de Time Records

**Arquivo**: `src/apps/api/views/templates/time-records/show.html`

- Modificar verificação de botões de aprovar/desaprovar para incluir professores
- Botões devem aparecer para `IsAdmin` OU `IsTeacher` quando status for PENDING
- Remover campo de seleção de professor (teacher_id) para professores (só admin precisa)

### 5. Verificar e Atualizar Outras Views

**Arquivos a verificar**:

- `src/apps/api/handlers/views/students_views.go` - Já trata professores parcialmente
- `src/apps/api/handlers/views/time_records_views.go` - Já usa filtros corretos
- `src/apps/api/handlers/views/internships_views.go` - Verificar se funciona para professores
- `src/apps/api/handlers/views/internship_locations_views.go` - Verificar se funciona para professores

### 6. Bloquear Time Records Create para Professores

**Arquivo**: `src/apps/api/routes/views.go` e `src/apps/api/views/templates/layouts/base.html`

- Professores NÃO podem criar time records (apenas GET e PATCH segundo policy.csv)
- Opção 1: Ocultar link "Novo" na listagem de time records para professores
- Opção 2: Adicionar verificação na view CreatePage para redirecionar professores
- Preferir Opção 1 (ocultar link) - mais simples e UX melhor

### 7. Verificar e Atualizar Outras Views (já mencionado na seção 5, mas consolidar aqui)

**Arquivos**:

- `src/apps/api/handlers/views/internships_views.go` - Já funciona para professores (verifica IsAdmin apenas para teacher_id field)
- `src/apps/api/handlers/views/internship_locations_views.go` - Funciona para professores (sem verificações de role)
- `src/apps/api/handlers/views/students_views.go` - Já funciona para professores (usa filtros corretos)
- `src/apps/api/handlers/views/time_records_views.go` - Precisa ocultar link "Novo" para professores
- `src/apps/api/handlers/views/dashboard_views.go` - Já funciona para professores (aplica filtros corretos)

### 8. Nota sobre Renderer

**Arquivo**: `src/apps/api/views/renderer.go`

- Funções `isTeacher` e `isAdmin` JÁ estão registradas no FuncMap (linha 33-34)
- Não é necessário registrar novamente
- Podemos usar `.User.IsTeacher` e `.User.IsAdmin` nos templates, ou `isTeacher .User.RoleName`

## Arquivos a Modificar

1. `src/apps/api/handlers/views/helpers/context.go` - Adicionar campo IsTeacher ao UserInfo
2. `src/apps/api/handlers/views/auth_views.go` - Redirecionar professores após login (LoginPage e Login)
3. `src/apps/api/views/templates/layouts/base.html` - Sidebar condicional (ocultar seção Admin, mover Internship Locations)
4. `src/apps/api/views/templates/time-records/show.html` - Botões de aprovar/desaprovar para professores, remover campo teacher_id para professores
5. `src/apps/api/views/templates/time-records/list.html` - Ocultar botão "Novo" para professores (verificar se existe)

## Arquivos a Verificar (pode não precisar mudanças)

- `src/apps/api/handlers/views/students_views.go` - Já funciona corretamente para professores
- `src/apps/api/handlers/views/time_records_views.go` - Verificar se CreatePage precisa bloquear professores (ou apenas ocultar link)
- `src/apps/api/handlers/views/internships_views.go` - Já funciona corretamente para professores
- `src/apps/api/handlers/views/internship_locations_views.go` - Já funciona corretamente para professores
- `src/apps/api/handlers/views/dashboard_views.go` - Já funciona corretamente para professores
- `src/apps/api/routes/views.go` - Rotas já estão corretas (sem AdminAuthorize onde professores precisam acessar)
- `src/apps/api/views/renderer.go` - Funções helper já estão registradas, não precisa mudanças