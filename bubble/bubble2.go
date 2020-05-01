package bubble

import (
	"github.com/algorithmsRepeat/soft"
	"time"
)

type Bubble2 struct {
	*soft.Soft
}

func NewBubble2(arr []int) *Bubble2 {
	return &Bubble2{
		Soft: soft.NewSoft(arr),
	}
}
func (b *Bubble2) SortTest() {
	before := time.Now()
	b.Sort()
	end := time.Now()
	timeStr := end.Sub(before).String()
	b.String("bubble2 sort", timeStr)
}
func (b *Bubble2) Sort() {
	for end := len(b.Array) - 1; end > 0; {
		lastIndex := 0
		for begin := 0; begin < end; begin++ {
			if b.Compare(begin, begin+1) > 0 {
				b.Swap(begin,begin+1)
				lastIndex = begin
			}
		}
		end = lastIndex
	}

}
