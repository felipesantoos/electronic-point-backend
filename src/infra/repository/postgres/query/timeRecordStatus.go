package query

const (
	TimeRecordStatusID   = "time_record_status_id"
	TimeRecordStatusName = "time_record_status_name"
)

type TimeRecordStatusQueryBuilder interface {
	Select() TimeRecordStatusQuerySelectBuilder
}

type TimeRecordStatusQuerySelectBuilder interface {
	All() string
	ByID() string
}

type timeRecordStatusQueryBuilder struct{}

func TimeRecordStatus() TimeRecordStatusQueryBuilder {
	return &timeRecordStatusQueryBuilder{}
}

func (*timeRecordStatusQueryBuilder) Select() TimeRecordStatusQuerySelectBuilder {
	return &timeRecordStatusQuerySelectBuilder{}
}

type timeRecordStatusQuerySelectBuilder struct{}

func (*timeRecordStatusQuerySelectBuilder) All() string {
	return `
        SELECT time_record_status.id AS time_record_status_id, 
               time_record_status.name AS time_record_status_name
        FROM time_record_status;
    `
}

func (*timeRecordStatusQuerySelectBuilder) ByID() string {
	return `
        SELECT time_record_status.id AS time_record_status_id, 
               time_record_status.name AS time_record_status_name
        FROM time_record_status
        WHERE time_record_status.id = $1;
    `
}
