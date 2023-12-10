//go:build go1.1
// +build go1.1

package simplejson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
)

// Implements the json.Unmarshaler interface.
func (j *Json) UnmarshalJSON(p []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(p))
	dec.UseNumber()
	return dec.Decode(&j.data)
}

// NewFromReader returns a *Json by decoding from an io.Reader
func NewFromReader(r io.Reader) (*Json, error) {
	j := new(Json)
	dec := json.NewDecoder(r)
	dec.UseNumber()
	err := dec.Decode(&j.data)
	return j, err
}

// Float64 coerces into a float64
func (j *Json) Float64() (float64, error) {
	switch j.data.(type) {
	case json.Number:
		return j.data.(json.Number).Float64()
	case float32, float64:
		return reflect.ValueOf(j.data).Float(), nil
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(j.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(j.data).Uint()), nil
	}
	f, err := floatFromString(j.data)
	if err == nil {
		return f, nil
	} else {
		return 0, errors.New("invalid value type")
	}
}

// Int coerces into an int
func (j *Json) Int() (int, error) {
	switch j.data.(type) {
	case json.Number:
		i, err := j.data.(json.Number).Int64()
		return int(i), err
	case float32, float64:
		return int(reflect.ValueOf(j.data).Float()), nil
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(j.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(j.data).Uint()), nil
	}
	i, err := intFromString(j.data)
	if err == nil {
		return i, nil
	} else {
		return 0, errors.New("invalid value type")
	}
}

// Int64 coerces into an int64
func (j *Json) Int64() (int64, error) {
	switch j.data.(type) {
	case json.Number:
		return j.data.(json.Number).Int64()
	case float32, float64:
		return int64(reflect.ValueOf(j.data).Float()), nil
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(j.data).Int(), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(j.data).Uint()), nil
	}
	i, err := int64FromString(j.data)
	if err == nil {
		return i, nil
	} else {
		return 0, errors.New("invalid value type")
	}
}

// Uint64 coerces into an uint64
func (j *Json) Uint64() (uint64, error) {
	switch j.data.(type) {
	case json.Number:
		return strconv.ParseUint(j.data.(json.Number).String(), 10, 64)
	case float32, float64:
		return uint64(reflect.ValueOf(j.data).Float()), nil
	case int, int8, int16, int32, int64:
		return uint64(reflect.ValueOf(j.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(j.data).Uint(), nil
	}
	i, err := uint64FromString(j.data)
	if err == nil {
		return i, nil
	} else {
		return 0, errors.New("invalid value type")
	}
}

func floatFromString(raw interface{}) (float64, error) {
	str, ok := raw.(string)
	if !ok {
		return 0, errors.New(fmt.Sprintf("unable to parse, value not string: %T", raw))
	}
	flt, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return flt, nil
}

func intFromString(raw interface{}) (int, error) {
	str, ok := raw.(string)
	if !ok {
		return 0, errors.New(fmt.Sprintf("unable to parse, value not string: %T", raw))
	}
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func int64FromString(raw interface{}) (int64, error) {
	str, ok := raw.(string)
	if !ok {
		return 0, errors.New(fmt.Sprintf("unable to parse, value not string: %T", raw))
	}
	n, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func uint64FromString(raw interface{}) (uint64, error) {
	str, ok := raw.(string)
	if !ok {
		return 0, errors.New(fmt.Sprintf("unable to parse, value not string: %T", raw))
	}
	n, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		return 0, err
	}
	return n, nil
}
