package gocommon

import (
	"strings"
)

var replacerChar = strings.NewReplacer("ﻙ", "ک", "ك", "ک", "ي", "ی")
var replacerCharAndNumber = strings.NewReplacer("ﻙ", "ک", "ك", "ک", "ي", "ی", "0", "۰", "1", "۱", "2", "۲", "3", "۳", "4", "۴", "5", "۵", "6", "۶", "7", "۷", "8", "۸", "9", "۹")

// ToPersian convert string to standard persian unicode
func ToPersian(inputString string) string {
	return replacerChar.Replace(inputString)
}

// ToPersianWithNumber convert string to standard persian unicode and convert numbers to
func ToPersianWithNumber(inputString string) string {
	return replacerCharAndNumber.Replace(inputString)
}
