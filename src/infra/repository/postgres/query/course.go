package query

const (
	CourseID   = "course_id"
	CourseName = "course_name"
)

type CourseQueryBuilder interface {
	Select() CourseQuerySelectBuilder
	Insert() string
	Update() string
	Delete() string
}

type courseQueryBuilder struct{}

func Course() CourseQueryBuilder {
	return &courseQueryBuilder{}
}

func (*courseQueryBuilder) Select() CourseQuerySelectBuilder {
	return &courseQuerySelectBuilder{}
}

func (*courseQueryBuilder) Insert() string {
	return `
		INSERT INTO course (name)
		VALUES ($1)
		RETURNING id
	`
}

func (*courseQueryBuilder) Update() string {
	return `
		UPDATE course
		SET name = $2, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
}

func (*courseQueryBuilder) Delete() string {
	return `
		UPDATE course
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
}

type CourseQuerySelectBuilder interface {
	All() string
	ByID() string
}

type courseQuerySelectBuilder struct{}

func (courseQuerySelectBuilder *courseQuerySelectBuilder) All() string {
	return `
		SELECT
			course.id AS course_id,
			course.name AS course_name
		FROM course
		WHERE course.deleted_at IS NULL
			AND course.name LIKE '%' || COALESCE($1, course.name) || '%'
		ORDER BY course.name ASC
	`
}

func (courseQuerySelectBuilder *courseQuerySelectBuilder) ByID() string {
	return `
		SELECT
			course.id AS course_id,
			course.name AS course_name
		FROM course
		WHERE course.id = $1 AND course.deleted_at IS NULL
	`
}
