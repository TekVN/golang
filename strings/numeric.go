package strings

import (
	"strconv"
	"strings"
)

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func IsFloat(s string) bool {
	if strings.Count(s, ".") != 1 {
		return false
	}

	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
