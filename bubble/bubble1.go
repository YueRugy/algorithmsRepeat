package bubble

import (
	"github.com/algorithmsRepeat/soft"
	"time"
)

type Bubble1 struct {
	*soft.Soft
}

func NewBubble1(arr []int) *Bubble1 {
	return &Bubble1{
		Soft: soft.NewSoft(arr),
	}
}
func (b *Bubble1) SortTest() {
	before := time.Now()
	b.Sort()
	end := time.Now()
	timeStr := end.Sub(before).String()
	b.String("bubble1 sort", timeStr)
}

func (b *Bubble1) Sort() {

	for end := len(b.Array) - 1; end > 0; end-- {
		flag := true
		for begin := 0; begin < end; begin++ {
			if b.Compare(begin, begin+1) > 0 {
				flag = false
				b.Swap(begin, begin+1)
			}
		}
		if flag {
			break
		}
	}
}
