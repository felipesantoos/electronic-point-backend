package response

import (
	"eletronic_point/src/core/domain/student"
)

type TimeRecordsByStudent struct {
	StudentList
	TimeRecords []TimeRecord `json:"time_records"`
}

type timeRecordsByStudentBuilder struct{}

func TimeRecordsByStudentBuilder() *timeRecordsByStudentBuilder {
	return &timeRecordsByStudentBuilder{}
}

func (*timeRecordsByStudentBuilder) BuildFromDomain(data student.Student) TimeRecordsByStudent {
	return TimeRecordsByStudent{
		StudentList: StudentListBuilder().BuildFromDomain(data),
		TimeRecords: TimeRecordBuilder().BuildFromDomainList(data.FrequencyHistory()),
	}
}

func (*timeRecordsByStudentBuilder) BuildFromDomainList(data []student.Student) []TimeRecordsByStudent {
	timeRecordsByStudents := make([]TimeRecordsByStudent, 0)
	for _, record := range data {
		timeRecordsByStudents = append(timeRecordsByStudents, TimeRecordsByStudentBuilder().BuildFromDomain(record))
	}
	return timeRecordsByStudents
}
