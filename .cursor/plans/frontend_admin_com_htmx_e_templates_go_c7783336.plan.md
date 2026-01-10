---
name: Frontend Admin com HTMX e Templates Go
overview: Criar um frontend administrativo completo usando HTMX e Go HTML templates padrão para gerenciar todas as funcionalidades da API de ponto eletrônico. O frontend será integrado na mesma aplicação Go usando Echo framework.
todos:
  - id: setup-renderer
    content: Criar renderer de templates Go com funções auxiliares (formatação de datas, UUIDs, etc.)
    status: completed
  - id: create-view-helpers
    content: Criar helpers de views para conversão de domínios, filtros, validação e contexto de usuário
    status: completed
  - id: setup-base-templates
    content: Criar layout base e componentes reutilizáveis (form, table, modal, alerts)
    status: completed
    dependencies:
      - setup-renderer
  - id: implement-auth-views
    content: Implementar páginas e handlers de autenticação (login, logout, reset password)
    status: completed
    dependencies:
      - setup-base-templates
      - create-view-helpers
  - id: implement-dashboard
    content: Criar dashboard principal com visão geral do sistema
    status: completed
    dependencies:
      - setup-base-templates
  - id: implement-accounts-views
    content: Implementar CRUD completo de contas (listagem admin, perfil, atualização)
    status: completed
    dependencies:
      - setup-base-templates
      - create-view-helpers
  - id: implement-students-views
    content: Implementar CRUD completo de estudantes (incluindo upload de foto)
    status: completed
    dependencies:
      - setup-base-templates
      - create-view-helpers
  - id: implement-internships-views
    content: Implementar CRUD completo de estágios
    status: completed
    dependencies:
      - setup-base-templates
      - create-view-helpers
  - id: implement-internship-locations-views
    content: Implementar CRUD completo de locais de estágio
    status: completed
    dependencies:
      - setup-base-templates
      - create-view-helpers
  - id: implement-time-records-views
    content: Implementar CRUD completo de registros de tempo + aprovar/reprovar
    status: completed
    dependencies:
      - setup-base-templates
      - create-view-helpers
  - id: implement-time-record-status-views
    content: Implementar visualização de status de registros de tempo
    status: completed
    dependencies:
      - setup-base-templates
      - create-view-helpers
  - id: implement-readonly-views
    content: Implementar listagens read-only (cursos, campus, instituições)
    status: completed
    dependencies:
      - setup-base-templates
      - create-view-helpers
  - id: setup-routes-views
    content: Criar rotas de frontend e integrar com main.go
    status: completed
    dependencies:
      - implement-auth-views
      - implement-dashboard
  - id: implement-view-auth-middleware
    content: Criar middleware de autenticação específico para views (redirecionar para login)
    status: completed
  - id: add-static-files
    content: Configurar serviço de arquivos estáticos (CSS, JS) e integrar HTMX/Tailwind
    status: completed
    dependencies:
      - setup-base-templates
  - id: add-error-pages
    content: Criar páginas de erro personalizadas (404, 403, 500, 401) e sistema de tratamento de erros da API
    status: completed
    dependencies:
      - setup-base-templates
  - id: implement-permissions-system
    content: Implementar sistema de permissões baseado em roles (helpers de template, menu condicional, filtros por role)
    status: completed
    dependencies:
      - setup-base-templates
      - create-view-helpers
  - id: implement-file-upload
    content: Implementar upload de arquivos com preview (foto de perfil de estudantes) e validação
    status: completed
    dependencies:
      - setup-base-templates
      - implement-students-views
  - id: implement-filters-search
    content: Criar componentes de filtro e busca reutilizáveis com HTMX (debounce, query params, clear)
    status: completed
    dependencies:
      - setup-base-templates
      - implement-students-views
  - id: implement-flash-messages
    content: Implementar sistema de flash messages/toast notifications para feedback de ações
    status: completed
    dependencies:
      - setup-base-templates
  - id: add-htmx-configuration
    content: Configurar HTMX adequadamente (hx-push-url, hx-indicator, hx-headers, tratamento de erros)
    status: completed
    dependencies:
      - setup-base-templates
      - implement-view-auth-middleware
  - id: add-responsive-design
    content: Garantir responsividade mobile (sidebar colapsável, tabelas scrolláveis, modais fullscreen)
    status: completed
    dependencies:
      - setup-base-templates
