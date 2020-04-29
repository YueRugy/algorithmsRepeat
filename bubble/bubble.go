package bubble

import (
	"github.com/algorithmsRepeat/soft"
	"time"
)

type Bubble struct {
	*soft.Soft
}

func NewBubble(arr []int) *Bubble {
	return &Bubble{
		Soft: soft.NewSoft(arr),
	}
}

func (b *Bubble) SortTest() {
	before := time.Now()
	b.Sort()
	end := time.Now()
	timeStr := end.Sub(before).String()
	b.String("bubble sort", timeStr)
}

func (b *Bubble) Sort() {
	length := len(b.Array)
	for end := length - 1; end > 0; end-- {
		for begin := 0; begin < end; begin++ {
			if b.Compare(begin, begin+1) > 0 {
				b.Swap(begin, begin+1)
			}
		}
	}
}
