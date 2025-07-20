package entity

type Entity[T any] struct {
	val T
	id  ID
}

func New[T any](options ...func(Entity[T])) Entity[T] {
	e := Entity[T]{}
	for _, option := range options {
		option(e)
	}
	return e
}

func WithID[T any](id ID) func(Entity[T]) {
	return func(e Entity[T]) {
		e.id = id
	}
}

func WithIntID[T any](id int) func(Entity[T]) {
	return WithID[T](NewIntID(id))
}

func WithStrID[T any](id string) func(Entity[T]) {
	return WithID[T](NewStrID(id))
}

func (e Entity[T]) ID() ID {
	return e.id
}

func (e Entity[T]) StrID() string {
	return e.id.String()
}