---

# Implementação de Frontend Admin com HTMX e Templates Go

## Visão Geral

Criar um frontend administrativo completo integrado na aplicação Go existente usando:

- **HTMX** para interatividade sem JavaScript pesado
- **Go HTML templates** padrão para renderização server-side
- **Tailwind CSS** via CDN para estilização rápida
- **Echo framework** já utilizado pela API para servir templates

## Estrutura de Arquivos

```
src/apps/api/
├── views/                     # Templates HTML e renderer
│   ├── renderer.go            # Renderer customizado para templates (Go)
│   ├── templates/             # Templates HTML organizados por módulo
│   │   ├── layouts/
│   │   │   └── base.html          # Layout base com navegação e estrutura comum
│   │   ├── components/
│   │   │   ├── form.html          # Componentes de formulário reutilizáveis
│   │   │   ├── table.html         # Tabelas com paginação
│   │   │   ├── modal.html         # Modais para ações
│   │   │   └── alerts.html        # Mensagens de erro/sucesso
│   │   ├── errors/
│   │   │   ├── 404.html           # Página não encontrada
│   │   │   ├── 403.html           # Acesso negado
│   │   │   ├── 401.html           # Não autenticado
│   │   │   └── 500.html           # Erro interno
│   │   ├── auth/
│   │   │   ├── login.html         # Página de login
│   │   │   ├── reset-password.html # Solicitar reset de senha
│   │   │   └── reset-password-confirm.html # Confirmar reset de senha
│   │   ├── dashboard.html         # Dashboard principal
│   │   ├── accounts/
│   │   │   ├── list.html          # Listagem de contas (admin)
│   │   │   ├── create.html        # Criar conta
│   │   │   └── profile.html       # Perfil do usuário
│   │   ├── students/
│   │   │   ├── list.html          # Listagem de estudantes
│   │   │   ├── create.html        # Criar estudante
│   │   │   ├── edit.html          # Editar estudante
│   │   │   └── show.html          # Visualizar estudante
│   │   ├── internships/
│   │   │   ├── list.html
│   │   │   ├── create.html
│   │   │   ├── edit.html
│   │   │   └── show.html
│   │   ├── internship-locations/
│   │   │   ├── list.html
│   │   │   ├── create.html
│   │   │   ├── edit.html
│   │   │   └── show.html
│   │   ├── time-records/
│   │   │   ├── list.html
│   │   │   ├── create.html
│   │   │   ├── edit.html
│   │   │   └── show.html
│   │   ├── time-record-status/
│   │   │   ├── list.html
│   │   │   └── show.html
│   │   ├── courses/
│   │   │   └── list.html          # Apenas listagem (read-only)
│   │   ├── campus/
│   │   │   └── list.html          # Apenas listagem (read-only)
│   │   └── institutions/
│   │       └── list.html          # Apenas listagem (read-only)
├── static/                    # Arquivos estáticos (CSS, JS, imagens)
│   ├── css/
│   │   └── custom.css         # CSS customizado adicional
│   ├── js/
│   │   └── app.js             # JavaScript mínimo para funcionalidades extras
│   └── images/                # Imagens, ícones, etc.
```

## Implementação

### 1. Renderer de Templates

Criar `src/apps/api/views/renderer.go` para configurar o renderer de templates do Echo com funções auxiliares úteis.

**Estrutura**: `views/` contém tanto o código Go (`renderer.go`) quanto os templates HTML (`templates/`), mantendo tudo relacionado à camada de apresentação junto e dentro da aplicação API (`src/apps/api/`).

