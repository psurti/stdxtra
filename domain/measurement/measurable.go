package measurement

import (
	"fmt"
)

type Measurable interface {
	fmt.Stringer
	Equal(other Measurable) bool
}
