package query

const (
	TimeRecordID            = "id"
	TimeRecordDate          = "date"
	TimeRecordEntryTime     = "entry_time"
	TimeRecordExitTime      = "exit_time"
	TimeRecordLocation      = "location"
	TimeRecordIsOffSite     = "is_off_site"
	TimeRecordJustification = "justification"
	TimeRecordStudentID     = "student_id"
)
type TimeRecordQueryBuilder interface {
	Select() TimeRecordQuerySelectBuilder
	Insert() string
	Update() string
	Delete() string
}

type timeRecordQueryBuilder struct{}

func TimeRecord() TimeRecordQueryBuilder {
	return &timeRecordQueryBuilder{}
}

func (*timeRecordQueryBuilder) Select() TimeRecordQuerySelectBuilder {
	return &timeRecordQuerySelectBuilder{}
}

func (*timeRecordQueryBuilder) Insert() string {
	return `
		INSERT INTO time_record (
			id, date, entry_time, exit_time, location, is_off_site, justification, student_id, created_at
		) VALUES (
			DEFAULT, $1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP
		) RETURNING id
	`
}

func (*timeRecordQueryBuilder) Update() string {
	return `
		UPDATE time_record
		SET
			date = $2,
			entry_time = $3,
			exit_time = $4,
			location = $5,
			is_off_site = $6,
			justification = $7,
			student_id = $8,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
		RETURNING id
	`
}

func (*timeRecordQueryBuilder) Delete() string {
	return `
		UPDATE time_record
		SET
			deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
}

type TimeRecordQuerySelectBuilder interface {
	All() string
	ByID() string
}

type timeRecordQuerySelectBuilder struct{}

func (timeRecordQuerySelectBuilder *timeRecordQuerySelectBuilder) All() string {
	return `
		SELECT
			id AS time_record_id,
			date AS time_record_date,
			entry_time AS time_record_entry_time,
			exit_time AS time_record_exit_time,
			location AS time_record_location,
			is_off_site AS time_record_is_off_site,
			justification AS time_record_justification,
			student_id AS time_record_student_id
		FROM time_record
		WHERE deleted_at IS NULL
		ORDER BY date ASC
	`
}

func (timeRecordQuerySelectBuilder *timeRecordQuerySelectBuilder) ByID() string {
	return `
		SELECT
			id AS time_record_id,
			date AS time_record_date,
			entry_time AS time_record_entry_time,
			exit_time AS time_record_exit_time,
			location AS time_record_location,
			is_off_site AS time_record_is_off_site,
			justification AS time_record_justification,
			student_id AS time_record_student_id
		FROM time_record
		WHERE id = $1 AND deleted_at IS NULL
	`
}
