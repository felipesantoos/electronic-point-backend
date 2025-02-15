package services

import (
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"
)

type courseServices struct {
	repository secondary.CoursePort
}

func NewCourseServices(repository secondary.CoursePort) primary.CoursePort {
	return &courseServices{repository}
}

func (this *courseServices) List(_filters filters.CourseFilters) ([]course.Course, errors.Error) {
	return this.repository.List(_filters)
}
