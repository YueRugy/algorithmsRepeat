package unionfind

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const max = 50000

func TestUnionFind_F_Union(t *testing.T) {
	uf := NewUnionFind_F(max)
	foo(uf)
}

func TestUnionFind_U_Union(t *testing.T) {
	uf := NewUnionFind_U(max)
	foo(uf)
}
func TestUnionFind_US_Union(t *testing.T) {
	uf := NewUnionFind_US(max)
	foo(uf)
}

func TestUnionFind_UR_Union(t *testing.T) {
	uf := NewUnionFind_UR(max)
	foo(uf)
}

func TestUnionFind_UR_PC_Union(t *testing.T) {
	uf := NewUnionFind_UR_PC(max)
	foo(uf)
}

func TestNewUnionFind_UR_PS_Union(t *testing.T) {
	uf := NewUnionFind_UR_PS(max)
	foo(uf)
}

func TestUnionFind_UR_PH_Union(t *testing.T) {
	uf := NewUnionFind_UR_PH(max)
	foo(uf)
}

func foo(uf Union) {
	rand.Seed(time.Now().Unix())
	uf.Union(0, 1)
	uf.Union(0, 3)
	uf.Union(0, 4)
	uf.Union(2, 3)
	uf.Union(2, 5)
	uf.Union(6, 7)
	uf.Union(8, 10)
	uf.Union(9, 10)
	uf.Union(9, 11)

	fmt.Println(uf.IsSame(2, 7))
	uf.Union(4, 6)
	fmt.Println(uf.IsSame(2, 7))
	begin := time.Now()
	for index := 0; index < max; index++ {
		uf.Union(rand.Intn(max), rand.Intn(max))
	}
	for index := 0; index < max; index++ {
		uf.IsSame(rand.Intn(max), rand.Intn(max))
	}

	str := time.Now().Sub(begin).String()
	fmt.Println("use " + str)
}
