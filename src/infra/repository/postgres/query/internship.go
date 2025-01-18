package query

const (
	InternshipID        = "internship_id"
	InternshipStartedIn = "internship_started_in"
	InternshipEndedIn   = "internship_ended_in"
)

type InternshipQueryBuilder interface {
	Select() InternshipQuerySelectBuilder
	Insert() string
	Update() string
	Delete() string
}

type internshipQueryBuilder struct{}

func Internship() InternshipQueryBuilder {
	return &internshipQueryBuilder{}
}

func (*internshipQueryBuilder) Select() InternshipQuerySelectBuilder {
	return &internshipQuerySelectBuilder{}
}

func (*internshipQueryBuilder) Insert() string {
	return `
		INSERT INTO internship (
			student_id, internship_location_id, started_in, ended_in
		) VALUES (
			$1, $2, $3, $4
		) RETURNING id
	`
}

func (*internshipQueryBuilder) Update() string {
	return `
		UPDATE internship
		SET
			student_id = $2,
			internship_location_id = $3,
			started_in = $4,
			ended_in = $5,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
}

func (*internshipQueryBuilder) Delete() string {
	return `
		UPDATE internship
		SET
			deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
}

type InternshipQuerySelectBuilder interface {
	All() string
	ByID() string
	ByStudentID() string
}

type internshipQuerySelectBuilder struct{}

func (internshipQuerySelectBuilder *internshipQuerySelectBuilder) All() string {
	return `
		SELECT
			internship.id AS internship_id,
			internship.started_in AS internship_started_in,
			internship.ended_in AS internship_ended_in,
			internship_location.id AS internship_location_id,
			internship_location.name AS internship_location_name,
			internship_location.number AS internship_location_number,
			internship_location.street AS internship_location_street,
			internship_location.neighborhood AS internship_location_neighborhood,
			internship_location.city AS internship_location_city,
			internship_location.zip_code AS internship_location_zip_code,
			internship_location.lat AS internship_location_lat,
			internship_location.long AS internship_location_long,
			person.id AS person_id,
			person.name AS person_name,
			student.profile_picture AS student_profile_picture,
			campus.id AS campus_id,
			campus.name AS campus_name,
			institution.id AS institution_id,
			institution.name AS institution_name,
			course.id AS course_id,
			course.name AS course_name
		FROM internship
			INNER JOIN student ON student.id = internship.student_id
			INNER JOIN person ON person.id = student.id
			INNER JOIN campus ON campus.id = student.campus_id
			INNER JOIN institution ON institution.id = campus.institution_id
			INNER JOIN course ON course.id = student.course_id
			INNER JOIN internship_location ON internship_location.id = internship.internship_location_id
		WHERE internship.deleted_at IS NULL
			AND ($1::uuid IS NULL OR internship.student_id = $1)
		ORDER BY person.name ASC, internship.created_at DESC
	`
}

func (internshipQuerySelectBuilder *internshipQuerySelectBuilder) ByID() string {
	return `
		SELECT
			internship.id AS internship_id,
			internship.started_in AS internship_started_in,
			internship.ended_in AS internship_ended_in,
			internship_location.id AS internship_location_id,
			internship_location.name AS internship_location_name,
			internship_location.number AS internship_location_number,
			internship_location.street AS internship_location_street,
			internship_location.neighborhood AS internship_location_neighborhood,
			internship_location.city AS internship_location_city,
			internship_location.zip_code AS internship_location_zip_code,
			internship_location.lat AS internship_location_lat,
			internship_location.long AS internship_location_long,
			person.id AS person_id,
			person.name AS person_name,
			student.profile_picture AS student_profile_picture,
			campus.id AS campus_id,
			campus.name AS campus_name,
			institution.id AS institution_id,
			institution.name AS institution_name,
			course.id AS course_id,
			course.name AS course_name
		FROM internship
			INNER JOIN student ON student.id = internship.student_id
			INNER JOIN person ON person.id = student.id
			INNER JOIN campus ON campus.id = student.campus_id
			INNER JOIN institution ON institution.id = campus.institution_id
			INNER JOIN course ON course.id = student.course_id
			INNER JOIN internship_location ON internship_location.id = internship.internship_location_id
		WHERE internship.id = $1 AND internship.deleted_at IS NULL
	`
}

func (internshipQuerySelectBuilder *internshipQuerySelectBuilder) ByStudentID() string {
	return `
		SELECT
			internship.id AS internship_id,
			internship.started_in AS internship_started_in,
			internship.ended_in AS internship_ended_in,
			internship_location.id AS internship_location_id,
			internship_location.name AS internship_location_name,
			internship_location.number AS internship_location_number,
			internship_location.street AS internship_location_street,
			internship_location.neighborhood AS internship_location_neighborhood,
			internship_location.city AS internship_location_city,
			internship_location.zip_code AS internship_location_zip_code,
			internship_location.lat AS internship_location_lat,
			internship_location.long AS internship_location_long,
			person.id AS person_id,
			person.name AS person_name,
			student.profile_picture AS student_profile_picture,
			campus.id AS campus_id,
			campus.name AS campus_name,
			institution.id AS institution_id,
			institution.name AS institution_name,
			course.id AS course_id,
			course.name AS course_name
		FROM internship
			INNER JOIN student ON student.id = internship.student_id
			INNER JOIN person ON person.id = student.id
			INNER JOIN campus ON campus.id = student.campus_id
			INNER JOIN institution ON institution.id = campus.institution_id
			INNER JOIN course ON course.id = student.course_id
			INNER JOIN internship_location ON internship_location.id = internship.internship_location_id
		WHERE internship.student_id = $1
		ORDER BY internship.created_at DESC
	`
}
