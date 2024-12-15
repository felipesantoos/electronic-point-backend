package query

const (
	InternshipLocationID      = "internship_location_id"
	InternshipLocationName    = "internship_location_name"
	InternshipLocationAddress = "internship_location_address"
	InternshipLocationCity    = "internship_location_city"
	InternshipLocationLat     = "internship_location_lat"
	InternshipLocationLong    = "internship_location_long"
)

type InternshipLocationQueryBuilder interface {
	Select() InternshipLocationQuerySelectBuilder
	Insert() string
	Update() string
	Delete() string
}

type internshipLocationQueryBuilder struct{}

func InternshipLocation() InternshipLocationQueryBuilder {
	return &internshipLocationQueryBuilder{}
}

func (*internshipLocationQueryBuilder) Select() InternshipLocationQuerySelectBuilder {
	return &internshipLocationQuerySelectBuilder{}
}

func (*internshipLocationQueryBuilder) Insert() string {
	return `
		INSERT INTO internship_location (
			name, address, city, lat, long
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING id
	`
}

func (*internshipLocationQueryBuilder) Update() string {
	return `
		UPDATE internship_location
		SET
			name = $2,
			address = $3,
			city = $4,
			lat = $5,
			long = $6,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
}

func (*internshipLocationQueryBuilder) Delete() string {
	return `
		UPDATE internship_location
		SET
			deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
}

type InternshipLocationQuerySelectBuilder interface {
	All() string
	ByID() string
}

type internshipLocationQuerySelectBuilder struct{}

func (internshipLocationQuerySelectBuilder *internshipLocationQuerySelectBuilder) All() string {
	return `
		SELECT DISTINCT
			internship_location.id AS internship_location_id,
			internship_location.name AS internship_location_name,
			internship_location.address AS internship_location_address,
			internship_location.city AS internship_location_city,
			internship_location.lat AS internship_location_lat,
			internship_location.long AS internship_location_long
		FROM internship_location
			LEFT JOIN internship ON internship.internship_location_id = internship_location.id
			LEFT JOIN student ON student.id = internship.student_id
		WHERE internship_location.deleted_at IS NULL
			AND ($1::uuid IS NULL OR student.id = $1)
		ORDER BY internship_location.name ASC
	`
}

func (internshipLocationQuerySelectBuilder *internshipLocationQuerySelectBuilder) ByID() string {
	return `
		SELECT
			internship_location.id AS internship_location_id,
			internship_location.name AS internship_location_name,
			internship_location.address AS internship_location_address,
			internship_location.city AS internship_location_city,
			internship_location.lat AS internship_location_lat,
			internship_location.long AS internship_location_long
		FROM internship_location
		WHERE internship_location.id = $1 AND internship_location.deleted_at IS NULL
	`
}
