package helper

import "strings"

func RemoveCommas(text string) string {
	return strings.ReplaceAll(text, ",", "")
}
