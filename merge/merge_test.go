package merge

import (
	"github.com/algorithmsRepeat/util"
	"sort"
	"testing"
)

func TestMerge_SortTest(t *testing.T) {
	arr := util.NewSlice(100000, 1000000)
	m := NewMerge(arr)
	m.SortTest()
	if sort.IntsAreSorted(m.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
