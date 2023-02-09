package domain

import (
	"fmt"
	"strings"
)

func BuildMapWithParentName(data map[string]interface{}, parentName string) map[string]interface{} {
	var newData = map[string]interface{}{}
	for k, v := range data {
		if strings.Contains(k, parentName) {
			newData[strings.ReplaceAll(k, fmt.Sprintf("%s_", parentName), "")] = v
		}
	}
	return newData
}
