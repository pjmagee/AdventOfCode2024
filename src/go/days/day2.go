package days

import (
	"context"
	"math"
	"strconv"
	"strings"
)

type Day2 struct {
}

func (d *Day2) IsSafe(levels []int) bool {

	increasing := true
	decreasing := true

	for level := 1; level < len(levels); level++ {
		diff := levels[level] - levels[level-1]

		if diff < 0 {
			increasing = false
		}

		if diff > 0 {
			decreasing = false
		}
	}

	withinDiffRange := true

	for level := 1; level < len(levels); level++ {
		diff := int(math.Abs(float64(levels[level] - levels[level-1])))

		if diff < 1 || diff > 3 {
			withinDiffRange = false
			break
		}
	}

	return (increasing || decreasing) && withinDiffRange
}

func (d *Day2) Execute(ctx context.Context, input string) (string, error) {

	reports := strings.Split(input, "\n")
	results := make([]bool, len(reports))

	for reportIdx, report := range reports {
		numbers := strings.Fields(report)
		levels := make([]int, len(numbers))
		for levelIdx, number := range numbers {
			level, _ := strconv.ParseInt(number, 10, 64)
			levels[levelIdx] = int(level)
		}

		result := d.IsSafe(levels)

		if !result {

			// remove a level and try again
			for levelIdx := 0; levelIdx < len(levels); levelIdx++ {
				levelsCopy := make([]int, len(levels))
				copy(levelsCopy, levels)
				levelsCopy = append(levelsCopy[:levelIdx], levelsCopy[levelIdx+1:]...)
				result = d.IsSafe(levelsCopy)
				if result {
					results[reportIdx] = result
				}
			}
		} else {
			results[reportIdx] = result
		}
	}

	safe := 0

	for _, result := range results {
		if result {
			safe++
		}
	}

	return strconv.Itoa(safe), nil
}
