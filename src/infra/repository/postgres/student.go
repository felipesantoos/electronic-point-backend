package postgres

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/student"
	"eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/constraints"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/infra/repository/postgres/queryObject"
	"strings"

	"github.com/google/uuid"
)

type studentRepository struct{}

func NewStudentRepository() secondary.StudentPort {
	return &studentRepository{}
}

func (this studentRepository) Create(_student student.Student) (*uuid.UUID, errors.Error) {
	var id uuid.UUID
	rows, err := repository.Queryx(query.Student().Insert(), _student.Name(), _student.Registration(),
		_student.ProfilePicture(), _student.Institution(), _student.Course(), _student.InternshipLocationName(),
		_student.InternshipAddress(), _student.InternshipLocation(), _student.TotalWorkload())
	if err != nil {
		logger.Error().Msg(err.String())
		if strings.Contains(err.String(), constraints.StudentRegistrationUK) {
			return nil, errors.NewConflictFromString(messages.StudentRegistrationIsAlreadyInUseErrorMessage)
		}
		return nil, errors.NewUnexpected()
	}
	if !rows.Next() {
		return nil, errors.NewUnexpected()
	}
	scanError := rows.Scan(&id)
	if scanError != nil {
		logger.Error().Msg(scanError.Error())
		return nil, errors.NewUnexpected()
	}
	return &id, nil
}

func (this studentRepository) Update(_student student.Student) errors.Error {
	_, err := repository.ExecQuery(query.Student().Update(), _student.ID(), _student.Name(), _student.Registration(),
		_student.ProfilePicture(), _student.Institution(), _student.Course(), _student.InternshipLocationName(),
		_student.InternshipAddress(), _student.InternshipLocation(), _student.TotalWorkload())
	if err != nil {
		logger.Error().Msg(err.String())
		if strings.Contains(err.String(), constraints.StudentRegistrationUK) {
			return errors.NewConflictFromString(messages.StudentRegistrationIsAlreadyInUseErrorMessage)
		}
		return errors.NewUnexpected()
	}
	return nil
}

func (this studentRepository) Delete(id uuid.UUID) errors.Error {
	_, err := repository.ExecQuery(query.Student().Delete(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return errors.NewUnexpected()
	}
	return nil
}

func (this studentRepository) List() ([]student.Student, errors.Error) {
	rows, err := repository.Queryx(query.Student().Select().All())
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	students, err := queryObject.Student().FromRows(rows)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	for i := range students {
		timeRecordRepository := NewTimeRecordRepository()
		studentID := students[i].ID()
		_filters := filters.TimeRecordFilters{StudentID: &studentID}
		timeRecords, err := timeRecordRepository.List(_filters)
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, err
		}
		err = students[i].SetFrequencyHistory(timeRecords)
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, err
		}
		workloadCompleted := calculateWorkloadCompleted(timeRecords)
		students[i].SetWorkloadCompleted(workloadCompleted)
		students[i].SetPendingWorkload(calculatePendingWorkload(students[i].TotalWorkload(), workloadCompleted))
	}
	return students, nil
}

func (this studentRepository) Get(id uuid.UUID) (student.Student, errors.Error) {
	rows, err := repository.Queryx(query.Student().Select().ByID(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.NewFromString(messages.StudentNotFoundErrorMessage)
	}
	var serializedStudent = map[string]interface{}{}
	nativeError := rows.MapScan(serializedStudent)
	if nativeError != nil {
		logger.Error().Msg(nativeError.Error())
		return nil, errors.NewUnexpected()
	}
	_student, err := queryObject.Student().FromMap(serializedStudent)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	timeRecordRepository := NewTimeRecordRepository()
	studentID := _student.ID()
	_filters := filters.TimeRecordFilters{StudentID: &studentID}
	timeRecords, err := timeRecordRepository.List(_filters)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	err = _student.SetFrequencyHistory(timeRecords)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	workloadCompleted := calculateWorkloadCompleted(timeRecords)
	_student.SetWorkloadCompleted(workloadCompleted)
	_student.SetPendingWorkload(calculatePendingWorkload(_student.TotalWorkload(), workloadCompleted))
	return _student, nil
}

func calculateWorkloadCompleted(timeRecords []timeRecord.TimeRecord) int {
	sum := 0
	for _, _timeRecord := range timeRecords {
		duration := _timeRecord.ExitTime().Sub(_timeRecord.EntryTime())
		hours := duration.Hours()
		fullHours := int(hours)
		sum += fullHours
	}
	return sum
}

func calculatePendingWorkload(totalWorkload, workloadCompleted int) int {
	return totalWorkload - workloadCompleted
}
