package heap

import (
	"sort"
	"testing"

	"github.com/algorithmsRepeat/util"
)

func TestHeap_SortTest(t *testing.T) {
	arr := util.NewSlice(1000000, 1000000)
	h := NewHeap(arr)
	h.SortTest()
	if sort.IntsAreSorted(h.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	//fmt.Println(h.Array)
}
