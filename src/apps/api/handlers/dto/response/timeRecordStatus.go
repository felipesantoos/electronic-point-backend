package response

import (
	"eletronic_point/src/core/domain/timeRecordStatus"

	"github.com/google/uuid"
)

type TimeRecordStatus struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type timeRecordStatusBuilder struct{}

func TimeRecordStatusBuilder() *timeRecordStatusBuilder {
	return &timeRecordStatusBuilder{}
}

func (*timeRecordStatusBuilder) BuildFromDomain(data timeRecordStatus.TimeRecordStatus) TimeRecordStatus {
	return TimeRecordStatus{
		ID:   data.ID(),
		Name: data.Name(),
	}
}

func (*timeRecordStatusBuilder) BuildFromDomainList(data []timeRecordStatus.TimeRecordStatus) []TimeRecordStatus {
	timeRecordStatuses := make([]TimeRecordStatus, 0)
	for _, status := range data {
		timeRecordStatuses = append(timeRecordStatuses, TimeRecordStatusBuilder().BuildFromDomain(status))
	}
	return timeRecordStatuses
}
