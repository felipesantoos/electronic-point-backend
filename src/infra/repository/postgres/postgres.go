package postgres

import (
	"dit_backend/src/infra"
	"dit_backend/src/infra/repository"
)

func defaultExecQuery(sqlQuery string, args ...interface{}) infra.Error {
	result, err := repository.ExecQuery(sqlQuery, args...)
	if err != nil {
		return err
	} else if rowsAff, err := result.RowsAffected(); err != nil {
		return infra.NewInternalSourceErr(err)
	} else if rowsAff == 0 {
		return infra.NewUnexpectedSourceErr()
	}
	return nil
}

func txQueryRowReturningID(tx *repository.SQLTransaction, sqlQuery string, args ...interface{}) (string, infra.Error) {
	row := tx.QueryRow(sqlQuery, args...)
	if err := row.Err(); err != nil {
		rollBackErr := tx.Rollback(err)
		if rollBackErr != nil {
			return "", repository.TranslateError(rollBackErr)
		}
		return "", repository.TranslateError(row.Err())
	}
	var strUUID = ""
	scanErr := row.Scan(&strUUID)
	if scanErr != nil {
		rollBackErr := tx.Rollback(row.Err())
		if rollBackErr != nil {
			return "", repository.TranslateError(rollBackErr)
		}
		return "", repository.TranslateError(scanErr)
	}
	return strUUID, nil
}
