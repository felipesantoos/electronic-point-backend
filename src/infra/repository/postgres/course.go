package postgres

import (
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/infra/repository/postgres/queryObject"
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
