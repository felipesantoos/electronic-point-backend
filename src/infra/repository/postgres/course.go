package postgres

import (
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/infra/repository/postgres/queryObject"

	"github.com/google/uuid"
)

type courseRepository struct{}

func NewCourseRepository() secondary.CoursePort {
	return &courseRepository{}
}

func (this courseRepository) List(_filters filters.CourseFilters) ([]course.Course, errors.Error) {
	rows, err := repository.Queryx(query.Course().Select().All(), _filters.Name)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	defer rows.Close()
	courses, err := queryObject.Course().FromRows(rows)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return courses, nil
}

func (this courseRepository) Get(id uuid.UUID) (course.Course, errors.Error) {
	rows, err := repository.Queryx(query.Course().Select().ByID(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	defer rows.Close()
	courses, err := queryObject.Course().FromRows(rows)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	if len(courses) == 0 {
		return nil, errors.NewFromString("course not found")
	}
	return courses[0], nil
}

func (this courseRepository) Create(data course.Course) (*uuid.UUID, errors.Error) {
	return execQueryReturningID(query.Course().Insert(), data.Name())
}

func (this courseRepository) Update(data course.Course) errors.Error {
	return defaultExecQuery(query.Course().Update(), data.ID(), data.Name())
}

func (this courseRepository) Delete(id uuid.UUID) errors.Error {
	return defaultExecQuery(query.Course().Delete(), id)
}
