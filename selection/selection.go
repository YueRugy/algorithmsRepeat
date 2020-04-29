package selection

import (
	"time"

	"github.com/algorithmsRepeat/soft"
)

type Selection struct {
	*soft.Soft
}

func NewSelection(arr []int) *Selection {
	return &Selection{
		Soft: soft.NewSoft(arr),
	}
}
func (s *Selection) SortTest() {
	before := time.Now()
	s.Sort()
	end := time.Now()
	timeStr := end.Sub(before).String()
	s.String("bubble1 sort", timeStr)
}

func (s *Selection) Sort() {

	for end := len(s.Array) - 1; end > 0; end-- {
		maxIndex := end
		for begin := 0; begin < end; begin++ {
			if s.Compare(begin, maxIndex) > 0 {
				maxIndex = begin
			}
		}
		if maxIndex != end {
			s.Swap(maxIndex, end)
		}
	}
}
