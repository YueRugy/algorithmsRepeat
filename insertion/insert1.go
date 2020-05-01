package insertion

import (
	"github.com/algorithmsRepeat/soft"
	"time"
)

type Insert1 struct {
	*soft.Soft
}

func NewInsert1(arr []int) *Insert1 {
	return &Insert1{
		Soft: soft.NewSoft(arr),
	}
}
func (i *Insert1) SortTest() {
	before := time.Now()
	i.Sort()
	end := time.Now()
	timeStr := end.Sub(before).String()
	i.String("insert1 sort", timeStr)
}

func (i *Insert1) Sort() {
	for begin := 1; begin < len(i.Array); begin++ {
		cur := begin
		num := i.Array[cur]
		for ; cur > 0; cur-- {
			if i.CompareNum(num, i.Array[cur-1]) < 0 {
				i.Array[cur] = i.Array[cur-1]
			} else {
				break
			}
		}
		i.Array[cur] = num
	}
}
