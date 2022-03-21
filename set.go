package gf

type Set[T any] map[any]struct{}

func NewSet[T any]() Set[T] {
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
