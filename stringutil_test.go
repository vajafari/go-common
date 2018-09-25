package gocommon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestToPersian convert string to standard persian unicode
func TestToPersian(t *testing.T) {
	var stringTable = []struct {
		input    string // input
		expected string // expected result
	}{
		{"ﻙجا", "کجا"},
		{"كجا", "کجا"},
		{"وحيد", "وحید"},
		{"ﻙجايي", "کجایی"},
		{"كجايی", "کجایی"},
		{"كجایي", "کجایی"},
	}
	for _, tt := range stringTable {
		result := ToPersian(tt.input)
		assert.Equal(t, tt.expected, result)
	}
}

func TestToPersianWithNumber(t *testing.T) {
	assert.Equal(t, 123, 123)
	var stringTable = []struct {
		input    string // input
		expected string // expected result
	}{
		{"ﻙجا", "کجا"},
		{"كجا", "کجا"},
		{"وحيد", "وحید"},
		{"ﻙجايي", "کجایی"},
		{"كجايی", "کجایی"},
		{"كجایي", "کجایی"},
		{"0123456789کجایي", "۰۱۲۳۴۵۶۷۸۹کجایی"},
	}
	for _, tt := range stringTable {
		result := ToPersianWithNumber(tt.input)
		assert.Equal(t, tt.expected, result)
	}
}
