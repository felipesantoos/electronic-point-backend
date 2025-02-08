package postgres

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/core/domain/timeRecordStatus"
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

type timeRecordRepository struct{}

func NewTimeRecordRepository() secondary.TimeRecordPort {
	return &timeRecordRepository{}
}

func (this timeRecordRepository) Create(_timeRecord timeRecord.TimeRecord) (*uuid.UUID, errors.Error) {
	transaction, err := repository.BeginTransaction()
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	defer transaction.CloseConn()
	id, err := txQueryRowReturningID(transaction, query.TimeRecord().Insert(),
		_timeRecord.Date(),
		_timeRecord.EntryTime(),
		_timeRecord.ExitTime(),
		_timeRecord.Location(),
		_timeRecord.IsOffSite(),
		_timeRecord.Justification(),
		_timeRecord.StudentID(),
	)
	if err != nil {
		logger.Error().Msg(err.String())
		if strings.Contains(err.String(), constraints.TimeRecordStudentFK) {
			return nil, errors.NewValidationFromString(messages.StudentNotFoundErrorMessage)
		}
		return nil, errors.NewUnexpected()
	}
	timeRecordID, convErr := uuid.Parse(id)
	if convErr != nil {
		logger.Error().Msg(convErr.Error())
		return nil, errors.NewUnexpected()
	}
	_, movementErr := txQueryRowReturningID(transaction, query.TimeRecordStatusMovement().Insert(),
		timeRecordID,
		timeRecordStatus.Pending.ID(),
		_timeRecord.StudentID(),
		nil,
	)
	if movementErr != nil {
		logger.Error().Msg(movementErr.String())
		return nil, errors.NewUnexpected()
	}
	err = transaction.Commit()
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return &timeRecordID, nil
}

func (this timeRecordRepository) Update(_timeRecord timeRecord.TimeRecord) errors.Error {
	_, err := repository.ExecQuery(query.TimeRecord().Update(), _timeRecord.ID(),
		_timeRecord.Date(), _timeRecord.EntryTime(), _timeRecord.ExitTime(), _timeRecord.Location(),
		_timeRecord.IsOffSite(), _timeRecord.Justification(), _timeRecord.StudentID(),
	)
	if err != nil {
		logger.Error().Msg(err.String())
		if strings.Contains(err.String(), constraints.TimeRecordStudentFK) {
			return errors.NewValidationFromString(messages.StudentNotFoundErrorMessage)
		}
		return errors.NewUnexpected()
	}
	return nil
}

func (this timeRecordRepository) Delete(id uuid.UUID) errors.Error {
	_, err := repository.ExecQuery(query.TimeRecord().Delete(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return errors.NewUnexpected()
	}
	return nil
}

func (this timeRecordRepository) List(_filters filters.TimeRecordFilters) ([]timeRecord.TimeRecord, errors.Error) {
	rows, err := repository.Queryx(query.TimeRecord().Select().All(), _filters.StudentID,
		_filters.StartDate, _filters.EndDate, _filters.TeacherID)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	defer rows.Close()
	timeRecords, err := queryObject.TimeRecord().FromRows(rows)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return timeRecords, nil
}

func (this timeRecordRepository) Get(id uuid.UUID, _filters filters.TimeRecordFilters) (timeRecord.TimeRecord, errors.Error) {
	rows, err := repository.Queryx(query.TimeRecord().Select().ByID(), id, _filters.StudentID, _filters.TeacherID)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.NewFromString(messages.TimeRecordNotFoundErrorMessage)
	}
	var serializedTimeRecord = map[string]interface{}{}
	nativeError := rows.MapScan(serializedTimeRecord)
	if nativeError != nil {
		logger.Error().Msg(nativeError.Error())
		return nil, errors.NewUnexpected()
	}
	_timeRecord, err := queryObject.TimeRecord().FromMap(serializedTimeRecord)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return _timeRecord, nil
}
