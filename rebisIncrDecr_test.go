package rebis

import (
	"testing"
	"time"
)

func TestIncrementExpiration(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tint", 1, 5*time.Millisecond)
	<-time.After(10 * time.Millisecond)
	err = tc.Increment("tint", 2)
	if err == nil {
		t.Error("Error found expired:", err)
	}
}

func TestIncrementWithInt(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tint", 1, DefaultExpiration)
	err = tc.Increment("tint", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tint")
	if !found {
		t.Error("tint was not found")
	}
	if x.(int) != 3 {
		t.Error("tint is not 3:", x)
	}
	tc.Set("tstr", "str", DefaultExpiration)
	err = tc.Increment("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestIncrementWithInt8(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tint8", int8(1), DefaultExpiration)
	err = tc.Increment("tint8", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tint8")
	if !found {
		t.Error("tint8 was not found")
	}
	if x.(int8) != 3 {
		t.Error("tint8 is not 3:", x)
	}
}

func TestIncrementWithInt16(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tint16", int16(1), DefaultExpiration)
	err = tc.Increment("tint16", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tint16")
	if !found {
		t.Error("tint16 was not found")
	}
	if x.(int16) != 3 {
		t.Error("tint16 is not 3:", x)
	}
}

func TestIncrementWithInt32(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tint32", int32(1), DefaultExpiration)
	err = tc.Increment("tint32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tint32")
	if !found {
		t.Error("tint32 was not found")
	}
	if x.(int32) != 3 {
		t.Error("tint32 is not 3:", x)
	}
}

func TestIncrementWithInt64(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tint64", int64(1), DefaultExpiration)
	err = tc.Increment("tint64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tint64")
	if !found {
		t.Error("tint64 was not found")
	}
	if x.(int64) != 3 {
		t.Error("tint64 is not 3:", x)
	}
}

func TestIncrementWithUint(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tuint", uint(1), DefaultExpiration)
	err = tc.Increment("tuint", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tuint")
	if !found {
		t.Error("tuint was not found")
	}
	if x.(uint) != 3 {
		t.Error("tuint is not 3:", x)
	}
}

func TestIncrementWithUintptr(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tuintptr", uintptr(1), DefaultExpiration)
	err = tc.Increment("tuintptr", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}

	x, found := tc.Get("tuintptr")
	if !found {
		t.Error("tuintptr was not found")
	}
	if x.(uintptr) != 3 {
		t.Error("tuintptr is not 3:", x)
	}
}

func TestIncrementWithUint8(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tuint8", uint8(1), DefaultExpiration)
	err = tc.Increment("tuint8", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tuint8")
	if !found {
		t.Error("tuint8 was not found")
	}
	if x.(uint8) != 3 {
		t.Error("tuint8 is not 3:", x)
	}
}

func TestIncrementWithUint16(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tuint16", uint16(1), DefaultExpiration)
	err = tc.Increment("tuint16", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}

	x, found := tc.Get("tuint16")
	if !found {
		t.Error("tuint16 was not found")
	}
	if x.(uint16) != 3 {
		t.Error("tuint16 is not 3:", x)
	}
}

func TestIncrementWithUint32(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tuint32", uint32(1), DefaultExpiration)
	err = tc.Increment("tuint32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tuint32")
	if !found {
		t.Error("tuint32 was not found")
	}
	if x.(uint32) != 3 {
		t.Error("tuint32 is not 3:", x)
	}
}

func TestIncrementWithUint64(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tuint64", uint64(1), DefaultExpiration)
	err = tc.Increment("tuint64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}

	x, found := tc.Get("tuint64")
	if !found {
		t.Error("tuint64 was not found")
	}
	if x.(uint64) != 3 {
		t.Error("tuint64 is not 3:", x)
	}
}

func TestIncrementWithFloat32(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("float32", float32(1.5), DefaultExpiration)
	err = tc.Increment("float32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3.5 {
		t.Error("float32 is not 3.5:", x)
	}
}

func TestIncrementWithFloat64(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("float64", float64(1.5), DefaultExpiration)
	err = tc.Increment("float64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3.5 {
		t.Error("float64 is not 3.5:", x)
	}
}

func TestIncrementFloatWithFloat32(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("float32", float32(1.5), DefaultExpiration)
	err = tc.IncrementFloat("float32", 2)
	if err != nil {
		t.Error("Error incrementfloating:", err)
	}
	x, found := tc.Get("float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3.5 {
		t.Error("float32 is not 3.5:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	err = tc.IncrementFloat("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	err = tc.IncrementFloat("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestIncrementFloatWithFloat64(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("float64", float64(1.5), DefaultExpiration)
	err = tc.IncrementFloat("float64", 2)
	if err != nil {
		t.Error("Error incrementfloating:", err)
	}
	x, found := tc.Get("float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3.5 {
		t.Error("float64 is not 3.5:", x)
	}
}

func TestDecrementWithInt(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("int", int(5), DefaultExpiration)
	err = tc.Decrement("int", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("int")
	if !found {
		t.Error("int was not found")
	}
	if x.(int) != 3 {
		t.Error("int is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	err = tc.Decrement("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	err = tc.Decrement("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestDecrementWithInt8(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("int8", int8(5), DefaultExpiration)
	err = tc.Decrement("int8", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("int8")
	if !found {
		t.Error("int8 was not found")
	}
	if x.(int8) != 3 {
		t.Error("int8 is not 3:", x)
	}
}

func TestDecrementWithInt16(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("int16", int16(5), DefaultExpiration)
	err = tc.Decrement("int16", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("int16")
	if !found {
		t.Error("int16 was not found")
	}
	if x.(int16) != 3 {
		t.Error("int16 is not 3:", x)
	}
}

func TestDecrementWithInt32(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("int32", int32(5), DefaultExpiration)
	err = tc.Decrement("int32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("int32")
	if !found {
		t.Error("int32 was not found")
	}
	if x.(int32) != 3 {
		t.Error("int32 is not 3:", x)
	}
}

func TestDecrementWithInt64(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("int64", int64(5), DefaultExpiration)
	err = tc.Decrement("int64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("int64")
	if !found {
		t.Error("int64 was not found")
	}
	if x.(int64) != 3 {
		t.Error("int64 is not 3:", x)
	}
}

func TestDecrementWithUint(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("uint", uint(5), DefaultExpiration)
	err = tc.Decrement("uint", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uint")
	if !found {
		t.Error("uint was not found")
	}
	if x.(uint) != 3 {
		t.Error("uint is not 3:", x)
	}
}

func TestDecrementWithUintptr(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("uintptr", uintptr(5), DefaultExpiration)
	err = tc.Decrement("uintptr", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uintptr")
	if !found {
		t.Error("uintptr was not found")
	}
	if x.(uintptr) != 3 {
		t.Error("uintptr is not 3:", x)
	}
}

func TestDecrementWithUint8(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("uint8", uint8(5), DefaultExpiration)
	err = tc.Decrement("uint8", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uint8")
	if !found {
		t.Error("uint8 was not found")
	}
	if x.(uint8) != 3 {
		t.Error("uint8 is not 3:", x)
	}
}

func TestDecrementWithUint16(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("uint16", uint16(5), DefaultExpiration)
	err = tc.Decrement("uint16", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uint16")
	if !found {
		t.Error("uint16 was not found")
	}
	if x.(uint16) != 3 {
		t.Error("uint16 is not 3:", x)
	}
}

func TestDecrementWithUint32(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("uint32", uint32(5), DefaultExpiration)
	err = tc.Decrement("uint32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uint32")
	if !found {
		t.Error("uint32 was not found")
	}
	if x.(uint32) != 3 {
		t.Error("uint32 is not 3:", x)
	}
}

func TestDecrementWithUint64(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("uint64", uint64(5), DefaultExpiration)
	err = tc.Decrement("uint64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uint64")
	if !found {
		t.Error("uint64 was not found")
	}
	if x.(uint64) != 3 {
		t.Error("uint64 is not 3:", x)
	}
}

func TestDecrementWithFloat32(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("float32", float32(5.5), DefaultExpiration)
	err = tc.Decrement("float32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3.5 {
		t.Error("float32 is not 3:", x)
	}
}

func TestDecrementWithFloat64(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("float64", float64(5.5), DefaultExpiration)
	err = tc.Decrement("float64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3.5 {
		t.Error("float64 is not 3:", x)
	}
}

func TestDecrementFloatWithFloat32(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("float32", float32(5.5), DefaultExpiration)
	err = tc.DecrementFloat("float32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3.5 {
		t.Error("float32 is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	err = tc.DecrementFloat("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	err = tc.DecrementFloat("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestDecrementFloatWithFloat64(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("float64", float64(5.5), DefaultExpiration)
	err = tc.DecrementFloat("float64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3.5 {
		t.Error("float64 is not 3:", x)
	}
}

func TestIncrementInt(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tint", 1, DefaultExpiration)
	n, err := tc.IncrementInt("tint", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tint")
	if !found {
		t.Error("tint was not found")
	}
	if x.(int) != 3 {
		t.Error("tint is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.IncrementInt("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.IncrementInt("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestIncrementInt8(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tint8", int8(1), DefaultExpiration)
	n, err := tc.IncrementInt8("tint8", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tint8")
	if !found {
		t.Error("tint8 was not found")
	}
	if x.(int8) != 3 {
		t.Error("tint8 is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.IncrementInt8("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.IncrementInt8("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestIncrementInt16(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tint16", int16(1), DefaultExpiration)
	n, err := tc.IncrementInt16("tint16", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tint16")
	if !found {
		t.Error("tint16 was not found")
	}
	if x.(int16) != 3 {
		t.Error("tint16 is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.IncrementInt16("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.IncrementInt16("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestIncrementInt32(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tint32", int32(1), DefaultExpiration)
	n, err := tc.IncrementInt32("tint32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tint32")
	if !found {
		t.Error("tint32 was not found")
	}
	if x.(int32) != 3 {
		t.Error("tint32 is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.IncrementInt32("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.IncrementInt32("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestIncrementInt64(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tint64", int64(1), DefaultExpiration)
	n, err := tc.IncrementInt64("tint64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tint64")
	if !found {
		t.Error("tint64 was not found")
	}
	if x.(int64) != 3 {
		t.Error("tint64 is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.IncrementInt64("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.IncrementInt64("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestIncrementUint(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tuint", uint(1), DefaultExpiration)
	n, err := tc.IncrementUint("tuint", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tuint")
	if !found {
		t.Error("tuint was not found")
	}
	if x.(uint) != 3 {
		t.Error("tuint is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.IncrementUint("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.IncrementUint("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestIncrementUintptr(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tuintptr", uintptr(1), DefaultExpiration)
	n, err := tc.IncrementUintptr("tuintptr", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tuintptr")
	if !found {
		t.Error("tuintptr was not found")
	}
	if x.(uintptr) != 3 {
		t.Error("tuintptr is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.IncrementUintptr("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.IncrementUintptr("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestIncrementUint8(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tuint8", uint8(1), DefaultExpiration)
	n, err := tc.IncrementUint8("tuint8", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tuint8")
	if !found {
		t.Error("tuint8 was not found")
	}
	if x.(uint8) != 3 {
		t.Error("tuint8 is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.IncrementUint8("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.IncrementUint8("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestIncrementUint16(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tuint16", uint16(1), DefaultExpiration)
	n, err := tc.IncrementUint16("tuint16", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tuint16")
	if !found {
		t.Error("tuint16 was not found")
	}
	if x.(uint16) != 3 {
		t.Error("tuint16 is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.IncrementUint16("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.IncrementUint16("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestIncrementUint32(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tuint32", uint32(1), DefaultExpiration)
	n, err := tc.IncrementUint32("tuint32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tuint32")
	if !found {
		t.Error("tuint32 was not found")
	}
	if x.(uint32) != 3 {
		t.Error("tuint32 is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.IncrementUint32("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.IncrementUint32("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestIncrementUint64(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tuint64", uint64(1), DefaultExpiration)
	n, err := tc.IncrementUint64("tuint64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tuint64")
	if !found {
		t.Error("tuint64 was not found")
	}
	if x.(uint64) != 3 {
		t.Error("tuint64 is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.IncrementUint64("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.IncrementUint64("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestIncrementFloat32(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("float32", float32(1.5), DefaultExpiration)
	n, err := tc.IncrementFloat32("float32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3.5 {
		t.Error("Returned number is not 3.5:", n)
	}
	x, found := tc.Get("float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3.5 {
		t.Error("float32 is not 3.5:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.IncrementFloat32("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.IncrementFloat32("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestIncrementFloat64(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("float64", float64(1.5), DefaultExpiration)
	n, err := tc.IncrementFloat64("float64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3.5 {
		t.Error("Returned number is not 3.5:", n)
	}
	x, found := tc.Get("float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3.5 {
		t.Error("float64 is not 3.5:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.IncrementFloat64("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.IncrementFloat64("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestDecrementInt(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("tint", 5, DefaultExpiration)
	n, err := tc.DecrementInt("tint", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tint")
	if !found {
		t.Error("tint was not found")
	}
	if x.(int) != 3 {
		t.Error("tint is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.DecrementInt("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.DecrementInt("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestDecrementInt8(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("int8", int8(5), DefaultExpiration)
	n, err := tc.DecrementInt8("int8", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("int8")
	if !found {
		t.Error("int8 was not found")
	}
	if x.(int8) != 3 {
		t.Error("int8 is not 3:", x)
	}
	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.DecrementInt8("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.DecrementInt8("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestDecrementInt16(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("int16", int16(5), DefaultExpiration)
	n, err := tc.DecrementInt16("int16", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("int16")
	if !found {
		t.Error("int16 was not found")
	}
	if x.(int16) != 3 {
		t.Error("int16 is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.DecrementInt16("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.DecrementInt16("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestDecrementInt32(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("int32", int32(5), DefaultExpiration)
	n, err := tc.DecrementInt32("int32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("int32")
	if !found {
		t.Error("int32 was not found")
	}
	if x.(int32) != 3 {
		t.Error("int32 is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.DecrementInt32("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.DecrementInt32("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestDecrementInt64(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("int64", int64(5), DefaultExpiration)
	n, err := tc.DecrementInt64("int64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("int64")
	if !found {
		t.Error("int64 was not found")
	}
	if x.(int64) != 3 {
		t.Error("int64 is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.DecrementInt64("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.DecrementInt64("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestDecrementUint(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("uint", uint(5), DefaultExpiration)
	n, err := tc.DecrementUint("uint", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("uint")
	if !found {
		t.Error("uint was not found")
	}
	if x.(uint) != 3 {
		t.Error("uint is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.DecrementUint("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.DecrementUint("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestDecrementUintptr(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("uintptr", uintptr(5), DefaultExpiration)
	n, err := tc.DecrementUintptr("uintptr", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("uintptr")
	if !found {
		t.Error("uintptr was not found")
	}
	if x.(uintptr) != 3 {
		t.Error("uintptr is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.DecrementUintptr("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.DecrementUintptr("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestDecrementUint8(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("uint8", uint8(5), DefaultExpiration)
	n, err := tc.DecrementUint8("uint8", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("uint8")
	if !found {
		t.Error("uint8 was not found")
	}
	if x.(uint8) != 3 {
		t.Error("uint8 is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.DecrementUint8("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.DecrementUint8("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestDecrementUint16(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("uint16", uint16(5), DefaultExpiration)
	n, err := tc.DecrementUint16("uint16", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("uint16")
	if !found {
		t.Error("uint16 was not found")
	}
	if x.(uint16) != 3 {
		t.Error("uint16 is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.DecrementUint16("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.DecrementUint16("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestDecrementUint32(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("uint32", uint32(5), DefaultExpiration)
	n, err := tc.DecrementUint32("uint32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("uint32")
	if !found {
		t.Error("uint32 was not found")
	}
	if x.(uint32) != 3 {
		t.Error("uint32 is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.DecrementUint32("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.DecrementUint32("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestDecrementUint64(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("uint64", uint64(5), DefaultExpiration)
	n, err := tc.DecrementUint64("uint64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("uint64")
	if !found {
		t.Error("uint64 was not found")
	}
	if x.(uint64) != 3 {
		t.Error("uint64 is not 3:", x)
	}

	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.DecrementUint64("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.DecrementUint64("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestDecrementFloat32(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("float32", float32(5), DefaultExpiration)
	n, err := tc.DecrementFloat32("float32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3 {
		t.Error("float32 is not 3:", x)
	}
	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.DecrementFloat32("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.DecrementFloat32("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestDecrementFloat64(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("float64", float64(5), DefaultExpiration)
	n, err := tc.DecrementFloat64("float64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3 {
		t.Error("float64 is not 3:", x)
	}
	tc.Set("tstr", "str", DefaultExpiration)
	_, err = tc.DecrementFloat64("tstr", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
	_, err = tc.DecrementFloat64("not", 2)
	if err == nil {
		t.Error("Not check str:", err)
	}
}

func TestIncrementOverflowInt(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("int8", int8(127), DefaultExpiration)
	err = tc.Increment("int8", 1)
	if err != nil {
		t.Error("Error incrementing int8:", err)
	}
	x, _ := tc.Get("int8")
	int8 := x.(int8)
	if int8 != -128 {
		t.Error("int8 did not overflow as expected; value:", int8)
	}
}

func TestIncrementOverflowUint(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("uint8", uint8(255), DefaultExpiration)
	err = tc.Increment("uint8", 1)
	if err != nil {
		t.Error("Error incrementing int8:", err)
	}
	x, _ := tc.Get("uint8")
	uint8 := x.(uint8)
	if uint8 != 0 {
		t.Error("uint8 did not overflow as expected; value:", uint8)
	}
}

func TestDecrementUnderflowUint(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("uint8", uint8(0), DefaultExpiration)
	err = tc.Decrement("uint8", 1)
	if err != nil {
		t.Error("Error decrementing int8:", err)
	}
	x, _ := tc.Get("uint8")
	uint8 := x.(uint8)
	if uint8 != 255 {
		t.Error("uint8 did not underflow as expected; value:", uint8)
	}
}
