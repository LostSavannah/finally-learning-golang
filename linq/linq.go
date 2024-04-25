package linq

type Generic[T any] chan T
type MapGeneric[T any, U any] chan T

func Channel[T any](self []T) Generic[T] {
	result := make(Generic[T])
	go func() {
		for _, v := range self {
			result <- v
		}
	}()
	return result
}

func (self Generic[T]) Where(filter func(T) bool) Generic[T] {
	result := make(Generic[T])
	go func() {
		for v := range self {
			if filter(v) {
				result <- v
			}
		}
	}()
	return result
}

func (self MapGeneric[T, U]) Select(mapper func(T) U) Generic[U] {
	result := make(Generic[U])
	go func() {
		for v := range self {
			result <- mapper(v)
		}
	}()
	return result
}

func (self Generic[T]) Aggregate(reducer func(T, T) T, initial T) T {
	for v := range self {
		initial = reducer(v, initial)
	}
	return initial
}

func (self Generic[T]) AddRange(source Generic[T]) {
	go func() {
		for v := range source {
			self <- v
		}
	}()
}
