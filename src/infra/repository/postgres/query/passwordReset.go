package query

import "fmt"

type PasswordResetQueryBuilder interface {
	Select() PasswordResetQuerySelectBuilder
	Insert() string
	Delete() PasswordResetQueryDeleteBuilder
}

type passwordResetQueryBuilder struct{}

func PasswordReset() PasswordResetQueryBuilder {
	return &passwordResetQueryBuilder{}
}

type PasswordResetQuerySelectBuilder interface {
	All() string
	ByAccountId() string
	ByToken() string
}

type passwordResetQuerySelectBuilder struct{}

type PasswordResetQueryDeleteBuilder interface {
	ByAccountID() string
}

type passwordResetQueryDeleteBuilder struct{}

func (*passwordResetQueryBuilder) Select() PasswordResetQuerySelectBuilder {
	return &passwordResetQuerySelectBuilder{}
}

func (*passwordResetQueryBuilder) Insert() string {
	return `
		INSERT INTO password_reset (account_id, token)
		VALUES ($1, $2)
	`
}

func (*passwordResetQueryBuilder) Delete() PasswordResetQueryDeleteBuilder {
	return &passwordResetQueryDeleteBuilder{}
}

func (q *passwordResetQuerySelectBuilder) All() string {
	return q.defaultStatement("")
}

func (q *passwordResetQuerySelectBuilder) ByAccountId() string {
	return q.defaultStatement("account_id=$1")
}

func (q *passwordResetQuerySelectBuilder) ByToken() string {
	return q.defaultStatement("token=$1")
}

func (*passwordResetQuerySelectBuilder) defaultStatement(whereClause string) string {
	if whereClause != "" {
		whereClause = "WHERE " + whereClause
	}
	return fmt.Sprintf(`
		SELECT account_id FROM password_reset
		%s;
	`, whereClause)
}

func (*passwordResetQueryDeleteBuilder) ByAccountID() string {
	return `
		DELETE FROM password_reset WHERE account_id=$1;
	`
}
