package pair

type Pair[T, F any] struct {
	First T
	Second F
}

func New[T, F any](first T, second F) Pair[T, F] {
	return Pair[T, F]{
		First: first,
		Second: second,
	}
}