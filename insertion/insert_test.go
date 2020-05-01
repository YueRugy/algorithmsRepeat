package insertion

import (
	"fmt"
	"github.com/algorithmsRepeat/util"
	"sort"
	"testing"
)

func TestSelection_SortTest(t *testing.T) {
	arr := util.NewSlice(10000, 100000)
	s := NewInsert(arr)
	s.SortTest()
	if sort.IntsAreSorted(s.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	fmt.Println(s.Array)
}
func TestSelection1_SortTest(t *testing.T) {
	arr := util.NewSlice(10000, 100000)
	s := NewInsert1(arr)
	s.SortTest()
	if sort.IntsAreSorted(s.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
func TestSelection2_SortTest(t *testing.T) {
	arr := util.NewSlice(100000, 1000000)
	s := NewInsert2(arr)
	s.SortTest()
	if sort.IntsAreSorted(s.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
