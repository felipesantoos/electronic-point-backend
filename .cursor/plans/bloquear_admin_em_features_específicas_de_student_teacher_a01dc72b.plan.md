---
name: Bloquear Admin em Features Específicas de Student/Teacher
overview: Analisar e bloquear o admin nas políticas de acesso (policy.csv) e no frontend para operações específicas que requerem student ID ou teacher ID (como criar time records, aprovar records, criar internships), mantendo as verificações nos handlers.
todos:
  - id: analyze-policy
    content: Analisar policy.csv atual e mapear quais rotas devem ser removidas do acesso do admin
    status: pending
  - id: update-policy-csv
    content: Atualizar policy.csv removendo acesso do admin a rotas específicas de student/teacher (POST/PUT/PATCH em time-records, internships, internship-locations)
    status: pending
    dependencies:
      - analyze-policy
  - id: verify-handlers
    content: Verificar que handlers já têm verificações corretas e documentar que devem ser mantidas
    status: pending
  - id: analyze-views
    content: Analisar templates HTML para identificar onde ocultar botões/formulários para admin
    status: pending
  - id: update-templates
    content: Atualizar templates HTML para ocultar criar time records, aprovar/reprovar, criar internships/locations para admin
    status: pending
    dependencies:
      - analyze-views
  - id: review-view-routes
    content: Revisar routes/views.go e remover AdminAuthorize de rotas que admin não deve acessar (como /internships/new, /internship-locations/new)
    status: pending
    dependencies:
      - analyze-policy
  - id: verify-student-create
    content: Verificar se admin pode criar students e se isso é apropriado
    status: pending
---

# Bloquear Admin em Features Específicas de Student/Teacher

## Objetivo

Bloquear o admin de acessar funcionalidades que são específicas de students ou teachers, tanto nas políticas de acesso (policy.csv) quanto no frontend. Isso é necessário porque essas operações requerem o ID do student ou teacher para salvar no banco de dados.

## Análise das Operações por Role

### Operações Específicas de STUDENT

- **TimeRecord.Create**: Requer `studentID` do contexto (handlers/timeRecord.go:55-83)
- **TimeRecord.Update**: Requer `studentID` do contexto (handlers/timeRecord.go:105-139)

### Operações Específicas de TEACHER

- **TimeRecord.Approve**: Requer `teacherID` (approvedBy) do contexto (handlers/timeRecord.go:312-327)
- **TimeRecord.Disapprove**: Requer `teacherID` (disapprovedBy) do contexto (handlers/timeRecord.go:347-362)
- **Internship.Create**: Apenas teacher (handlers/internship.go:52-71)
- **Internship.Update**: Apenas teacher (handlers/internship.go:93-118)
- **InternshipLocation.Create**: Apenas teacher (handlers/internshipLocation.go:52-71)
- **InternshipLocation.Update**: Apenas teacher (handlers/internshipLocation.go:93-122)

### Operações que ADMIN pode realizar

- **TimeRecord.List/Get**: Admin pode ver todos (sem filtros)
- **TimeRecord.Delete**: Verificar se admin deve poder deletar
- **Student.List/Get/Create/Update/Delete**: Admin pode gerenciar students
- **Internship.List/Get/Delete**: Admin pode visualizar/deletar internships
- **InternshipLocation.List/Get/Delete**: Admin pode visualizar/deletar locations
- **Accounts**: Admin tem acesso total via `/admin/accounts/*`

## Plano de Implementação

### Tarefa 1: Analisar Policy CSV Atual

- Revisar `src/apps/api/config/policy.csv`
- Identificar linha 6: `p, admin, /api/*, *` que dá acesso total
- Mapear quais rotas precisam ser removidas do acesso do admin

### Tarefa 2: Atualizar Policy CSV

Remover acesso do admin às seguintes rotas:

- `/api/time-records` (POST) - criar time records
- `/api/time-records/{id}` (PUT) - atualizar time records
- `/api/time-records/{id}/approve` (PATCH) - aprovar time records
- `/api/time-records/{id}/disapprove` (PATCH) - reprovar time records
- `/api/internships` (POST) - criar internships
- `/api/internships/{id}` (PUT) - atualizar internships
- `/api/internship-locations` (POST) - criar internship locations
- `/api/internship-locations/{id}` (PUT) - atualizar internship locations

Manter acesso do admin a:

- `/api/time-records` (GET) - listar time records
- `/api/time-records/{id}` (GET) - ver time record específico
- `/api/time-records/{id}` (DELETE) - deletar time records (se apropriado)
- `/api/internships` (GET) - listar internships
- `/api/internships/{id}` (GET, DELETE) - ver/deletar internship
- `/api/internship-locations` (GET) - listar locations
- `/api/internship-locations/{id}` (GET, DELETE) - ver/deletar location
- `/api/students/*` - todas operações de students
- `/api/admin/*` - todas operações de admin

### Tarefa 3: Verificar Handlers

- Confirmar que handlers já têm verificações corretas:
- TimeRecord.Create/Update: `if ctx.RoleName() != role.STUDENT_ROLE_CODE`
- TimeRecord.Approve/Disapprove: `if ctx.RoleName() != role.TEACHER_ROLE_CODE`
- Internship.Create/Update: `if ctx.RoleName() != role.TEACHER_ROLE_CODE`
- InternshipLocation.Create/Update: `if ctx.RoleName() != role.TEACHER_ROLE_CODE`
- Manter essas verificações como camada adicional de segurança

### Tarefa 4: Analisar Views/Templates do Frontend

- Identificar templates que mostram botões/formulários para:
- Criar time records (`/time-records/new`)
- Aprovar/reprovar time records (botões approve/disapprove)
- Criar internships (`/internships/new`)
- Criar internship locations (`/internship-locations/new`)
- Verificar como verificar role do usuário nos templates (usar `{{if eq .User.RoleName "student"}}` ou similar)

### Tarefa 5: Atualizar Views/Templates

- Ocultar/desabilitar botões de criar time record para admin
- Ocultar botões de aprovar/reprovar para admin
- Ocultar botões de criar internship para admin
- Ocultar botões de criar internship location para admin
- Ocultar links de navegação para páginas de criação se necessário

### Tarefa 6: Verificar Rotas de Views

- Revisar `src/apps/api/routes/views.go`
- Verificar se há rotas protegidas que admin não deve acessar:
- `/time-records/new` - deve ser apenas para students
- `/internships/new` - já tem `AdminAuthorize` (linha 105), remover
- `/internship-locations/new` - já tem `AdminAuthorize` (linha 97), remover
- Adicionar middleware de verificação de role se necessário

### Tarefa 7: Verificar Account Handler (Student.Create)

- Verificar `src/apps/api/handlers/student.go` método `Create`
- Confirmar se admin pode criar students (parece que sim, já que policy permite)
- Se admin pode criar, verificar se precisa de teacher ID (linha 131: `SetResponsibleTeacherID`)

## Arquivos a Modificar

1. **src/apps/api/config/policy.csv** - Remover rotas específicas do acesso do admin
2. **src/apps/api/routes/views.go** - Remover AdminAuthorize de rotas que admin não deve acessar
3. **src/apps/api/views/templates/** - Ocultar botões/formulários para admin nos templates HTML

- `time-records/list.html` - ocultar botão criar
- `time-records/show.html` - ocultar botões aprovar/reprovar
- `internships/list.html` - ocultar botão criar
- `internship-locations/list.html` - ocultar botão criar

## Observações Importantes

- As verificações nos handlers devem ser mantidas como estão (já estão corretas)
- A policy.csv é a primeira camada de defesa
- O frontend deve ocultar features para melhor UX
- Admin ainda pode ver todos os dados (list/get), mas não pode criar/editar operações que requerem student/teacher ID