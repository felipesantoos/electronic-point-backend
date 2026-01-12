---
name: Suporte a login e funcionalidades de estudantes no frontend
overview: Implementar suporte completo para o papel de Estudante no frontend, permitindo login, visualização de dashboard filtrada e gerenciamento de seus próprios registros de ponto.
todos:
  - id: update-permissions-entries
    content: Definir studentEntries em permissions.go
    status: completed
  - id: update-auth-redirect
    content: Permitir redirecionamento de estudantes em auth_views.go
    status: completed
  - id: apply-student-filters
    content: Auto-aplicar ProfileID nos filtros para estudantes em filters.go
    status: completed
  - id: filter-dashboard-data-by-student
    content: Filtrar dados do dashboard por StudentID em dashboard_views.go
    status: completed
  - id: optimize-time-records-handler-for-students
    content: Remover carregamento desnecessário de dados em time_records_views.go para estudantes
    status: completed
  - id: update-sidebar-nav-for-students
    content: Atualizar navegação lateral em base.html para estudantes
    status: completed
  - id: adjust-dashboard-metrics-template-for-students
    content: Ajustar métricas do dashboard em dashboard.html para estudantes
    status: completed
  - id: hide-student-info-in-list-template
    content: Esconder filtros e colunas de estudantes em time-records/list.html
    status: completed
  - id: remove-redundant-profile-link-in-show-template
    content: Remover link redundante de perfil em time-records/show.html para estudantes
    status: completed
---

# Suporte a Login e Funcionalidades de Estudantes

Este plano descreve as alterações necessárias para habilitar o acesso de estudantes ao frontend do sistema de ponto eletrônico, garantindo que eles vejam apenas seus próprios dados e tenham uma interface simplificada.

## Alterações no Backend

### 1. Permissões e Segurança

- Atualizar [`src/apps/api/middlewares/permissions/permissions.go`](src/apps/api/middlewares/permissions/permissions.go) para definir `studentEntries`, permitindo o acesso via API para os endpoints definidos no `policy.csv`.
- Ajustar [`src/apps/api/handlers/views/auth_views.go`](src/apps/api/handlers/views/auth_views.go) para que o método `LoginPage` redirecione estudantes para o dashboard caso já estejam logados.

### 2. Filtros e Handlers de Visualização

- Modificar [`src/apps/api/handlers/views/helpers/filters.go`](src/apps/api/handlers/views/helpers/filters.go) para que, se o usuário for um estudante, os filtros de registros de ponto e estudantes sejam automaticamente travados no seu próprio `ProfileID`.
- Atualizar [`src/apps/api/handlers/views/dashboard_views.go`](src/apps/api/handlers/views/dashboard_views.go) para filtrar estatísticas do dashboard pelo ID do estudante.
- Ajustar [`src/apps/api/handlers/views/time_records_views.go`](src/apps/api/handlers/views/time_records_views.go) para simplificar a listagem de pontos quando o usuário é um estudante (removendo lista de seleção de estudantes).

## Alterações no Frontend

### 1. Layout e Navegação

- Modificar [`src/apps/api/views/templates/layouts/base.html`](src/apps/api/views/templates/layouts/base.html) para exibir apenas o menu relevante para estudantes:
    - Dashboard
    - Registros de Ponto
    - Locais de Estágio (leitura)
    - Meus Dados (Perfil)

### 2. Templates de Conteúdo

- **Dashboard**: Ajustar [`src/apps/api/views/templates/dashboard.html`](src/apps/api/views/templates/dashboard.html) para mostrar métricas pertinentes ao estudante.
- **Lista de Pontos**: Ajustar [`src/apps/api/views/templates/time-records/list.html`](src/apps/api/views/templates/time-records/list.html) para esconder filtros de busca de outros estudantes e a coluna de nome do estudante na tabela.
- **Detalhes do Ponto**: Ajustar [`src/apps/api/views/templates/time-records/show.html`](src/apps/api/views/templates/time-records/show.html) para esconder o link de perfil do estudante, já que o próprio estudante estará vendo seu registro.