**Funções auxiliares para templates**:

- Formatação de datas (`formatDate`, `formatDateTime`)
- Formatação de UUIDs (`formatUUID`)
- Formatação de CPF, telefone, etc.
- Helpers de role/permissões (`isAdmin`, `isTeacher`, `canEdit`)
- Helpers de formatação de valores (`formatCurrency`, se necessário)
- Helpers de URL (`buildURL`, `buildPaginationURL`)

### 2. Cliente HTTP para API

**Decisão de Arquitetura**: Usar os services diretamente através do dicontainer ao invés de fazer chamadas HTTP internas. Isso é mais eficiente e mantém a tipagem forte.

- Criar `src/apps/api/handlers/views/helpers/` com helpers para:
  - Converter dados de domínio para templates
  - Extrair filtros de query parameters
  - Validar parâmetros de formulário
  - Preparar dados de contexto para templates (role, permissions, etc.)

### 3. Handlers de Views

Criar handlers em `src/apps/api/handlers/views/` para cada módulo:

- `auth_views.go` - Login, logout, reset password
- `dashboard_views.go` - Dashboard principal
- `accounts_views.go` - Gerenciamento de contas
- `students_views.go` - CRUD de estudantes
- `internships_views.go` - CRUD de estágios
- `internship_locations_views.go` - CRUD de locais de estágio
- `time_records_views.go` - CRUD de registros de tempo
- `time_record_status_views.go` - Visualização de status
- `courses_views.go` - Listagem de cursos
- `campus_views.go` - Listagem de campus
- `institutions_views.go` - Listagem de instituições

### 4. Rotas de Views

Criar `src/apps/api/routes/views.go` para registrar todas as rotas do frontend, mantendo separação entre rotas da API (`/api/*`) e rotas de views (`/`).

### 5. Middleware de Autenticação para Views

Ajustar middleware existente para redirecionar usuários não autenticados para página de login quando acessarem rotas protegidas via navegador.

### 6. Templates Base e Componentes

Implementar templates reutilizáveis:

- **Layout base** (`layouts/base.html`):
  - Sidebar de navegação com menu baseado em roles (mostrar/ocultar itens baseado em permissões)
  - Header com informações do usuário logado e logout
  - Área de conteúdo principal
  - Sistema de breadcrumbs
  - Suporte a flash messages (mensagens de sucesso/erro temporárias)

- **Componentes de formulário** (`components/form.html`):
  - Campos de input com validação HTML5 e feedback visual
  - Campos de seleção (select) dinâmicos com HTMX para carregar opções
  - Upload de arquivos com preview (foto de perfil)
  - Campos de data/time com formatadores apropriados
  - Mensagens de erro de validação inline

- **Tabelas** (`components/table.html`):
  - Tabelas responsivas com Tailwind
  - Sistema de filtros reutilizável (por nome, data, status, etc.)
  - Busca com debounce via HTMX
  - Indicadores de loading durante requisições
  - Empty states quando não há dados

- **Modais** (`components/modal.html`):
  - Modais de confirmação para ações destrutivas (delete)
  - Modais de formulário para criar/editar
  - Suporte a HTMX para carregar conteúdo dinamicamente

- **Alertas** (`components/alerts.html`):
  - Toast notifications para feedback de ações
  - Flash messages que aparecem após redirects
  - Mensagens de erro da API formatadas de forma amigável

### 7. Integração HTMX

Configurar HTMX com:

- `hx-target` para atualização parcial de seções específicas
- `hx-swap` com estratégias apropriadas (innerHTML, outerHTML, beforeend, etc.)
- `hx-push-url="true"` para manter histórico do navegador correto
- `hx-indicator` para mostrar loading states durante requisições
- `hx-trigger` com eventos customizados (debounce para busca, confirm para delete)
- `hx-headers` para incluir token de autenticação em requisições
- `hx-vals` para passar dados adicionais nas requisições

