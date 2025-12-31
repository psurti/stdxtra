package entity

import (
	"fmt"
)

// Entity[K, V] is generic over an ID type K and a payload/value type V.
type Entity[K any, V any] struct {
	id K
	val V
}

// Option is a functional option for Entity[K, V].
type Option[K any, V any] func(*Entity[K, V])

// New constructs an Entity[K, V] and applies provided options.
func New[K any, V any](opts ...Option[K, V]) Entity[K, V] {
	e := &Entity[K, V]{}
	for _, o := range opts {
		if o == nil {
			continue
		}
		o(e)
	}
	return *e
}

// WithID sets the entity ID (typed to the concrete ID type K).
func WithID[K any, V any](id K) Option[K, V] {
	return func(e *Entity[K, V]) {
		e.id = id
	}
}

// WithValue sets the entity payload/value (typed to V).
func WithValue[K any, V any](v V) Option[K, V] {
	return func(e *Entity[K, V]) {
		e.val = v
	}
}

// ID returns the stored ID (zero value if unset).
func (e Entity[K, V]) ID() K {
	return e.id
}

// Value returns the stored payload/value.
func (e Entity[K, V]) Value() V {
	return e.val
}

// StrID returns the string representation of the ID or an empty string when the
// underlying ID is nil (for pointer/interface types). If the ID implements
// fmt.Stringer its String() will be used; otherwise fmt.Sprintf("%v", id).
//
// Note: for primitive zero-values (e.g. int == 0 or string == ""), this will
// return the formatted representation ("0" or "") respectively.
func (e Entity[K, V]) StrID() string {
	if any(e.id) == nil {
		return ""
	}
	if s, ok := any(e.id).(fmt.Stringer); ok {
		return s.String()
	}
	return fmt.Sprintf("%v", any(e.id))
}

// -- Convenience helpers and aliases to reduce repetition at call-sites --

// Option aliases for common primitive ID types.
type OptionString[V any] = Option[string, V]
type OptionInt[V any] = Option[int, V]

// NewWithStringID constructs an Entity whose ID type is string and only requires V.
func NewWithStringID[V any](opts ...OptionString[V]) Entity[string, V] {
	return New[string, V](opts...)
}

// NewWithIntID constructs an Entity whose ID type is int and only requires V.
func NewWithIntID[V any](opts ...OptionInt[V]) Entity[int, V] {
	return New[int, V](opts...)
}

// WithStringID returns an Option for entities whose ID type is string.
func WithStringID[V any](s string) OptionString[V] {
	return func(e *Entity[string, V]) {
		e.id = s
	}
}

// WithIntID returns an Option for entities whose ID type is int.
func WithIntID[V any](i int) OptionInt[V] {
	return func(e *Entity[int, V]) {
		e.id = i
	}
}

// WithValueForStringID is a convenience setter for the value when K == string.
func WithValueForStringID[V any](v V) OptionString[V] {
	return WithValue[string, V](v)
}

// WithValueForIntID is a convenience setter for the value when K == int.
func WithValueForIntID[V any](v V) OptionInt[V] {
	return WithValue[int, V](v)
}

// -- One-shot constructors --

// NewStringIDEntity creates and returns an Entity[string, V] initialized with the given id and value.
func NewStringIDEntity[V any](id string, v V) Entity[string, V] {
	return Entity[string, V]{id: id, val: v}
}

// NewIntIDEntity creates and returns an Entity[int, V] initialized with the given id and value.
func NewIntIDEntity[V any](id int, v V) Entity[int, V] {
	return Entity[int, V]{id: id, val: v}
}

// NewEntityWithID creates and returns an Entity[K, V] initialized with the given id and value.
func NewEntityWithID[K any, V any](id K, v V) Entity[K, V] {
	return Entity[K, V]{id: id, val: v}
}
