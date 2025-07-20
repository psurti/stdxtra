package entity

import "strconv"

type ID interface {
	String() string
}

type StrID struct {
	id string
}

func NewStrID(val string) StrID {
	return StrID{id: val}
}

func (s StrID) String() string {
	return s.id
}

type IntID struct {
	id int
}

func NewIntID(val int) IntID {
	return IntID{id: val}
}
func (i IntID) String() string {
	return strconv.Itoa(i.id)
}
