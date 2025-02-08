package response

import (
	"eletronic_point/src/core/domain/timeRecord"
	"time"

	"github.com/google/uuid"
)

type TimeRecord struct {
	ID               uuid.UUID        `json:"id"`
	Date             time.Time        `json:"date"`
	EntryTime        time.Time        `json:"entry_time"`
	ExitTime         *time.Time       `json:"exit_time"`
	Location         string           `json:"location"`
	IsOffSite        bool             `json:"is_off_site"`
	Justification    *string          `json:"justification"`
	StudentID        uuid.UUID        `json:"student_id"`
	TimeRecordStatus TimeRecordStatus `json:"time_record_status"`
}

type timeRecordBuilder struct{}

func TimeRecordBuilder() *timeRecordBuilder {
	return &timeRecordBuilder{}
}

func (*timeRecordBuilder) BuildFromDomain(data timeRecord.TimeRecord) TimeRecord {
	return TimeRecord{
		ID:               data.ID(),
		Date:             data.Date(),
		EntryTime:        data.EntryTime(),
		ExitTime:         data.ExitTime(),
		Location:         data.Location(),
		IsOffSite:        data.IsOffSite(),
		Justification:    data.Justification(),
		StudentID:        data.StudentID(),
		TimeRecordStatus: TimeRecordStatusBuilder().BuildFromDomain(data.TimeRecordStatus()),
	}
}

func (*timeRecordBuilder) BuildFromDomainList(data []timeRecord.TimeRecord) []TimeRecord {
	timeRecords := make([]TimeRecord, 0)
	for _, record := range data {
		timeRecords = append(timeRecords, TimeRecordBuilder().BuildFromDomain(record))
	}
	return timeRecords
}
