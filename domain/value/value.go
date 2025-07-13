package value

type Value[T any] struct {
	val T
}

//go:noinline
func Get[T any](v *Value[T]) T {
	return v.val
}

//go:noinline
func Set[T any](v *Value[T], val T) {
	v.val = val
}

//go:noinline
func New[T any](val T) *Value[T] {
	return &Value[T]{
		val: val,
	}
}

//go:noinline
func (v *Value[T]) Get() T {
	return v.val
}

//go:noinline
func (v *Value[T]) Set(val T) {
	v.val = val
}
