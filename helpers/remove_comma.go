package helpers

import "strings"

func RemoveCommas(text string) string {
	return strings.ReplaceAll(text, ",", "")
}
