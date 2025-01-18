package query

const (
	StudentID             = "student_id"
	StudentRegistration   = "student_registration"
	StudentProfilePicture = "student_profile_picture"
	StudentTotalWorkload  = "student_total_workload"
)

type StudentQueryBuilder interface {
	Select() StudentQuerySelectBuilder
	Insert() string
	Update() string
	Delete() string
}

type studentQueryBuilder struct{}

func Student() StudentQueryBuilder {
	return &studentQueryBuilder{}
}

func (*studentQueryBuilder) Select() StudentQuerySelectBuilder {
	return &studentQuerySelectBuilder{}
}

func (*studentQueryBuilder) Insert() string {
	return `
		INSERT INTO student (
			id, registration, profile_picture, campus_id, course_id, total_workload
		) VALUES (
			$1, $2, $3, $4, $5, $6
		) RETURNING id
	`
}

func (*studentQueryBuilder) Update() string {
	return `
		UPDATE student
		SET
			registration = $2,
			profile_picture = $3,
			campus_id = $4,
			course_id = $5,
			total_workload = $6,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
}

func (*studentQueryBuilder) Delete() string {
	return `
		UPDATE student
		SET
			deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
}

type StudentQuerySelectBuilder interface {
	All() string
	ByID() string
}

type studentQuerySelectBuilder struct{}

func (studentQuerySelectBuilder *studentQuerySelectBuilder) All() string {
	return `
		SELECT
			person.id AS person_id,
			person.name AS person_name,
			person.birth_date AS person_birth_date,
			person.email AS person_email,
			person.cpf AS person_cpf,
			person.phone AS person_phone,
			student.registration AS student_registration,
			student.profile_picture AS student_profile_picture,
			campus.id AS campus_id,
			campus.name AS campus_name,
			institution.id AS institution_id,
			institution.name AS institution_name,
			course.id AS course_id,
			course.name AS course_name,
			student.total_workload AS student_total_workload,
			internship_location.id AS internship_location_id,
			internship_location.name AS internship_location_name,
			internship_location.number AS internship_location_number,
			internship_location.street AS internship_location_street,
			internship_location.neighborhood AS internship_location_neighborhood,
			internship_location.city AS internship_location_city,
			internship_location.zip_code AS internship_location_zip_code,
			internship_location.lat AS internship_location_lat,
			internship_location.long AS internship_location_long,
			internship.id AS internship_id,
			internship.started_in AS internship_started_in,
			internship.ended_in AS internship_ended_in
		FROM student
			INNER JOIN person ON person.id = student.id
			INNER JOIN campus ON campus.id = student.campus_id
			INNER JOIN institution ON institution.id = campus.institution_id
			INNER JOIN course ON course.id = student.course_id
			LEFT JOIN LATERAL (
				SELECT 
					internship.id, 
					internship.internship_location_id, 
					internship.started_in, 
					internship.ended_in
				FROM internship
				WHERE internship.student_id = student.id
				ORDER BY internship.created_at DESC
				LIMIT 1
			) internship ON true
			LEFT JOIN internship_location ON internship_location.id = internship.internship_location_id
		WHERE student.deleted_at IS NULL
		ORDER BY person.name ASC
	`
}

func (studentQuerySelectBuilder *studentQuerySelectBuilder) ByID() string {
	return `
		SELECT
			person.id AS person_id,
			person.name AS person_name,
			person.birth_date AS person_birth_date,
			person.email AS person_email,
			person.cpf AS person_cpf,
			person.phone AS person_phone,
			student.registration AS student_registration,
			student.profile_picture AS student_profile_picture,
			campus.id AS campus_id,
			campus.name AS campus_name,
			institution.id AS institution_id,
			institution.name AS institution_name,
			course.id AS course_id,
			course.name AS course_name,
			student.total_workload AS student_total_workload,
			internship_location.id AS internship_location_id,
			internship_location.name AS internship_location_name,
			internship_location.number AS internship_location_number,
			internship_location.street AS internship_location_street,
			internship_location.neighborhood AS internship_location_neighborhood,
			internship_location.city AS internship_location_city,
			internship_location.zip_code AS internship_location_zip_code,
			internship_location.lat AS internship_location_lat,
			internship_location.long AS internship_location_long,
			internship.id AS internship_id,
			internship.started_in AS internship_started_in,
			internship.ended_in AS internship_ended_in
		FROM student
			INNER JOIN person ON person.id = student.id
			INNER JOIN campus ON campus.id = student.campus_id
			INNER JOIN institution ON institution.id = campus.institution_id
			INNER JOIN course ON course.id = student.course_id
			LEFT JOIN LATERAL (
				SELECT 
					internship.id, 
					internship.internship_location_id, 
					internship.started_in, 
					internship.ended_in
				FROM internship
				WHERE internship.student_id = student.id
				ORDER BY internship.created_at DESC
				LIMIT 1
			) internship ON true
			LEFT JOIN internship_location ON internship_location.id = internship.internship_location_id
		WHERE person.id = $1 AND student.deleted_at IS NULL
	`
}
