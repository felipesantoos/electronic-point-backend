package repository

import (
	"database/sql"
	"dit_backend/src/core/utils"
	"dit_backend/src/infra"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

var logger = infra.Logger()
var keyConstraintCompiler = regexp.MustCompile(`^.+?_(.*)_key`)
var primaryKeyConstraintCompiler = regexp.MustCompile(`"(\w.+)_pkey?`)
var foreignKeyConstraintCompiler = regexp.MustCompile(`(\w.+)_fk`)

type SQLTransaction struct {
	QueryRow  func(query string, args ...interface{}) *sql.Row
	ExecQuery func(query string, args ...interface{}) (sql.Result, infra.Error)
	Rollback  func(err error) error
	CloseConn func() error
	Commit    func() error
}

func getDatabaseSchema() string {
	return utils.GetenvWithDefault("DATABASE_SCHEMA", "postgres")
}

func getDatabaseURI() string {
	schema := getDatabaseSchema()
	user := os.Getenv("DATABASE_USER")
	pwd := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("DATABASE_NAME")
	authentication := fmt.Sprintf("%s:%s", user, pwd)
	dst := fmt.Sprintf("%s:%s/%s", host, port, name)
	sslMode := utils.GetenvWithDefault("DATABASE_SSL_MODE", "disable")
	return fmt.Sprintf("%s://%s@%s?sslmode=%s", schema, authentication, dst, sslMode)
}

func getConnection() (*sqlx.DB, infra.Error) {
	schema := getDatabaseSchema()
	connection, err := sqlx.Open(schema, getDatabaseURI())
	if err != nil {
		logger.Error().Msg("Error while acessing database: " + err.Error())
		return nil, TranslateError(err)
	}
	connection.SetConnMaxLifetime(time.Minute * 3)
	connection.SetMaxOpenConns(10)
	connection.SetMaxIdleConns(10)
	return connection, nil
}

func closeConnection(conn *sqlx.DB) error {
	err := conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func internalRollback(tx *sql.Tx, err error) error {
	rollbackErr := tx.Rollback()
	if rollbackErr != nil {
		return rollbackErr
	}
	return err
}

func internalCommit(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func Queryx(sqlQuery string, args ...interface{}) (*sqlx.Rows, infra.Error) {
	conn, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer closeConnection(conn)
	rows, queryErr := conn.Queryx(sqlQuery, args...)
	if queryErr != nil {
		return nil, TranslateError(queryErr)
	}
	return rows, nil
}

func ExecQuery(sqlQuery string, args ...interface{}) (sql.Result, infra.Error) {
	conn, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer closeConnection(conn)
	result, queryErr := conn.Exec(sqlQuery, args...)
	if queryErr != nil {
		return nil, TranslateError(queryErr)
	}
	return result, nil
}

func BeginTransaction() (*SQLTransaction, infra.Error) {
	conn, err := getConnection()
	if err != nil {
		return nil, err
	}
	tx, beginErr := conn.Begin()
	if beginErr != nil {
		return nil, TranslateError(beginErr)
	}
	queryRow := func(query string, args ...interface{}) *sql.Row {
		return tx.QueryRow(query, args...)
	}
	execQuery := func(query string, args ...interface{}) (sql.Result, infra.Error) {
		result, err := tx.Exec(query, args...)
		return result, TranslateError(err)
	}
	closeConnection := func() error {
		return closeConnection(conn)
	}
	rollBack := func(err error) error {
		defer closeConnection()
		return internalRollback(tx, err)
	}
	commit := func() error {
		defer closeConnection()
		return internalCommit(tx)
	}
	transaction := &SQLTransaction{
		QueryRow:  queryRow,
		ExecQuery: execQuery,
		Rollback:  rollBack,
		CloseConn: closeConnection,
		Commit:    commit,
	}
	return transaction, nil
}

func TranslateError(err error) infra.Error {
	if err == nil {
		return nil
	}
	logger.Error().Msg(err.Error())
	if strings.Contains(err.Error(), "duplicate key") {
		keyMatch := keyConstraintCompiler.FindStringSubmatch(err.Error())
		var key = ""
		if keyMatch == nil {
			pkeyMatch := primaryKeyConstraintCompiler.FindStringSubmatch(err.Error())
			if pkeyMatch == nil {
				return infra.NewInternalSourceErr(err)
			}
			key = pkeyMatch[1]
		} else {
			key = keyMatch[1]
		}
		return infra.NewSourceErrFromStr(fmt.Sprintf("the value defined for %s is already in use", key))
	} else if strings.Contains(err.Error(), "violates foreign key constraint") {
		keyMatch := foreignKeyConstraintCompiler.FindStringSubmatch(err.Error())
		if keyMatch == nil {
			return infra.NewInternalSourceErr(err)
		}
		key := keyMatch[1]
		return infra.NewSourceErrFromStr(fmt.Sprintf("the key defined for %s doesn't exists", key))
	}
	return infra.NewUnexpectedSourceErr()
}
