package timeRecordStatus

import "errors"

type TimeRecordStatusValues int

const (
	Pending TimeRecordStatusValues = iota
	Approved
	Disapproved
)

var statusIDs = []string{
	"52613242-6b50-490a-9b4c-90cc3f263e9a",
	"faa4a69d-fe41-4ffe-b8d0-f752085f016a",
	"7f58a284-c8a5-4f89-a18e-320e8ea8960f",
}

var statusNames = []string{
	"pending",
	"approved",
	"disapproved",
}

func (s TimeRecordStatusValues) ID() string {
	return statusIDs[s]
}

func (s TimeRecordStatusValues) Name() string {
	return statusNames[s]
}

func ParseStatusByID(id string) (TimeRecordStatusValues, error) {
	for i, statusID := range statusIDs {
		if statusID == id {
			return TimeRecordStatusValues(i), nil
		}
	}
	return -1, errors.New("invalid status ID")
}

func ParseStatusByName(name string) (TimeRecordStatusValues, error) {
	for i, statusName := range statusNames {
		if statusName == name {
			return TimeRecordStatusValues(i), nil
		}
	}
	return -1, errors.New("invalid status name")
}
