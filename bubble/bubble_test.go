package bubble

import (
	"fmt"
	"github.com/algorithmsRepeat/util"
	"sort"
	"testing"
)

func TestBubble_SortTest(t *testing.T) {
	arr := util.NewSlice(10000, 100000)
	b := NewBubble(arr)
	b.SortTest()
	if sort.IntsAreSorted(b.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	//fmt.Println(b.Array)
}
func TestBubble1_SortTest(t *testing.T) {
	arr := util.NewSlice(10000, 100000)
	b := NewBubble1(arr)
	b.SortTest()
	if sort.IntsAreSorted(b.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	fmt.Println(b.Array)
}
func TestBubble2_SortTest(t *testing.T) {
	arr := util.NewSlice(10000, 100000)
	b := NewBubble2(arr)
	b.SortTest()
	if sort.IntsAreSorted(b.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	fmt.Println(b.Array)
}
