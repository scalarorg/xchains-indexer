package common

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseAttributeValue(value string) any {
	result := strings.Trim(value, "\"")
	if strings.HasPrefix(result, "[") && strings.HasSuffix(result, "]") {
		items := strings.Split(result[1:len(result)-1], ",")
		hexValue := "0x"
		for _, item := range items {
			v, e := strconv.Atoi(item)
			if e == nil {
				hexValue = fmt.Sprintf("%s%x", hexValue, v)
			}
		}
		return hexValue
	}
	return result
}
