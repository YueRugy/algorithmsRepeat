package quick

import (
	"fmt"
	"sort"
	"testing"

	"github.com/algorithmsRepeat/util"
)

func TestQuick_SortTest(t *testing.T) {
	arr := util.NewSlice(100000, 1000000)
	s := NewQuick(arr)
	s.SortTest()
	if sort.IntsAreSorted(s.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	fmt.Println(s.Array)
}
