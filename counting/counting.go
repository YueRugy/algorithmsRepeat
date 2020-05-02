package counting

import (
	"time"

	"github.com/algorithmsRepeat/soft"
)

type Counting struct {
	*soft.Soft
	countArray []int
}

func NewCounting(arr []int) *Counting {
	c := &Counting{
		Soft: soft.NewSoft(arr),
	}
	return c

}

func (c *Counting) SortTest() {
	before := time.Now()
	c.Sort()
	end := time.Now()
	timeStr := end.Sub(before).String()
	c.String("counting sort", timeStr)
}

func (c *Counting) Sort() {
	//find maxnum's index
	max := c.Array[0]
	for index := 1; index < len(c.Array); index++ {
		if c.CompareNum(max, c.Array[index]) < 0 {
			max = c.Array[index]
		}
	}
	c.countArray = make([]int, max+1)
	for index := 0; index < len(c.Array); index++ {
		c.countArray[c.Array[index]]++
	}
	count := 0
	for index := 0; index < len(c.countArray); index++ {
		for c.countArray[index] > 0 {
			c.countArray[index]--
			c.Array[count] = index
			count++
		}
	}
}
