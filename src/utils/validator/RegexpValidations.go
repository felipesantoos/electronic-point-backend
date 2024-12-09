package validator

import "regexp"

func HostAddressIsValid(hostAddress string) bool {
	ipv4Regex := `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	re, _ := regexp.Compile(ipv4Regex)
	return re.MatchString(hostAddress) || len(hostAddress) > 0
}

func TextHasOnlyNumbers(text string) bool {
	regex := regexp.MustCompile("^[0-9]+$")
	return regex.MatchString(text)
}
