package d2

import (
	"advent/register"
	"context"
)

type Day2 struct {
}

func (d *Day2) Run(ctx context.Context) (string, error) {
	return "Hello, World!", nil
}

func init() {
	register.Register(&Day2{})
}
