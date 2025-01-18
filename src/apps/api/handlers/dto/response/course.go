package response

import (
	"eletronic_point/src/core/domain/course"

	"github.com/google/uuid"
)

type Course struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type courseBuilder struct{}

func CourseBuilder() *courseBuilder {
	return &courseBuilder{}
}

func (*courseBuilder) BuildFromDomain(data course.Course) Course {
	return Course{
		ID:   data.ID(),
		Name: data.Name(),
	}
}

func (*courseBuilder) BuildFromDomainList(data []course.Course) []Course {
	courses := make([]Course, 0)
	for _, _course := range data {
		courses = append(courses, CourseBuilder().BuildFromDomain(_course))
	}
	return courses
}
