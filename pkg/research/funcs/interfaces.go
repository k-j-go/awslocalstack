package funcs

type InterfacesMapper[T any] func(value T) T
type InterfacesFilter[T any] func(value T) bool

type Interfaces[T any] struct {
	values []T
}

func NewFromSlice[T any](values []T) *Interfaces[T] {
	this := new(Interfaces[T])
	this.values = values
	return this
}

func (C *Interfaces[T]) Map(mapper InterfacesMapper[T]) *Interfaces[T] {
	new_ := make([]T, 0, len(C.values))
	for _, v := range C.values {
		new_ = append(new_, mapper(v))
	}

	return &Interfaces[T]{values: new_}
}

func (C *Interfaces[T]) Filter(filter InterfacesFilter[T]) *Interfaces[T] {
	new_ := make([]T, 0, len(C.values))
	for _, v := range C.values {
		if filter(v) {
			new_ = append(new_, v)
		}
	}
	return &Interfaces[T]{values: new_}
}
