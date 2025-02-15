package query

const (
	CourseID   = "course_id"
	CourseName = "course_name"
)

type CourseQueryBuilder interface {
	Select() CourseQuerySelectBuilder
}

type courseQueryBuilder struct{}

func Course() CourseQueryBuilder {
	return &courseQueryBuilder{}
}

func (*courseQueryBuilder) Select() CourseQuerySelectBuilder {
	return &courseQuerySelectBuilder{}
}

type CourseQuerySelectBuilder interface {
	All() string
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
