package gocommon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// LitToPInt64 Convert litral to pointer
func TestLitToPInt64(t *testing.T) {
	pointerToNumber := LitToPInt64(45)
	assert.Equal(t, int64(45), *pointerToNumber)
}

// LitToPInt32 Convert litral to pointer
func TestLitToPInt32(t *testing.T) {
	pointerToNumber := LitToPInt32(45)
	assert.Equal(t, int32(45), *pointerToNumber)
}

// LitToPInt16 Convert litral to pointer
func TestTestLitToPInt16(t *testing.T) {
	pointerToNumber := LitToPInt16(45)
	assert.Equal(t, int16(45), *pointerToNumber)
}

// LitToPInt8 Convert litral to pointer
func TestLitToPInt8(t *testing.T) {
	pointerToNumber := LitToPInt8(45)
	assert.Equal(t, int8(45), *pointerToNumber)
}

// LitToPInt Convert litral to pointer
func TestLitToPInt(t *testing.T) {
	pointerToNumber := LitToPInt(45)
	assert.Equal(t, int(45), *pointerToNumber)
}

// LitToPFloat32 Convert litral to pointer
func TestLitToPFloat32(t *testing.T) {
	pointerToNumber := LitToPFloat32(45.12)
	assert.Equal(t, float32(45.12), *pointerToNumber)
}

// LitToPFloat64 Convert litral to pointer
func TestLitToPFloat64(t *testing.T) {
	pointerToNumber := LitToPFloat64(45.12)
	assert.Equal(t, float64(45.12), *pointerToNumber)
}

// LitToPBool Convert litral to pointer
func TestTestLitToPBool(t *testing.T) {
	pointerToNumber := LitToPBool(true)
	assert.Equal(t, true, *pointerToNumber)
}
