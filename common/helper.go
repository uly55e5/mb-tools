package common

import "golang.org/x/exp/constraints"

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
