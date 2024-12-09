package validator

func IsPaginationValid(pagination int) bool {
	if &pagination == nil || pagination <= 0 {
		return false
	}
	return true
}
