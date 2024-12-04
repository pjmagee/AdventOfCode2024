package main

import (
	"advent/days"
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	input := string(data)

	number, _ := strconv.Atoi(os.Args[1])

	for _, day := range days.Days {
		// get name of type and check if ends in day
		structType := fmt.Sprintf("%T", day)

		if strings.HasSuffix(structType, fmt.Sprintf("%d", number)) {
			output, _ := day.Execute(context.Background(), input)
			fmt.Print(output)
		}
	}
}
