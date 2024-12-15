package query

const (
	StudentWorksAtInternshipLocationID        = "student_works_at_internship_location_id"
	StudentWorksAtInternshipLocationStartedIn = "student_works_at_internship_location_started_in"
	StudentWorksAtInternshipLocationEndedIn   = "student_works_at_internship_location_ended_in"
)

type StudentWorksAtInternshipLocationQueryBuilder interface {
	Select() StudentWorksAtInternshipLocationQuerySelectBuilder
}

type studentWorksAtInternshipLocationQueryBuilder struct{}

func StudentWorksAtInternshipLocation() StudentWorksAtInternshipLocationQueryBuilder {
	return &studentWorksAtInternshipLocationQueryBuilder{}
}

func (*studentWorksAtInternshipLocationQueryBuilder) Select() StudentWorksAtInternshipLocationQuerySelectBuilder {
	return &studentWorksAtInternshipLocationQuerySelectBuilder{}
}

type StudentWorksAtInternshipLocationQuerySelectBuilder interface {
	ByStudentID() string
}

type studentWorksAtInternshipLocationQuerySelectBuilder struct{}

func (studentWorksAtInternshipLocationQuerySelectBuilder *studentWorksAtInternshipLocationQuerySelectBuilder) ByStudentID() string {
	return `
		SELECT
			student_works_at_internship_location.id AS student_works_at_internship_location_id,
			student_works_at_internship_location.started_in AS student_works_at_internship_location_started_in,
			student_works_at_internship_location.ended_in AS student_works_at_internship_location_ended_in,
			internship_location.id AS internship_location_id,
			internship_location.name AS internship_location_name,
			internship_location.address AS internship_location_address,
			internship_location.city AS internship_location_city,
			internship_location.lat AS internship_location_lat,
			internship_location.long AS internship_location_long
		FROM student_works_at_internship_location
			INNER JOIN internship_location ON internship_location.id = student_works_at_internship_location.internship_location_id
		WHERE student_works_at_internship_location.student_id = $1
		ORDER BY student_works_at_internship_location.created_at DESC
	`
}
