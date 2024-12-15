package query

const (
	InternshipID        = "internship_id"
	InternshipStartedIn = "internship_started_in"
	InternshipEndedIn   = "internship_ended_in"
)

type InternshipQueryBuilder interface {
	Select() InternshipQuerySelectBuilder
}

type internshipQueryBuilder struct{}

func Internship() InternshipQueryBuilder {
	return &internshipQueryBuilder{}
}

func (*internshipQueryBuilder) Select() InternshipQuerySelectBuilder {
	return &internshipQuerySelectBuilder{}
}

type InternshipQuerySelectBuilder interface {
	ByStudentID() string
}

type internshipQuerySelectBuilder struct{}

func (internshipQuerySelectBuilder *internshipQuerySelectBuilder) ByStudentID() string {
	return `
		SELECT
			internship.id AS internship_id,
			internship.started_in AS internship_started_in,
			internship.ended_in AS internship_ended_in,
			internship_location.id AS internship_location_id,
			internship_location.name AS internship_location_name,
			internship_location.address AS internship_location_address,
			internship_location.city AS internship_location_city,
			internship_location.lat AS internship_location_lat,
			internship_location.long AS internship_location_long
		FROM internship
			INNER JOIN internship_location ON internship_location.id = internship.internship_location_id
		WHERE internship.student_id = $1
		ORDER BY internship.created_at DESC
	`
}
