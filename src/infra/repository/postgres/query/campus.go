package query

const (
	CampusID   = "campus_id"
	CampusName = "campus_name"
)

type CampusQueryBuilder interface {
	Select() CampusQuerySelectBuilder
}

type campusQueryBuilder struct{}

func Campus() CampusQueryBuilder {
	return &campusQueryBuilder{}
}

func (*campusQueryBuilder) Select() CampusQuerySelectBuilder {
	return &campusQuerySelectBuilder{}
}

type CampusQuerySelectBuilder interface {
	All() string
}

type campusQuerySelectBuilder struct{}

func (campusQuerySelectBuilder *campusQuerySelectBuilder) All() string {
	return `
		SELECT
			campus.id AS campus_id,
			campus.name AS campus_name
		FROM campus
		WHERE campus.deleted_at IS NULL
			AND campus.name LIKE '%' || COALESCE($1, campus.name) || '%'
		ORDER BY campus.name ASC
	`
}
