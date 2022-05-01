// nolint
package rebis

import (
	"fmt"
	"time"
)

/*
	Expired returns true if the item has expired.
*/
func (item Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}

	return time.Now().UnixNano() > item.Expiration
}

/*
	Increment an item of type int, int8, int16, int32, int64, uintptr, uint,
	uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
	item's value is not an integer, if it was not found, or if it is not
	possible to increment it by n. To retrieve the incremented value, use one
	of the specialized methods, e.g. IncrementInt64.
*/
func (c *cache) Increment(k string, n int64) error {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return fmt.Errorf("item %s not found", k)
	}
	switch v.Value.(type) {
	case int:
		v.Value = v.Value.(int) + int(n)
	case int8:
		v.Value = v.Value.(int8) + int8(n)
	case int16:
		v.Value = v.Value.(int16) + int16(n)
	case int32:
		v.Value = v.Value.(int32) + int32(n)
	case int64:
		v.Value = v.Value.(int64) + int64(n)
	case uint:
		v.Value = v.Value.(uint) + uint(n)
	case uintptr:
		v.Value = v.Value.(uintptr) + uintptr(n)
	case uint8:
		v.Value = v.Value.(uint8) + uint8(n)
	case uint16:
		v.Value = v.Value.(uint16) + uint16(n)
	case uint32:
		v.Value = v.Value.(uint32) + uint32(n)
	case uint64:
		v.Value = v.Value.(uint64) + uint64(n)
	case float32:
		v.Value = v.Value.(float32) + float32(n)
	case float64:
		v.Value = v.Value.(float64) + float64(n)
	default:
		c.mu.Unlock()
		return fmt.Errorf("the value for %s is not an integer", k)
	}
	c.items[k] = v
	c.mu.Unlock()
	return nil
}

/*
	Increment an item of type float32 or float64 by n. Returns an error if the
 	item's value is not floating point, if it was not found, or if it is not
 	possible to increment it by n. Pass a negative number to decrement the
 	value. To retrieve the incremented value, use one of the specialized methods,
 	e.g. IncrementFloat64.
*/
func (c *cache) IncrementFloat(k string, n float64) error {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return fmt.Errorf("item %s not found", k)
	}
	switch v.Value.(type) {
	case float32:
		v.Value = v.Value.(float32) + float32(n)
	case float64:
		v.Value = v.Value.(float64) + float64(n)
	default:
		c.mu.Unlock()
		return fmt.Errorf("the value for %s does not have type float32 or float64", k)
	}
	c.items[k] = v
	c.mu.Unlock()
	return nil
}

/*
	Increment an item of type int by n. Returns an error if the item's value is
	not an int, or if it was not found. If there is no error, the incremented
	value is returned.
*/
func (c *cache) IncrementInt(k string, n int) (int, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(int)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an int", k)
	}
	nv := rv + n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Increment an item of type int8 by n. Returns an error if the item's value is
	not an int8, or if it was not found. If there is no error, the incremented
	value is returned.
*/
func (c *cache) IncrementInt8(k string, n int8) (int8, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(int8)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an int8", k)
	}
	nv := rv + n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Increment an item of type int16 by n. Returns an error if the item's value is
	not an int16, or if it was not found. If there is no error, the incremented
	value is returned.
*/
func (c *cache) IncrementInt16(k string, n int16) (int16, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(int16)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an int16", k)
	}
	nv := rv + n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Increment an item of type int32 by n. Returns an error if the item's value is
	not an int32, or if it was not found. If there is no error, the incremented
	value is returned.
*/
func (c *cache) IncrementInt32(k string, n int32) (int32, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(int32)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an int32", k)
	}
	nv := rv + n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Increment an item of type int64 by n. Returns an error if the item's value is
	not an int64, or if it was not found. If there is no error, the incremented
	value is returned.
*/
func (c *cache) IncrementInt64(k string, n int64) (int64, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(int64)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an int64", k)
	}
	nv := rv + n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Increment an item of type uint by n. Returns an error if the item's value is
	not an uint, or if it was not found. If there is no error, the incremented
	value is returned.
