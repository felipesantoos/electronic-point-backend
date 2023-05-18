package query

import "fmt"

type AccountQueryBuilder interface {
	Select() AccountQuerySelectBuilder
	Insert() string
	Update() AccountQueryUpdateBuilder
}

type accountQueryBuilder struct{}

func Account() AccountQueryBuilder {
	return &accountQueryBuilder{}
}

type AccountQuerySelectBuilder interface {
	All() string
	ByID() string
	PasswordByID() string
	ByCredentials() string
	SimplifiedByID() string
	SimplifiedByEmail() string
}

type accountQuerySelectBuilder struct{}

type AccountQueryUpdateBuilder interface {
	Password() string
	Profile() string
}

type accountQueryUpdateBuilder struct{}

func (*accountQueryBuilder) Select() AccountQuerySelectBuilder {
	return &accountQuerySelectBuilder{}
}

func (*accountQueryBuilder) Insert() string {
	return `
		INSERT INTO account (email, password, person_id, role_id)
		SELECT $1, $2, $3, ar.id FROM account_role ar WHERE lower(ar.code)=lower($4)
		RETURNING id;
	`
}

func (*accountQueryBuilder) Update() AccountQueryUpdateBuilder {
	return &accountQueryUpdateBuilder{}
}

func (q *accountQuerySelectBuilder) All() string {
	return q.defaultStatement("")
}

func (q *accountQuerySelectBuilder) ByID() string {
	return q.defaultStatement("a.id=$1")
}

func (q *accountQuerySelectBuilder) PasswordByID() string {
	return `
		SELECT a.password FROM account a WHERE id=$1;
	`
}

func (q *accountQuerySelectBuilder) SimplifiedByID() string {
	return q.defaultSimplifiedStatement("a.id=$1")
}

func (q *accountQuerySelectBuilder) SimplifiedByEmail() string {
	return q.defaultSimplifiedStatement("a.email=$1")
}

func (q *accountQuerySelectBuilder) ByCredentials() string {
	return `SELECT id, password FROM account WHERE email=$1;`
}

func (*accountQuerySelectBuilder) defaultStatement(whereClause string) string {
	if whereClause != "" {
		whereClause = fmt.Sprintf("WHERE %s", whereClause)
	}
	return fmt.Sprintf(`
		SELECT
			a.id AS id,
			a.email AS email,
			a.person_id AS person_id,
			p.name AS person_name,
			p.birth_date AS person_birth_date,
			p.phone AS person_phone,
			p.cpf AS person_cpf,
			p.created_at AS person_created_at,
			p.updated_at AS person_updated_at,
			a.password AS password,
			ar.id AS role_id,
			ar.name AS role_name,
			ar.code AS role_code,
			prof.id AS professional_id,
			a.created_at AS created_at,
			a.updated_at AS updated_at
		FROM
			account a
		INNER JOIN account_role ar
			ON ar.id = a.role_id
		INNER JOIN person p
			ON p.id = a.person_id

		LEFT JOIN professional prof
			ON prof.person_id = p.id
		%s
	`, whereClause)
}

func (*accountQuerySelectBuilder) defaultSimplifiedStatement(whereClause string) string {
	if whereClause != "" {
		whereClause = fmt.Sprintf("WHERE %s", whereClause)
	}
	return fmt.Sprintf(`
		SELECT
			a.id AS account_id,
			a.email AS account_email,
			p.name AS person_name,
			p.birth_date AS person_birth_date,
			p.cpf AS person_cpf
		FROM
			account a
		INNER JOIN person p
			ON p.id = a.person_id
		%s
	`, whereClause)
}

func (*accountQueryUpdateBuilder) Password() string {
	return `
		UPDATE account SET password=$1 WHERE id=$2;
	`
}

func (*accountQueryUpdateBuilder) Profile() string {
	return `
		UPDATE
			person
		SET
			name=$1,
			birth_date=$2,
			phone=$3
		WHERE id=$4
	`
}
