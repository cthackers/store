package tenant

type Batch[T any] struct {
}

func (c *Collection[T]) NewBatch() *Batch[T] {
	return &Batch[T]{}
}

func (b *Batch[T]) Put(data any) error {
	return nil
}

func (b *Batch[T]) Get(id int) (*Entry[T], error) {
	return nil, nil
}

func (b *Batch[T]) Delete(id int) error {
	return nil
}

func (b *Batch[T]) Update(id int, data any) error {
	return nil
}

func (b *Batch[T]) All() ([]*Entry[T], error) {
	return nil, nil
}

func (b *Batch[T]) Filter(fn func(e *Entry[T]) bool) ([]*Entry[T], error) {
	return nil, nil
}

func (b *Batch[T]) Find(fn func(e *Entry[T]) bool) (*Entry[T], error) {
	return nil, nil
}

func (b *Batch[T]) Commit() error {
	return nil
}

func (b *Batch[T]) Rollback() {
}
