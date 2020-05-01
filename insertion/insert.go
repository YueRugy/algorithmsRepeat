package insertion

import (
	"github.com/algorithmsRepeat/soft"
	"time"
)

type Insert struct {
	*soft.Soft
}

func NewInsert(arr []int) *Insert {
	return &Insert{
		Soft: soft.NewSoft(arr),
	}
}
func (i *Insert) SortTest() {
	before := time.Now()
	i.Sort()
	end := time.Now()
	timeStr := end.Sub(before).String()
	i.String("insert sort", timeStr)
}

func (i *Insert) Sort() {
	for begin := 1; begin < len(i.Array); begin++ {
		for cur := begin; cur > 0; cur-- {
			if i.Compare(cur, cur-1) < 0 {
				i.Swap(cur, cur-1)
			} else {
				break
			}
		}
	}
}
