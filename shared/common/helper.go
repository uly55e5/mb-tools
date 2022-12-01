package common

import (
	"encoding/json"
	"golang.org/x/exp/constraints"
	"math"
)

type jsonFloat64 float64

func (f jsonFloat64) MarshalJSON() ([]byte, error) {
	if math.IsNaN(float64(f)) {
		return []byte("\"NaN\""), nil
	} else if math.IsInf(float64(f), +1) {
		return []byte("\"+Inf\""), nil
	} else if math.IsInf(float64(f), -1) {
		return []byte("\"+Inf\""), nil
	}
	return json.Marshal(float64(f))
}

func Max[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m < v {
			m = v
		}
	}
	return m
}

func Min[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m > v {
			m = v
		}
	}
	return m
}

func Unique[T comparable](array []T) []T {
	var uniqeVals []T
	m := map[T]bool{}
	for _, v := range array {
		if !m[v] {
			m[v] = true
			uniqeVals = append(uniqeVals, v)
		}
	}
	return uniqeVals
}
