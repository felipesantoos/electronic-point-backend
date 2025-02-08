package query

const (
	TimeRecordID            = "time_record_id"
	TimeRecordDate          = "time_record_date"
	TimeRecordEntryTime     = "time_record_entry_time"
	TimeRecordExitTime      = "time_record_exit_time"
	TimeRecordLocation      = "time_record_location"
	TimeRecordIsOffSite     = "time_record_is_off_site"
	TimeRecordJustification = "time_record_justification"
	TimeRecordStudentID     = "time_record_student_id"
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
			date, entry_time, exit_time, location, is_off_site, justification, student_id
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
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
			time_record.id AS time_record_id,
			time_record.date AS time_record_date,
			time_record.entry_time AS time_record_entry_time,
			time_record.exit_time AS time_record_exit_time,
			time_record.location AS time_record_location,
			time_record.is_off_site AS time_record_is_off_site,
			time_record.justification AS time_record_justification,
			time_record.student_id AS time_record_student_id,
			time_record_status.id AS time_record_status_id,
			time_record_status.name AS time_record_status_name
		FROM time_record
			INNER JOIN student ON student.id = time_record.student_id
			INNER JOIN student_linked_to_teacher ON student_linked_to_teacher.student_id = student.id
			INNER JOIN person teacher ON teacher.id = student_linked_to_teacher.teacher_id
			INNER JOIN time_record_status_movement 
				ON time_record_status_movement.time_record_id = time_record.id
				AND time_record_status_movement.terminated_at IS NULL
			INNER JOIN time_record_status ON time_record_status.id = time_record_status_movement.status_id
		WHERE time_record.deleted_at IS NULL
			AND time_record.student_id = COALESCE($1, time_record.student_id)
			AND time_record.date BETWEEN COALESCE($2, time_record.date) AND COALESCE($3, time_record.date)
			AND teacher.id = COALESCE($4, teacher.id)
		ORDER BY time_record.date ASC
	`
}

func (timeRecordQuerySelectBuilder *timeRecordQuerySelectBuilder) ByID() string {
	return `
		SELECT
			time_record.id AS time_record_id,
			time_record.date AS time_record_date,
			time_record.entry_time AS time_record_entry_time,
			time_record.exit_time AS time_record_exit_time,
			time_record.location AS time_record_location,
			time_record.is_off_site AS time_record_is_off_site,
			time_record.justification AS time_record_justification,
			time_record.student_id AS time_record_student_id,
			time_record_status.id AS time_record_status_id,
			time_record_status.name AS time_record_status_name
		FROM time_record
			INNER JOIN student ON student.id = time_record.student_id
			INNER JOIN student_linked_to_teacher ON student_linked_to_teacher.student_id = student.id
			INNER JOIN person teacher ON teacher.id = student_linked_to_teacher.teacher_id
			INNER JOIN time_record_status_movement 
				ON time_record_status_movement.time_record_id = time_record.id
				AND time_record_status_movement.terminated_at IS NULL
			INNER JOIN time_record_status ON time_record_status.id = time_record_status_movement.status_id
		WHERE time_record.id = $1 
			AND time_record.deleted_at IS NULL
			AND student.id = COALESCE($2, student.id)
			AND teacher.id = COALESCE($3, teacher.id)
	`
}
