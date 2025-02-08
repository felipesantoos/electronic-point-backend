package query

const (
	TimeRecordStatusMovementID           = "time_record_status_movement_id"
	TimeRecordStatusMovementTimeRecordID = "time_record_status_movement_time_record_id"
	TimeRecordStatusMovementStatusID     = "time_record_status_movement_status_id"
	TimeRecordStatusMovementPersonID     = "time_record_status_movement_person_id"
	TimeRecordStatusMovementComments     = "time_record_status_movement_comments"
	TimeRecordStatusMovementCreatedAt    = "time_record_status_movement_created_at"
	TimeRecordStatusMovementTerminatedAt = "time_record_status_movement_terminated_at"
	TimeRecordStatusMovementUpdatedAt    = "time_record_status_movement_updated_at"
	TimeRecordStatusMovementDeletedAt    = "time_record_status_movement_deleted_at"
)

type TimeRecordStatusMovementQueryBuilder interface {
	Select() TimeRecordStatusMovementQuerySelectBuilder
	Insert() string
	Terminate() string
}

type timeRecordStatusMovementQueryBuilder struct{}

func TimeRecordStatusMovement() TimeRecordStatusMovementQueryBuilder {
	return &timeRecordStatusMovementQueryBuilder{}
}

func (*timeRecordStatusMovementQueryBuilder) Select() TimeRecordStatusMovementQuerySelectBuilder {
	return &timeRecordStatusMovementQuerySelectBuilder{}
}

func (*timeRecordStatusMovementQueryBuilder) Insert() string {
	return `
		INSERT INTO time_record_status_movement (
			time_record_id, status_id, person_id, comments
		) VALUES (
			$1, $2, $3, $4
		) RETURNING id
	`
}

func (*timeRecordStatusMovementQueryBuilder) Terminate() string {
	return `
		UPDATE time_record_status_movement
		SET
			terminated_at = CURRENT_TIMESTAMP
		WHERE time_record_id = $1
	`
}

type TimeRecordStatusMovementQuerySelectBuilder interface {
	All() string
	ByID() string
}

type timeRecordStatusMovementQuerySelectBuilder struct{}

func (*timeRecordStatusMovementQuerySelectBuilder) All() string {
	return `
		SELECT
			time_record_status_movement.id AS time_record_status_movement_id,
			time_record_status_movement.time_record_id AS time_record_status_movement_time_record_id,
			time_record_status_movement.status_id AS time_record_status_movement_status_id,
			time_record_status_movement.person_id AS time_record_status_movement_person_id,
			time_record_status_movement.comments AS time_record_status_movement_comments,
			time_record_status_movement.created_at AS time_record_status_movement_created_at,
			time_record_status_movement.terminated_at AS time_record_status_movement_terminated_at,
			time_record_status_movement.updated_at AS time_record_status_movement_updated_at,
			time_record_status_movement.deleted_at AS time_record_status_movement_deleted_at
		FROM time_record_status_movement
		WHERE time_record_status_movement.deleted_at IS NULL
		ORDER BY time_record_status_movement.created_at ASC
	`
}

func (*timeRecordStatusMovementQuerySelectBuilder) ByID() string {
	return `
		SELECT
			time_record_status_movement.id AS time_record_status_movement_id,
			time_record_status_movement.time_record_id AS time_record_status_movement_time_record_id,
			time_record_status_movement.status_id AS time_record_status_movement_status_id,
			time_record_status_movement.person_id AS time_record_status_movement_person_id,
			time_record_status_movement.comments AS time_record_status_movement_comments,
			time_record_status_movement.created_at AS time_record_status_movement_created_at,
			time_record_status_movement.terminated_at AS time_record_status_movement_terminated_at,
			time_record_status_movement.updated_at AS time_record_status_movement_updated_at,
			time_record_status_movement.deleted_at AS time_record_status_movement_deleted_at
		FROM time_record_status_movement
		WHERE time_record_status_movement.id = $1 
			AND time_record_status_movement.deleted_at IS NULL
	`
}
