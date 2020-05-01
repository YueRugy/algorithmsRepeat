package quick

import (
	"math/rand"
	"time"

	"github.com/algorithmsRepeat/soft"
)

type Quick struct {
	*soft.Soft
}

func NewQuick(arr []int) *Quick {
	return &Quick{
		Soft: soft.NewSoft(arr),
	}
}
func (q *Quick) SortTest() {
	rand.Seed(time.Now().Unix())
	before := time.Now()
	q.Sort(0, len(q.Array))
	end := time.Now()
	timeStr := end.Sub(before).String()
	q.String("quick sort", timeStr)
}
func (q *Quick) Sort(begin, end int) {
	if end-begin < 2 {
		return
	}
	pivotIndex := q.pivotIndex(begin, end)
	q.Sort(begin, pivotIndex)
	q.Sort(pivotIndex+1, end)
}
func (q *Quick) pivotIndex(begin, end int) int {
	end--
	randIndex := rand.Intn(end-begin) + begin
	if randIndex != begin {
		q.Swap(begin, randIndex)
	}
	num := q.Array[begin]
	for begin < end {
		for begin < end {
			if q.Array[end] > num {
				end--
			} else {
				q.Array[begin] = q.Array[end]
				begin++
				break
			}
		}

		for begin < end {
			if q.Array[begin] < num {
				begin++
			} else {
				q.Array[end] = q.Array[begin]
				end--
				break
			}
		}

	}
	q.Array[begin] = num
	return begin
}
