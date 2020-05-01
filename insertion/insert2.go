package insertion

import (
	"github.com/algorithmsRepeat/soft"
	"time"
)

type Insert2 struct {
	*soft.Soft
}

func NewInsert2(arr []int) *Insert2 {
	return &Insert2{
		Soft: soft.NewSoft(arr),
	}
}
func (i *Insert2) SortTest() {
	before := time.Now()
	i.Sort()
	end := time.Now()
	timeStr := end.Sub(before).String()
	i.String("insert2 sort", timeStr)
}

func (i *Insert2) Sort() {
	for begin := 1; begin < len(i.Array); begin++ {
		index := i.index(begin)
		num := i.Array[begin]
		for tempIndex := begin; tempIndex > index; tempIndex-- {
			i.Array[begin] = i.Array[begin-1]
		}
		i.Array[index] = num
	}
}

func (i *Insert2) index(in int) int {// 找到第一个比i.Array[in]大的索引
	begin, end, num := 0, in, i.Array[in]
	for begin < end {
		mid := (begin + end) >> 1
		if i.CompareNum(num, i.Array[mid]) < 0 {
			end = mid
		} else {
			begin = mid + 1
		}
	}
	return begin
}
