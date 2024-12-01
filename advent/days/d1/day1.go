package d1

import (
	"advent/register"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type Day1 struct {
	TotalDistance   int `json:"total_distance"`
	SimilarityScore int `json:"similarity_score"`
}

func init() {
	register.Register(&Day1{})
}

func (d *Day1) Run(ctx context.Context) (string, error) {

	bytes, _ := os.ReadFile("days/d1/INPUT")
	input := string(bytes)

	p1 := &PartOneQuestion{InputData: input}
	p1.Solve()
	d.TotalDistance = p1.PartOneResult.TotalDistance

	p2 := &PartTwoQuestion{InputData: input, PartOneResult: p1.PartOneResult}
	p2.Solve()
	d.SimilarityScore = p2.PartTwoResult.TotalSimilarity

	out, _ := json.Marshal(d)
	return string(out), nil
}

type PartOneQuestion struct {
	InputData     string
	PartOneResult PartOneResult
}

type PartTwoQuestion struct {
	InputData     string
	PartOneResult PartOneResult
	PartTwoResult PartTwoResult
}

type Lists struct {
	LeftList  []int
	RightList []int
}

func (l *Lists) Sort() {
	sort.Ints(l.LeftList)
	sort.Ints(l.RightList)
}

type PartOneResult struct {
	TotalDistance int
	LeftList      []int
	RightList     []int
}

type PartTwoResult struct {
	TotalSimilarity int
}

func (p *PartOneQuestion) Solve() {

	var lists = Lists{
		LeftList:  []int{},
		RightList: []int{},
	}

	for _, line := range strings.Split(p.InputData, "\n") {
		var left, right int
		_, _ = fmt.Sscanf(line, "%d %d", &left, &right)
		lists.LeftList = append(lists.LeftList, left)
		lists.RightList = append(lists.RightList, right)
	}

	lists.Sort()

	var totalDistance int = 0

	for i := 0; i < len(lists.LeftList); i++ {
		totalDistance += int(math.Abs(float64(lists.LeftList[i] - lists.RightList[i])))
	}

	p.PartOneResult = PartOneResult{
		TotalDistance: totalDistance,
		LeftList:      lists.LeftList,
		RightList:     lists.RightList,
	}
}

func (p *PartTwoQuestion) Solve() {

	var totalSimilarity int = 0

	for i := 0; i < len(p.PartOneResult.LeftList); i++ {

		occurrenceInRightList := 0

		for j := 0; j < len(p.PartOneResult.RightList); j++ {
			if p.PartOneResult.LeftList[i] == p.PartOneResult.RightList[j] {
				occurrenceInRightList++
			}
		}

		totalSimilarity += p.PartOneResult.LeftList[i] * occurrenceInRightList
	}

	p.PartTwoResult = PartTwoResult{
		TotalSimilarity: totalSimilarity,
	}
}