*/
func (c *cache) IncrementUint(k string, n uint) (uint, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(uint)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an uint", k)
	}
	nv := rv + n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Increment an item of type uintptr by n. Returns an error if the item's value is
	not an uintptr, or if it was not found. If there is no error, the incremented
	value is returned.
*/
func (c *cache) IncrementUintptr(k string, n uintptr) (uintptr, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(uintptr)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an uintptr", k)
	}
	nv := rv + n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Increment an item of type uint8 by n. Returns an error if the item's value is
	not an uint8, or if it was not found. If there is no error, the incremented
	value is returned.
*/
func (c *cache) IncrementUint8(k string, n uint8) (uint8, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(uint8)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an uint8", k)
	}
	nv := rv + n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Increment an item of type uint16 by n. Returns an error if the item's value is
	not an uint16, or if it was not found. If there is no error, the incremented
	value is returned.
*/
func (c *cache) IncrementUint16(k string, n uint16) (uint16, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(uint16)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an uint16", k)
	}
	nv := rv + n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Increment an item of type uint32 by n. Returns an error if the item's value is
	not an uint32, or if it was not found. If there is no error, the incremented
	value is returned.
*/
func (c *cache) IncrementUint32(k string, n uint32) (uint32, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(uint32)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an uint32", k)
	}
	nv := rv + n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Increment an item of type uint64 by n. Returns an error if the item's value is
	not an uint64, or if it was not found. If there is no error, the incremented
	value is returned.
*/
func (c *cache) IncrementUint64(k string, n uint64) (uint64, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(uint64)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an uint64", k)
	}
	nv := rv + n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Increment an item of type float32 by n. Returns an error if the item's value is
	not an float32, or if it was not found. If there is no error, the incremented
	value is returned.
*/
func (c *cache) IncrementFloat32(k string, n float32) (float32, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(float32)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an float32", k)
	}
	nv := rv + n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Increment an item of type float64 by n. Returns an error if the item's value is
	not an float64, or if it was not found. If there is no error, the incremented
	value is returned.
*/
func (c *cache) IncrementFloat64(k string, n float64) (float64, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(float64)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an float64", k)
	}
	nv := rv + n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Decrement an item of type int, int8, int16, int32, int64, uintptr, uint,
	uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
	item's value is not an integer, if it was not found, or if it is not
	possible to decrement it by n. To retrieve the decremented value, use one
	of the specialized methods, e.g. DecrementInt64.
*/
func (c *cache) Decrement(k string, n int64) error {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return fmt.Errorf("item not found")
	}
	switch v.Value.(type) {
	case int:
		v.Value = v.Value.(int) - int(n)
	case int8:
		v.Value = v.Value.(int8) - int8(n)
	case int16:
		v.Value = v.Value.(int16) - int16(n)
	case int32:
		v.Value = v.Value.(int32) - int32(n)
	case int64:
		v.Value = v.Value.(int64) - int64(n)
	case uint:
		v.Value = v.Value.(uint) - uint(n)
	case uintptr:
		v.Value = v.Value.(uintptr) - uintptr(n)
	case uint8:
		v.Value = v.Value.(uint8) - uint8(n)
	case uint16:
		v.Value = v.Value.(uint16) - uint16(n)
	case uint32:
		v.Value = v.Value.(uint32) - uint32(n)
	case uint64:
		v.Value = v.Value.(uint64) - uint64(n)
	case float32:
		v.Value = v.Value.(float32) - float32(n)
	case float64:
		v.Value = v.Value.(float64) - float64(n)
	default:
		c.mu.Unlock()
		return fmt.Errorf("the value for %s is not an integer", k)
	}
	c.items[k] = v
	c.mu.Unlock()
	return nil
}

/*
	Decrement an item of type float32 or float64 by n. Returns an error if the
	item's value is not floating point, if it was not found, or if it is not
	possible to decrement it by n. Pass a negative number to decrement the
	value. To retrieve the decremented value, use one of the specialized methods,
	e.g. DecrementFloat64.
*/
func (c *cache) DecrementFloat(k string, n float64) error {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return fmt.Errorf("item %s not found", k)
	}
	switch v.Value.(type) {
	case float32:
		v.Value = v.Value.(float32) - float32(n)
	case float64:
		v.Value = v.Value.(float64) - n
	default:
		c.mu.Unlock()
		return fmt.Errorf("the value for %s does not have type float32 or float64", k)
	}
	c.items[k] = v
	c.mu.Unlock()
	return nil
}

