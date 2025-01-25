package query

const (
	StudentLinkedToTeacherID        = "student_linked_to_teacher_id"
	StudentLinkedToTeacherStudentID = "student_linked_to_teacher_student_id"
	StudentLinkedToTeacherTeacherID = "student_linked_to_teacher_teacher_id"
	StudentLinkedToTeacherCreatedAt = "student_linked_to_teacher_created_at"
	StudentLinkedToTeacherUpdatedAt = "student_linked_to_teacher_updated_at"
	StudentLinkedToTeacherDeletedAt = "student_linked_to_teacher_deleted_at"
)

type StudentLinkedToTeacherQueryBuilder interface {
	Select() StudentLinkedToTeacherQuerySelectBuilder
	Insert() string
	Update() string
	Delete() string
}

type studentLinkedToTeacherQueryBuilder struct{}

func StudentLinkedToTeacher() StudentLinkedToTeacherQueryBuilder {
	return &studentLinkedToTeacherQueryBuilder{}
}

func (*studentLinkedToTeacherQueryBuilder) Select() StudentLinkedToTeacherQuerySelectBuilder {
	return &studentLinkedToTeacherQuerySelectBuilder{}
}

func (*studentLinkedToTeacherQueryBuilder) Insert() string {
	return `
		INSERT INTO student_linked_to_teacher (
			student_id, teacher_id
		) VALUES (
			$1, $2
		) RETURNING id
	`
}

func (*studentLinkedToTeacherQueryBuilder) Update() string {
	return `
		UPDATE student_linked_to_teacher
		SET
			student_id = $2,
			teacher_id = $3,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
}

func (*studentLinkedToTeacherQueryBuilder) Delete() string {
	return `
		UPDATE student_linked_to_teacher
		SET
			deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
}

type StudentLinkedToTeacherQuerySelectBuilder interface {
	All() string
	ByID() string
}

type studentLinkedToTeacherQuerySelectBuilder struct{}

func (*studentLinkedToTeacherQuerySelectBuilder) All() string {
	return `
		SELECT
			student_linked_to_teacher.id AS student_linked_to_teacher_id,
			student_linked_to_teacher.student_id AS student_linked_to_teacher_student_id,
			student_linked_to_teacher.teacher_id AS student_linked_to_teacher_teacher_id,
			student_linked_to_teacher.created_at AS student_linked_to_teacher_created_at,
			student_linked_to_teacher.updated_at AS student_linked_to_teacher_updated_at,
			student_linked_to_teacher.deleted_at AS student_linked_to_teacher_deleted_at
		FROM student_linked_to_teacher
		WHERE deleted_at IS NULL
			AND student_id = COALESCE($1, student_id)
			AND teacher_id = COALESCE($2, teacher_id)
		ORDER BY created_at ASC
	`
}

func (*studentLinkedToTeacherQuerySelectBuilder) ByID() string {
	return `
		SELECT
			student_linked_to_teacher.id AS student_linked_to_teacher_id,
			student_linked_to_teacher.student_id AS student_linked_to_teacher_student_id,
			student_linked_to_teacher.teacher_id AS student_linked_to_teacher_teacher_id,
			student_linked_to_teacher.created_at AS student_linked_to_teacher_created_at,
			student_linked_to_teacher.updated_at AS student_linked_to_teacher_updated_at,
			student_linked_to_teacher.deleted_at AS student_linked_to_teacher_deleted_at
		FROM student_linked_to_teacher
		WHERE id = $1 AND deleted_at IS NULL
	`
}
