package gocommon

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//RandInt generate random number inr range
func RandInt(min int, max int) int {
	return rand.Intn(max-min) + min
}

//RandInt32 generate random number inr range
func RandInt32(min int32, max int32) int32 {
	return rand.Int31n(max-min) + min
}

//RandInt64 generate random number inr range
func RandInt64(min int64, max int64) int64 {
	return rand.Int63n(max-min) + min
}
