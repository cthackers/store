package tenant

type Collection[T any] struct {
	tenant   *Tenant
	sequence int
	name     string
}

func GetCollection[T any](tenant *Tenant, name string) *Collection[T] {
	return nil
}

func CreateCollection[T any](tenant *Tenant, name string) (*Collection[T], error) {
	return nil, nil
}

func (c *Collection[T]) Count() int {
	return 0
}

func (c *Collection[T]) Put(data T) error {
	return nil
}

func (c *Collection[T]) Get(id int) (*Entry[T], error) {
	return nil, nil
}

func (c *Collection[T]) Delete(id int) error {
	return nil
}

func (c *Collection[T]) Update(id int, data T) error {
	return nil
}

func (c *Collection[T]) All() (*Iterator[T], error) {
	return nil, nil
}

func (c *Collection[T]) Filter(fn func(e *Entry[T]) bool) ([]*Entry[T], error) {
	return nil, nil
}

func (c *Collection[T]) Find(fn func(e *Entry[T]) bool) (*Entry[T], error) {
	return nil, nil
}

func (c *Collection[T]) Truncate() error {
	return nil
}
