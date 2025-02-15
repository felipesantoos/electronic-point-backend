package queryObject

import (
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/infra/repository/postgres/query"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CourseObjectBuilder interface {
	FromMap(data map[string]interface{}) (course.Course, errors.Error)
	FromRows(rows *sqlx.Rows) ([]course.Course, errors.Error)
}

type courseQueryObjectBuilder struct{}

func Course() CourseObjectBuilder {
	return &courseQueryObjectBuilder{}
}

func (this *courseQueryObjectBuilder) FromMap(data map[string]interface{}) (course.Course, errors.Error) {
	id, err := uuid.Parse(string(data[query.CourseID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	name := fmt.Sprint(data[query.CourseName])
	_course, validationError := course.NewBuilder().WithID(id).WithName(name).Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	return _course, nil
}

func (this *courseQueryObjectBuilder) FromRows(rows *sqlx.Rows) ([]course.Course, errors.Error) {
	if rows == nil {
		err := errors.NewFromString("row value cannot be nil")
		logger.Error().Msg(err.String())
		return nil, err
	}
	defer rows.Close()
	courses := make([]course.Course, 0)
	for rows.Next() {
		var serializedCourse = map[string]interface{}{}
		nativeError := rows.MapScan(serializedCourse)
		if nativeError != nil {
			logger.Error().Msg(nativeError.Error())
			return nil, errors.New(nativeError)
		}
		_course, err := this.FromMap(serializedCourse)
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, err
		}
		courses = append(courses, _course)
	}
	return courses, nil
}