Casos de uso:

- Atualização parcial de páginas após operações CRUD
- Submissão de formulários sem recarregar página completa
- Carregamento dinâmico de dados em modais
- Busca com debounce (300ms) sem refresh completo
- Validação de formulários em tempo real (usar `hx-trigger="blur"`)
- Carregamento de selects dependentes (ex: campus baseado em instituição)
- Paginação via `hx-get` com query parameters
- Loading indicators durante operações

### 8. Autenticação e Sessão

Implementar:

- Página de login integrada com `/api/auth/login`
- Armazenamento de token JWT em cookie/HTTP-only
- Middleware para verificar autenticação nas rotas de views
- Logout integrado com `/api/auth/logout`

### 9. Tratamento de Erros

- **Páginas de erro personalizadas**:
  - `views/errors/404.html` - Recurso não encontrado
  - `views/errors/403.html` - Acesso negado (com mensagem baseada em role)
  - `views/errors/500.html` - Erro interno do servidor
  - `views/errors/401.html` - Não autenticado (redirecionar para login)

- **Sistema de exibição de erros da API**:
  - Converter códigos de erro HTTP em mensagens amigáveis em português
  - Exibir erros de validação (422) junto aos campos correspondentes
  - Toast notifications para erros de operações (criar, atualizar, deletar)
  - Tratamento de erros de rede/timeout com mensagens apropriadas
  - Modal de erro para operações críticas

- **Tratamento de erros HTMX**:
  - Usar `hx-on::htmx:response-error` para capturar erros
  - Redirecionar para login em caso de 401
  - Mostrar mensagens de erro formatadas em caso de 4xx/5xx

## Endpoints Mapeados

### Autenticação

- `GET /login` - Página de login
- `POST /login` - Processar login (via API)
- `POST /logout` - Logout
- `GET /reset-password` - Solicitar reset
- `GET /reset-password/:token` - Confirmar token
- `PUT /reset-password/:token` - Atualizar senha

### Dashboard

- `GET /` - Dashboard principal

### Contas (Admin)

- `GET /admin/accounts` - Listar contas
- `GET /admin/accounts/new` - Formulário criar conta
- `POST /admin/accounts` - Criar conta (via API)
- `GET /accounts/profile` - Perfil do usuário
- `PUT /accounts/profile` - Atualizar perfil (via API)

### Estudantes

- `GET /students` - Listar estudantes
- `GET /students/new` - Formulário criar
- `POST /students` - Criar (via API)
- `GET /students/:id` - Visualizar
- `GET /students/:id/edit` - Formulário editar
- `PUT /students/:id` - Atualizar (via API)
- `DELETE /students/:id` - Deletar (via API)

### Estágios

- `GET /internships` - Listar
- `GET /internships/new` - Criar
- `GET /internships/:id` - Visualizar
- `GET /internships/:id/edit` - Editar
- Operações CRUD via API

### Locais de Estágio

- CRUD completo similar a estágios

### Registros de Tempo

- CRUD completo + ações de aprovar/reprovar

### Status de Registro de Tempo

- Apenas visualização (listagem e detalhes)

### Cursos, Campus, Instituições

- Apenas listagem (read-only)

## Arquivos Principais a Criar/Modificar

**Estrutura de Pastas** (tudo dentro de `src/apps/api/`):

- `src/apps/api/views/renderer.go` - Renderer de templates com funções auxiliares
- `src/apps/api/views/templates/` - Todos os templates HTML organizados por módulo
- `src/apps/api/handlers/views/helpers/` - Helpers Go para conversão, filtros, validação
- `src/apps/api/handlers/views/*.go` - Handlers de views (usam services via dicontainer)
- `src/apps/api/routes/views.go` - Rotas de frontend
- `src/apps/api/middlewares/views_auth.go` - Middleware de autenticação para views (redireciona para login)
- `src/apps/api/main.go` - Modificar para incluir renderer e servir arquivos estáticos (`/static/*`)
- `src/apps/api/static/css/custom.css` - CSS customizado adicional
- `src/apps/api/static/js/app.js` - JavaScript mínimo para preview de imagens, etc.

