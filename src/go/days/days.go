package days

import (
	"context"
)

type Day interface {
	Execute(ctx context.Context, input string) (string, error)
}

var Days = []Day{
	&Day1{},
	&Day2{},
	&Day3{},
	&Day4{},
}
