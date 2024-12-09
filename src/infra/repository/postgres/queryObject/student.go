package queryObject

import (
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/student"
	"backend_template/src/infra/repository/postgres/query"
	"backend_template/src/utils"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type StudentObjectBuilder interface {
	FromMap(data map[string]interface{}) (student.Student, errors.Error)
	FromRows(rows *sqlx.Rows) ([]student.Student, errors.Error)
}

type studentQueryObjectBuilder struct{}

func Student() StudentObjectBuilder {
	return &studentQueryObjectBuilder{}
}

func (s *studentQueryObjectBuilder) FromMap(data map[string]interface{}) (student.Student, errors.Error) {
	id, err := uuid.Parse(string(data[query.StudentID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	name := fmt.Sprint(data[query.StudentName])
	registration := fmt.Sprint(data[query.StudentRegistration])
	profilePicture := utils.GetNullableValue[string](data[query.StudentName])
	institution := fmt.Sprint(data[query.StudentInstitution])
	course := fmt.Sprint(data[query.StudentCourse])
	internshipLocation := fmt.Sprint(data[query.StudentInternshipLocation])
	totalWorkload, err := strconv.Atoi(fmt.Sprint(data[query.StudentTotalWorkload]))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	_student, validationError := student.NewBuilder().
		WithID(id).
		WithName(name).
		WithRegistration(registration).
		WithProfilePicture(profilePicture).
		WithInstitution(institution).
		WithCourse(course).
		WithInternshipLocation(internshipLocation).
		WithTotalWorkload(totalWorkload).
		Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	return _student, nil
}

func (s *studentQueryObjectBuilder) FromRows(rows *sqlx.Rows) ([]student.Student, errors.Error) {
	if rows == nil {
		err := errors.NewFromString("row value cannot be nil")
		logger.Error().Msg(err.String())
		return nil, err
	}
	defer rows.Close()
	students := make([]student.Student, 0)
	for rows.Next() {
		var serializedStudent = map[string]interface{}{}
		nativeError := rows.MapScan(serializedStudent)
		if nativeError != nil {
			logger.Error().Msg(nativeError.Error())
			return nil, errors.New(nativeError)
		}
		_student, err := s.FromMap(serializedStudent)
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, err
		}
		students = append(students, _student)
	}
	return students, nil
}