**Justificativa da Estrutura**:

- `views/` fica dentro de `src/apps/api/` porque faz parte da aplicação API
- Renderer e templates ficam juntos para facilitar manutenção
- Separação clara entre código Go (`renderer.go`) e templates HTML (`templates/`)
- Consistente com a organização atual: `handlers/`, `routes/`, `middlewares/` também estão dentro de `api/`

## Considerações Técnicas Adicionais

### Autenticação e Sessão

- Usar contexto do Echo (`RichContext`) para passar dados do usuário autenticado para templates
- Middleware específico para views que redireciona para `/login` se não autenticado
- Armazenar token JWT em cookie HTTP-only (já implementado na API)
- Passar token para requisições HTMX via header `Authorization: Bearer <token>`
- Verificar role do usuário nos templates para mostrar/ocultar funcionalidades

### Permissões Baseadas em Roles

- Criar helpers de template para verificar roles: `{{if isAdmin .User}}`, `{{if canEdit .User}}`
- Ocultar botões/links baseado em permissões (ex: apenas admin pode criar contas)
- Filtrar dados automaticamente baseado em role (ex: professores só veem seus estudantes)
- Mostrar diferentes menus na sidebar baseado em role

### Upload de Arquivos

- Formulários multipart/form-data para upload de fotos de perfil
- Preview de imagem antes de enviar usando JavaScript
- Progress indicator durante upload
- Validação de tipo e tamanho de arquivo no frontend
- Exibição de imagem atual ao editar (se existir)

### Filtros e Busca

- Componentes de filtro reutilizáveis para cada entidade
- Filtros comuns: nome, data (start/end), status, instituição, campus
- Busca com debounce (300ms) usando HTMX `hx-trigger="keyup changed delay:300ms"`
- Manter estado dos filtros na URL como query parameters
- Botão de limpar filtros

### Paginação

- **Nota**: API atual não retorna metadados de paginação. Opções:

  1. Implementar paginação no backend primeiro (retornar total, page, limit)
  2. Carregar todos os resultados e paginar no frontend (limitado para listas pequenas)
  3. Implementar "load more" infinito scroll com HTMX

- Se implementar paginação backend: componentes reutilizáveis de paginação

### Validação

- Validação HTML5 básica (required, type, pattern)
- Validação em tempo real com HTMX após blur dos campos
- Validação de formato (email, CPF, telefone, datas)
- Mensagens de erro em português próximas aos campos
- Desabilitar botão de submit durante validação/envio

### Performance

- Lazy loading de imagens (fotos de perfil, avatares)
- Cache de dados estáticos (roles, instituições, campus) no contexto da sessão
- Minimizar requisições desnecessárias usando `hx-trigger` apropriado
- Loading states para evitar cliques múltiplos

### Responsividade

- Layout responsivo com Tailwind (mobile-first)
- Sidebar colapsável em mobile
- Tabelas scrolláveis horizontalmente em mobile
- Modais fullscreen em mobile
- Formulários em uma coluna em mobile

### Acessibilidade

- Labels apropriados em todos os formulários
- Contraste adequado de cores
- Navegação por teclado funcional
- ARIA labels onde necessário
- Foco visível em elementos interativos

### Flash Messages

- Sistema de flash messages usando session/cookies
- Tipos: success, error, warning, info
- Auto-dismiss após 5 segundos
- Posicionamento não intrusivo (top-right)

### Integração com API

- **Decisão**: Usar services diretamente via dicontainer ao invés de HTTP interno
- Vantagens: tipagem forte, sem overhead de HTTP, melhor performance
- Handlers de views chamam services diretamente
- Converter domínios para DTOs de template (ViewModels)