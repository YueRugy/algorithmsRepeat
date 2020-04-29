package selection

import (
	"fmt"
	"github.com/algorithmsRepeat/util"
	"sort"
	"testing"
)

func TestSelection_SortTest(t *testing.T) {
	arr := util.NewSlice(10000, 100000)
	s := NewSelection(arr)
	s.SortTest()
	if sort.IntsAreSorted(s.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	fmt.Println(s.Array)
}
