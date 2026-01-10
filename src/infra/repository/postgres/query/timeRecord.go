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
	TimeRecordInternshipID  = "internship_id"
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
			date, entry_time, exit_time, location, is_off_site, justification, student_id, internship_id
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
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
			internship_id = $9,
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
			time_record_status.name AS time_record_status_name,
			person.name AS person_name,
			student.profile_picture AS student_profile_picture,
			campus.id AS campus_id,
			campus.name AS campus_name,
			institution.id AS institution_id,
			institution.name AS institution_name,
			course.id AS course_id,
			course.name AS course_name,
			internship.id AS internship_id,
			internship.started_in AS internship_started_in,
			internship.ended_in AS internship_ended_in,
			internship.schedule_entry_time AS internship_schedule_entry_time,
			internship.schedule_exit_time AS internship_schedule_exit_time,
			internship_location.id AS internship_location_id,
			internship_location.name AS internship_location_name
		FROM time_record
			INNER JOIN student ON student.id = time_record.student_id
			INNER JOIN person ON person.id = student.id
			INNER JOIN campus ON campus.id = student.campus_id
			INNER JOIN institution ON institution.id = campus.institution_id
			INNER JOIN course ON course.id = student.course_id
			INNER JOIN internship ON internship.id = time_record.internship_id
			LEFT JOIN internship_location ON internship_location.id = internship.internship_location_id
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
			AND time_record_status.id = COALESCE($5, time_record_status.id)
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
			time_record_status.name AS time_record_status_name,
			person.name AS person_name,
			student.profile_picture AS student_profile_picture,
			campus.id AS campus_id,
			campus.name AS campus_name,
			institution.id AS institution_id,
			institution.name AS institution_name,
			course.id AS course_id,
			course.name AS course_name,
			internship.id AS internship_id,
			internship.started_in AS internship_started_in,
			internship.ended_in AS internship_ended_in,
			internship.schedule_entry_time AS internship_schedule_entry_time,
			internship.schedule_exit_time AS internship_schedule_exit_time,
			internship_location.id AS internship_location_id,
			internship_location.name AS internship_location_name
		FROM time_record
			INNER JOIN student ON student.id = time_record.student_id
			INNER JOIN person ON person.id = student.id
			INNER JOIN campus ON campus.id = student.campus_id
			INNER JOIN institution ON institution.id = campus.institution_id
			INNER JOIN course ON course.id = student.course_id
			INNER JOIN internship ON internship.id = time_record.internship_id
			LEFT JOIN internship_location ON internship_location.id = internship.internship_location_id
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
