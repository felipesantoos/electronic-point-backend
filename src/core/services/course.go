package services

import (
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"
	"github.com/google/uuid"
)

type courseServices struct {
	repository secondary.CoursePort
}

func NewCourseService(repository secondary.CoursePort) primary.CoursePort {
	return &courseServices{repository}
}

func (this *courseServices) List(_filters filters.CourseFilters) ([]course.Course, errors.Error) {
	return this.repository.List(_filters)
}

func (this *courseServices) Get(id uuid.UUID) (course.Course, errors.Error) {
	return this.repository.Get(id)
}

func (this *courseServices) Create(data course.Course) (*uuid.UUID, errors.Error) {
	return this.repository.Create(data)
}

func (this *courseServices) Update(data course.Course) errors.Error {
	return this.repository.Update(data)
}

func (this *courseServices) Delete(id uuid.UUID) errors.Error {
	return this.repository.Delete(id)
}
