package main

import (
	_ "advent/days/d1"
	_ "advent/days/d2"
	"advent/register"
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {

	if len(os.Args) > 1 {
		for _, day := range register.Registry {
			name := reflect.TypeOf(day).String()
			if strings.HasSuffix(name, os.Args[1]) {
				out, _ := day.Run(context.Background())
				fmt.Println(out)
				break
			}
		}
	} else {
		for _, day := range register.Registry {
			out, _ := day.Run(context.Background())
			fmt.Println(out)
		}
	}
}
