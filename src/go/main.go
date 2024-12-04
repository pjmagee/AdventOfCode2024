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

	number, _ := strconv.Atoi(os.Args[1])

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
		os.Exit(1)
	}

	input := string(data)

	for _, day := range days.Days {
		// get name of type and check if ends in day
		structType := fmt.Sprintf("%T", day)

		if strings.HasSuffix(structType, fmt.Sprintf("%d", number)) {
			output, _ := day.Execute(context.Background(), input)
			fmt.Print(output)
		}
	}
}
