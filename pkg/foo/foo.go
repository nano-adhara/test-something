package foo

import (
	"github.com/nano-adhara/test-commons/pkg/other"
	"github.com/nano-adhara/test-commons/pkg/some"
)

func Foo(some some.Some) {
	other.Bar(some)
}
