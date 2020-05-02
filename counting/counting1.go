package counting

import (
	"time"

	"github.com/algorithmsRepeat/soft"
)

type Counting1 struct {
	*soft.Soft
	countArray []int
}

func NewCounting1(arr []int) *Counting1 {
	c := &Counting1{
		Soft: soft.NewSoft(arr),
	}
	return c

}

func (c *Counting1) SortTest() {
	before := time.Now()
	c.Sort()
	end := time.Now()
	timeStr := end.Sub(before).String()
	c.String("counting sort", timeStr)
}

func (c *Counting1) Sort() {
	min, max := c.Array[0], c.Array[0]
	//find min max index
	for index := 1; index < len(c.Array); index++ {
		if c.CompareNum(max, c.Array[index]) < 0 {
			max = c.Array[index]
		} else if c.CompareNum(min, c.Array[index]) < 0 {
			min = c.Array[index]
		}
	}

	c.countArray = make([]int, max-min+1)
	for index := 0; index < len(c.Array); index++ {
		c.countArray[c.Array[index]]++
	}

	for index := 1; index < len(c.countArray); index++ {
		c.countArray[index] += c.countArray[index-1]
	}

	newArray := make([]int, len(c.Array))
	for index := len(c.Array) - 1; index >= 0; index-- {
		newIndex := c.countArray[c.Array[index]] - 1
		newArray[newIndex] = c.Array[index]
		c.countArray[c.Array[index]]--
	}
}
