package query

const (
	CampusID            = "campus_id"
	CampusName          = "campus_name"
	CampusInstitutionID = "campus_institution_id"
)

type CampusQueryBuilder interface {
	Select() CampusQuerySelectBuilder
	Insert() string
	Update() string
	Delete() string
}

type campusQueryBuilder struct{}

func Campus() CampusQueryBuilder {
	return &campusQueryBuilder{}
}

func (*campusQueryBuilder) Select() CampusQuerySelectBuilder {
	return &campusQuerySelectBuilder{}
}

func (*campusQueryBuilder) Insert() string {
	return `
		INSERT INTO campus (name, institution_id)
		VALUES ($1, $2)
		RETURNING id
	`
}

func (*campusQueryBuilder) Update() string {
	return `
		UPDATE campus
		SET name = $2, institution_id = $3, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
}

func (*campusQueryBuilder) Delete() string {
	return `
		UPDATE campus
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
}

type CampusQuerySelectBuilder interface {
	All() string
	ByID() string
}

type campusQuerySelectBuilder struct{}

func (campusQuerySelectBuilder *campusQuerySelectBuilder) All() string {
	return `
		SELECT
			campus.id AS campus_id,
			campus.name AS campus_name,
			campus.institution_id AS campus_institution_id
		FROM campus
		WHERE campus.deleted_at IS NULL
			AND campus.name LIKE '%' || COALESCE($1, campus.name) || '%'
		ORDER BY campus.name ASC
	`
}

func (campusQuerySelectBuilder *campusQuerySelectBuilder) ByID() string {
	return `
		SELECT
			campus.id AS campus_id,
			campus.name AS campus_name,
			campus.institution_id AS campus_institution_id
		FROM campus
		WHERE campus.id = $1 AND campus.deleted_at IS NULL
	`
}
