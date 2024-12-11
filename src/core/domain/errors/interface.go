package errors

type Error interface {
	String() string
	Messages() []string
	CausedInternally() bool
	CausedByValidation() bool
	CausedByClient() bool
	CausedByForbiddenAccess() bool
	CausedByConflict() bool
	Metadata() map[string]interface{}
	ValidationMessagesByMetadataFields(field []string) []string
}
