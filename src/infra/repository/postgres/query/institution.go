package query

const (
	InstitutionID   = "institution_id"
	InstitutionName = "institution_name"
)

type InstitutionQueryBuilder interface {
	Select() InstitutionQuerySelectBuilder
	Insert() string
	Update() string
	Delete() string
}

type institutionQueryBuilder struct{}

func Institution() InstitutionQueryBuilder {
	return &institutionQueryBuilder{}
}

func (*institutionQueryBuilder) Select() InstitutionQuerySelectBuilder {
	return &institutionQuerySelectBuilder{}
}

func (*institutionQueryBuilder) Insert() string {
	return `
		INSERT INTO institution (name)
		VALUES ($1)
		RETURNING id
	`
}

func (*institutionQueryBuilder) Update() string {
	return `
		UPDATE institution
		SET name = $2, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
}

func (*institutionQueryBuilder) Delete() string {
	return `
		UPDATE institution
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
}

type InstitutionQuerySelectBuilder interface {
	All() string
	ByID() string
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

func (institutionQuerySelectBuilder *institutionQuerySelectBuilder) ByID() string {
	return `
		SELECT
			institution.id AS institution_id,
			institution.name AS institution_name
		FROM institution
		WHERE institution.id = $1 AND institution.deleted_at IS NULL
	`
}