/*
	Decrement an item of type int by n. Returns an error if the item's value is
	not an int, or if it was not found. If there is no error, the decremented
	value is returned.
*/
func (c *cache) DecrementInt(k string, n int) (int, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(int)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an int", k)
	}
	nv := rv - n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Decrement an item of type int8 by n. Returns an error if the item's value is
	not an int8, or if it was not found. If there is no error, the decremented
	value is returned.
*/
func (c *cache) DecrementInt8(k string, n int8) (int8, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(int8)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an int8", k)
	}
	nv := rv - n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Decrement an item of type int16 by n. Returns an error if the item's value is
	not an int16, or if it was not found. If there is no error, the decremented
	value is returned.
*/
func (c *cache) DecrementInt16(k string, n int16) (int16, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(int16)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an int16", k)
	}
	nv := rv - n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Decrement an item of type int32 by n. Returns an error if the item's value is
	not an int32, or if it was not found. If there is no error, the decremented
	value is returned.
*/
func (c *cache) DecrementInt32(k string, n int32) (int32, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(int32)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an int32", k)
	}
	nv := rv - n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Decrement an item of type int64 by n. Returns an error if the item's value is
	not an int64, or if it was not found. If there is no error, the decremented
	value is returned.
*/
func (c *cache) DecrementInt64(k string, n int64) (int64, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(int64)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an int64", k)
	}
	nv := rv - n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Decrement an item of type uint by n. Returns an error if the item's value is
	not an uint, or if it was not found. If there is no error, the decremented
	value is returned.
*/
func (c *cache) DecrementUint(k string, n uint) (uint, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(uint)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an uint", k)
	}
	nv := rv - n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Decrement an item of type uintptr by n. Returns an error if the item's value is
	not an uintptr, or if it was not found. If there is no error, the decremented
	value is returned.
*/
func (c *cache) DecrementUintptr(k string, n uintptr) (uintptr, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(uintptr)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an uintptr", k)
	}
	nv := rv - n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Decrement an item of type uint8 by n. Returns an error if the item's value is
	not an uint8, or if it was not found. If there is no error, the decremented
	value is returned.
*/
func (c *cache) DecrementUint8(k string, n uint8) (uint8, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(uint8)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an uint8", k)
	}
	nv := rv - n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Decrement an item of type uint16 by n. Returns an error if the item's value is
	not an uint16, or if it was not found. If there is no error, the decremented
	value is returned.
*/
func (c *cache) DecrementUint16(k string, n uint16) (uint16, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(uint16)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an uint16", k)
	}
	nv := rv - n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Decrement an item of type uint32 by n. Returns an error if the item's value is
	not an uint32, or if it was not found. If there is no error, the decremented
	value is returned.
*/
func (c *cache) DecrementUint32(k string, n uint32) (uint32, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(uint32)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an uint32", k)
	}
	nv := rv - n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Decrement an item of type uint64 by n. Returns an error if the item's value is
	not an uint64, or if it was not found. If there is no error, the decremented
	value is returned.
*/
func (c *cache) DecrementUint64(k string, n uint64) (uint64, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(uint64)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an uint64", k)
	}
	nv := rv - n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Decrement an item of type float32 by n. Returns an error if the item's value is
	not an float32, or if it was not found. If there is no error, the decremented
	value is returned.
*/
func (c *cache) DecrementFloat32(k string, n float32) (float32, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(float32)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an float32", k)
	}
	nv := rv - n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

/*
	Decrement an item of type float64 by n. Returns an error if the item's value is
	not an float64, or if it was not found. If there is no error, the decremented
	value is returned.
*/
func (c *cache) DecrementFloat64(k string, n float64) (float64, error) {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return 0, fmt.Errorf("item %s not found", k)
	}
	rv, ok := v.Value.(float64)
	if !ok {
		c.mu.Unlock()
		return 0, fmt.Errorf("the value for %s is not an float64", k)
	}
	nv := rv - n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}
