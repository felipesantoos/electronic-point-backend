package messages

const (
	TimeRecordID            = "time record id"
	TimeRecordDate          = "time record date"
	TimeRecordEntryTime     = "time record entry time"
	TimeRecordExitTime      = "time record exit time"
	TimeRecordLocation      = "time record location"
	TimeRecordIsOffSite     = "time record is off site"
	TimeRecordJustification = "time record justification"
	TimeRecordStudentID     = "time record student id"

	TimeRecordIDErrorMessage            = "time record id is invalid"
	TimeRecordDateErrorMessage          = "time record date is invalid"
	TimeRecordEntryTimeErrorMessage     = "time record entry time is invalid"
	TimeRecordExitTimeErrorMessage      = "time record exit time is invalid"
	TimeRecordLocationErrorMessage      = "time record location is invalid"
	TimeRecordIsOffSiteErrorMessage     = "time record off-site status is invalid"
	TimeRecordJustificationErrorMessage = "time record justification is invalid"
	TimeRecordStudentIDErrorMessage     = "time record student id is invalid"

	TimeRecordNotFoundErrorMessage = "time record not found"
	TimeRecordToleranceErrorMessage = "time record entry time exceeds the 30-minute tolerance"
	TimeRecordDailyLimitErrorMessage = "time record daily limit of 5 hours has been reached"
)
