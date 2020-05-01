package soft

import (
	"fmt"
	"strconv"
)

type Soft struct {
	CompareCount int
	SwapCount    int
	Array        []int
}

func NewSoft(array []int) *Soft {
	return &Soft{Array: array}
}

func (s *Soft) Compare(index1, index2 int) int {
	s.CompareCount++
	return s.Array[index1] - s.Array[index2]
}

func (s *Soft) CompareNum(num1, num2 int) int {
	s.CompareCount++
	return num1 - num2
}

func (s *Soft) Swap(index1, index2 int) {
	s.SwapCount++
	if index2==index1{
		return
	}
	s.Array[index1] ^= s.Array[index2]
	s.Array[index2] ^= s.Array[index1]
	s.Array[index1] ^= s.Array[index2]
}

func (s *Soft) String(str, time string) {
	fmt.Println("struct: " + str + " use " + time + "  compare count " +
		strconv.Itoa(s.CompareCount) + " swap count " + strconv.Itoa(s.SwapCount))
}
