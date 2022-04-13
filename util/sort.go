package util

import "sort"

// genericSlice implements sort.Interface for a slice of T.
type genericSlice[T any] struct {
	items []T
	less  func(T, T) bool
}

func (s genericSlice[T]) Len() int {
	return len(s.items)
}

func (s genericSlice[T]) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

func (s genericSlice[T]) Less(i, j int) bool {
	return s.less(s.items[i], s.items[j])
}

// GenericSort sorts s in place using a comparison function.
func GenericSort[T any](s []T, less func(T, T) bool) {
	sort.Sort(genericSlice[T]{s, less})
}
