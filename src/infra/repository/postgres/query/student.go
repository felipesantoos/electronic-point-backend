package query

const (
	StudentID             = "student_id"
	StudentRegistration   = "student_registration"
	StudentProfilePicture = "student_profile_picture"
	StudentInstitution    = "student_institution"
	StudentCourse         = "student_course"
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
			id, registration, profile_picture, institution, course, total_workload
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
			institution = $4,
			course = $5,
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
			student.institution AS student_institution,
			student.course AS student_course,
			student.total_workload AS student_total_workload,
			internship_location.id AS internship_location_id,
			internship_location.name AS internship_location_name,
			internship_location.address AS internship_location_address,
			internship_location.city AS internship_location_city,
			internship_location.lat AS internship_location_lat,
			internship_location.long AS internship_location_long,
			internship.id AS internship_id,
			internship.started_in AS internship_started_in,
			internship.ended_in AS internship_ended_in
		FROM student
			INNER JOIN person ON person.id = student.id
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
			student.institution AS student_institution,
			student.course AS student_course,
			student.total_workload AS student_total_workload,
			internship_location.id AS internship_location_id,
			internship_location.name AS internship_location_name,
			internship_location.address AS internship_location_address,
			internship_location.city AS internship_location_city,
			internship_location.lat AS internship_location_lat,
			internship_location.long AS internship_location_long,
			internship.id AS internship_id,
			internship.started_in AS internship_started_in,
			internship.ended_in AS internship_ended_in
		FROM student
			INNER JOIN person ON person.id = student.id
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
