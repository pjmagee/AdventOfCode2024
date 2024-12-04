package days

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day3 struct {
}

type Stack struct {
	chars []rune
}

func (s *Stack) Push(r rune) {
	s.chars = append(s.chars, r)
}

func (s *Stack) Reset() {
	s.chars = nil
}

func (s *Stack) String() string {
	return string(s.chars)
}

func (d *Day3) Execute(ctx context.Context, input string) (string, error) {

	regex := regexp.MustCompile("mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\)")
	matches := regex.FindAllString(input, -1)

	result := 0
	enabled := true
	for _, match := range matches {
		if match == "do()" {
			enabled = true
		} else if match == "don't()" {
			enabled = false
		} else {
			if enabled {
				numbers := make([]int, 2)
				for i, item := range strings.Split(match[4:len(match)-1], ",") {
					numbers[i], _ = strconv.Atoi(item)
				}

				result += numbers[0] * numbers[1]
			}
		}
	}

	return fmt.Sprintf("%d", result), nil
}
