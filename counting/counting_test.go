package counting

import (
	"fmt"
	"sort"
	"testing"

	"github.com/algorithmsRepeat/util"
)

func TestCounting_SortTest(t *testing.T) {
	arr := util.NewSlice(100000, 1000000)
	c := NewCounting(arr)
	c.SortTest()
	if sort.IntsAreSorted(c.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	fmt.Println(c.Array)
}
func TestCounting1_SortTest(t *testing.T) {
	arr := util.NewSlice(100000, 1000000)
	c := NewCounting1(arr)
	c.SortTest()
	if sort.IntsAreSorted(c.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	fmt.Println(c.Array)
}
