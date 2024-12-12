package queryObject

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/student"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/utils"
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
	id, err := uuid.Parse(string(data[query.PersonID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	name := fmt.Sprint(data[query.PersonName])
	birthDate := fmt.Sprint(data[query.PersonBirthDate])[:10]
	email := fmt.Sprint(data[query.PersonEmail])
	cpf := fmt.Sprint(data[query.PersonCPF])
	phone := fmt.Sprint(data[query.PersonPhone])
	registration := fmt.Sprint(data[query.StudentRegistration])
	profilePicture := utils.GetNullableValue[string](data[query.StudentProfilePicture])
	institution := fmt.Sprint(data[query.StudentInstitution])
	course := fmt.Sprint(data[query.StudentCourse])
	internshipLocationName := fmt.Sprint(data[query.StudentInternshipLocationName])
	internshipAddress := fmt.Sprint(data[query.StudentInternshipAddress])
	internshipLocation := fmt.Sprint(data[query.StudentInternshipLocation])
	totalWorkload, err := strconv.Atoi(fmt.Sprint(data[query.StudentTotalWorkload]))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	_person, validationError := person.NewBuilder().WithID(id).WithName(name).
		WithBirthDate(birthDate).WithEmail(email).WithCPF(cpf).WithPhone(phone).Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	_student, validationError := student.NewBuilder().
		WithPerson(_person).
		WithRegistration(registration).
		WithProfilePicture(profilePicture).
		WithInstitution(institution).
		WithCourse(course).
		WithInternshipLocationName(internshipLocationName).
		WithInternshipAddress(internshipAddress).
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
