package validator

import (
	"regexp"
	"strings"
)

var validDDDs = []string{
	"11", "12", "13", "14", "15", "16", "17", "18", "19",
	"21", "22", "24", "27", "28", "31", "32", "33", "34",
	"35", "37", "38", "41", "42", "43", "44", "45", "46",
	"47", "48", "49", "51", "53", "54", "55", "61", "62",
	"63", "64", "65", "66", "67", "68", "69", "71", "73",
	"74", "75", "77", "79", "81", "82", "83", "84", "85",
	"86", "87", "88", "89", "91", "92", "93", "94", "95",
	"96", "97", "98", "99",
}

func IsPhoneValid(phone string) bool {
	if phone == "" || len(phone) == 0 {
		return false
	}

	removeSpecialCharacters := regexp.MustCompile(`[^0-9]+`)

	parsedPhone := removeSpecialCharacters.ReplaceAllString(phone, "")
	islenWithDDDValid := len(parsedPhone) == 11
	isValidDDD := isValidDDD(parsedPhone[:2])

	return islenWithDDDValid && isValidDDD
}

func isValidDDD(separatedDDD string) bool {
	for _, ddd := range validDDDs {
		if strings.EqualFold(ddd, separatedDDD) {
			return true
		}
	}
	return false
}
