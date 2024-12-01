package register

import (
	"context"
)

type Day interface {
	Run(ctx context.Context) (string, error)
}

var Registry []Day

func Register(instance Day) {
	Registry = append(Registry, instance)
}
