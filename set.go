package gf

type Set map[any]struct{}

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(values ...any) {
	for i := range values {
		s[values[i]] = struct{}{}
	}
}

func (s Set) Has(value any) bool {
	_, ok := s[value]
	return ok
}
