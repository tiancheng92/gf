package gf

type Set[T comparable] map[any]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(values ...T) {
	for i := range values {
		s[values[i]] = struct{}{}
	}
}

func (s Set[T]) Has(value T) bool {
	_, ok := s[value]
	return ok
}
