package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func NewSlice(num, limit int) []int {
	if num <= 0 || limit <= 0 {
		return nil
	}
	sli := make([]int, num)
	for index := 0; index < num; index++ {
		sli[index] = rand.Intn(limit)
	}
	return sli
}
