package merge

import (
	"time"

	"github.com/algorithmsRepeat/soft"
)

type Merge struct {
	*soft.Soft
	leftArray []int
}

func NewMerge(arr []int) *Merge {
	m := &Merge{
		Soft: soft.NewSoft(arr),
	}
	m.leftArray = make([]int, len(arr)>>1+1)
	return m

}
func (m *Merge) SortTest() {
	before := time.Now()
	m.Sort(0, len(m.Array))
	end := time.Now()
	timeStr := end.Sub(before).String()
	m.String("merge sort", timeStr)
}

func (m *Merge) Sort(begin, end int) {
	if end-begin < 2 {
		return
	}
	mid := (begin + end) >> 1
	m.Sort(begin, mid)
	m.Sort(mid, end)
	m.merge(begin, mid, end)
}

func (m *Merge) merge(begin, mid, end int) {
	ls, le := 0, mid-begin
	rs, re := mid, end
	ai := begin
	for ls < le {
		if rs >= re || m.CompareNum(m.leftArray[ls], m.Array[rs]) <= 0 {
			m.Array[ai] = m.leftArray[ls]
			ls++
			ai++
		} else {
			m.Array[ai] = m.Array[rs]
			rs++
			ai++
		}
	}
}
