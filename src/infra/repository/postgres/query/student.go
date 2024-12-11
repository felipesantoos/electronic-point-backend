package query

const (
	StudentID                     = "student_id"
	StudentName                   = "student_name"
	StudentRegistration           = "student_registration"
	StudentProfilePicture         = "student_profile_picture"
	StudentInstitution            = "student_institution"
	StudentCourse                 = "student_course"
	StudentInternshipLocationName = "student_internship_location_name"
	StudentInternshipAddress      = "student_internship_address"
	StudentInternshipLocation     = "student_internship_location"
	StudentTotalWorkload          = "student_total_workload"
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
			name, registration, profile_picture, institution, course, internship_location_name, 
			internship_address, internship_location, total_workload
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		) RETURNING id
	`
}

func (*studentQueryBuilder) Update() string {
	return `
		UPDATE student
		SET
			name = $2,
			registration = $3,
			profile_picture = $4,
			institution = $5,
			course = $6,
			internship_location_name = $7,
			internship_address = $8,
			internship_location = $9,
			total_workload = $10,
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
			student.id AS student_id,
			student.name AS student_name,
			student.registration AS student_registration,
			student.profile_picture AS student_profile_picture,
			student.institution AS student_institution,
			student.course AS student_course,
			student.internship_location_name AS student_internship_location_name,
			student.internship_address AS student_internship_address,
			student.internship_location AS student_internship_location,
			student.total_workload AS student_total_workload
		FROM student
		WHERE student.deleted_at IS NULL
		ORDER BY student.name ASC
	`
}

func (studentQuerySelectBuilder *studentQuerySelectBuilder) ByID() string {
	return `
		SELECT
			student.id AS student_id,
			student.name AS student_name,
			student.registration AS student_registration,
			student.profile_picture AS student_profile_picture,
			student.institution AS student_institution,
			student.course AS student_course,
			student.internship_location_name AS student_internship_location_name,
			student.internship_address AS student_internship_address,
			student.internship_location AS student_internship_location,
			student.total_workload AS student_total_workload
		FROM student
		WHERE student.id = $1 AND student.deleted_at IS NULL
	`
}
