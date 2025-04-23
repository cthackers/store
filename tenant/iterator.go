package tenant

type Iterator[T any] struct {
}

func (i *Iterator[T]) Next() (*Entry[T], bool) {
	return nil, false
}

func (i *Iterator[T]) Close() {
	return
}
