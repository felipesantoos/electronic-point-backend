package query

const (
	InstitutionID   = "institution_id"
	InstitutionName = "institution_name"
)

type InstitutionQueryBuilder interface {
	Select() InstitutionQuerySelectBuilder
}

type institutionQueryBuilder struct{}

func Institution() InstitutionQueryBuilder {
	return &institutionQueryBuilder{}
}

func (*institutionQueryBuilder) Select() InstitutionQuerySelectBuilder {
	return &institutionQuerySelectBuilder{}
}

type InstitutionQuerySelectBuilder interface {
	All() string
}

type institutionQuerySelectBuilder struct{}

func (institutionQuerySelectBuilder *institutionQuerySelectBuilder) All() string {
	return `
		SELECT
			institution.id AS institution_id,
			institution.name AS institution_name
		FROM institution
		WHERE institution.deleted_at IS NULL
			AND institution.name LIKE '%' || COALESCE($1, institution.name) || '%'
		ORDER BY institution.name ASC
	`
}
