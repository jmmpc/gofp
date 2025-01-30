package stack

import (
	"slices"

	"github.com/jmmpc/gofp/list"
)

type Stack[T any] struct {
	list []T
}

func New[T any]() Stack[T] {
	return Stack[T]{
		list: []T(nil),
	}
}

func IsEmpty[T any](s Stack[T]) bool {
	return len(s.list) == 0
}

func Size[T any](s Stack[T]) int {
	return len(s.list)
}

func Pop[T any](s Stack[T]) (T, Stack[T]) {
	if len(s.list) == 0 {
		return *new(T), s
	}

	size := len(s.list)
	top := s.list[size-1]
	s.list = s.list[:size-1]

	return top, s
}

func Push[T any](s Stack[T], v T) Stack[T] {
	s.list = append(s.list, v)
	return s
}

func Top[T any](s Stack[T]) T {
	if len(s.list) == 0 {
		return *new(T)
	}

	return s.list[len(s.list)-1]
}

func Contains[T comparable](s Stack[T], v T) bool {
	return slices.Contains(s.list, v)
}

func ContainsFunc[T any](s Stack[T], f func(T) bool) bool {
	return slices.ContainsFunc(s.list, f)
}

func ToList[T any](s Stack[T]) []T {
	return s.list
}

func FromList[S ~[]T, T any](s S) Stack[T] {
	return Stack[T]{list: s}
}

func Map[T1, T2 any](s Stack[T1], f func(T1) T2) Stack[T2] {
	return Stack[T2]{list: list.Map(s.list, f)}
}

func Filter[T any](s Stack[T], f func(T) bool) Stack[T] {
	return Stack[T]{list: list.Filter(s.list, f)}
}
