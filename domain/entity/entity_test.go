package entity

import (
	"fmt"
	"testing"
)

type customID struct {
	id float64
}

func (c customID) String() string {
	return fmt.Sprintf("%v", c.id)
}

func TestValue(t *testing.T) {
	t.Run("NewWithIntID", func(t *testing.T) {
		New[string](WithIntID[string](100))
	})
	t.Run("NewWithStrID", func(t *testing.T) {
		New[string](WithStrID[string]("s100"))
	})
	t.Run("NewWithID", func(t *testing.T) {
		cid := customID{id: 100.2}
		New[string](WithID[string](cid))
	})
	t.Run("NewInt", func(t *testing.T) {
		New[int]()
	})
}
