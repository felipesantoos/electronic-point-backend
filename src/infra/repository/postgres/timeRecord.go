package postgres

import (
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/timeRecord"
	"backend_template/src/core/interfaces/adapters"
	"backend_template/src/core/messages"
	"backend_template/src/infra/repository"
	"backend_template/src/infra/repository/postgres/query"
	"backend_template/src/infra/repository/postgres/queryObject"

	"github.com/google/uuid"
)

type timeRecordRepository struct{}

func NewTimeRecordRepository() adapters.TimeRecordAdapter {
	return &timeRecordRepository{}
}

func (r timeRecordRepository) Create(tr timeRecord.TimeRecord) (*uuid.UUID, errors.Error) {
	var timeRecordID uuid.UUID
	rows, err := repository.Queryx(query.TimeRecord().Insert(),
		tr.Date,
		tr.EntryTime,
		tr.ExitTime,
		tr.Location,
		tr.IsOffSite,
		tr.Justification,
		tr.StudentID,
	)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	rows.Scan(&timeRecordID)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return &timeRecordID, nil
}

func (r timeRecordRepository) Update(tr timeRecord.TimeRecord) errors.Error {
	_, err := repository.ExecQuery(query.TimeRecord().Update(),
		tr.ID,
		tr.Date,
		tr.EntryTime,
		tr.ExitTime,
		tr.Location,
		tr.IsOffSite,
		tr.Justification,
		tr.StudentID,
	)
	if err != nil {
		logger.Error().Msg(err.String())
		return errors.NewUnexpected()
	}
	return nil
}

func (r timeRecordRepository) Delete(id uuid.UUID) errors.Error {
	_, err := repository.ExecQuery(query.TimeRecord().Delete(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return errors.NewUnexpected()
	}
	return nil
}

func (r timeRecordRepository) List() ([]timeRecord.TimeRecord, errors.Error) {
	rows, err := repository.Queryx(query.TimeRecord().Select().All())
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	tr, err := queryObject.TimeRecord().FromRows(rows)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return tr, nil
}

func (r timeRecordRepository) Get(id uuid.UUID) (timeRecord.TimeRecord, errors.Error) {
	rows, err := repository.Queryx(query.TimeRecord().Select().ByID(), id)
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
	tr, err := queryObject.TimeRecord().FromMap(serializedTimeRecord)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return tr, nil
}
