package unionfind

type UnionFind_US struct {
	parents []int
	size    []int
}

func NewUnionFind_US(capacity int) *UnionFind_US {
	uf := &UnionFind_US{
		parents: make([]int, capacity),
	}
	uf.size=make([]int,capacity)
	for index := 0; index < len(uf.parents); index++ {
		uf.parents[index] = index
		uf.size[index] = 1
	}
	return uf
}

func (uf *UnionFind_US) IsSame(v1, v2 int) bool {
	if uf.checkV(v1) || uf.checkV(v2) || v1 == v2 {
		return false
	}
	return uf.Find(v1) == uf.Find(v2)
}

func (uf *UnionFind_US) checkV(v int) bool {
	if v >= len(uf.parents) || v < 0 {
		return true
	}
	return false
}

func (uf *UnionFind_US) Find(v int) int {
	if uf.checkV(v) {
		return -1
	}

	for v != uf.parents[v] {
		v = uf.parents[v]
	}

	return v
}

func (uf *UnionFind_US) Union(v1, v2 int) {
	if uf.checkV(v1) || uf.checkV(v2) || v1 == v2 {
		return
	}

	p1 := uf.Find(v1)
	p2 := uf.Find(v2)
	if p1 == p2 {
		return
	}
	if uf.size[p1] < uf.size[p2] {
		uf.parents[p1] = p2
		uf.size[p2] += uf.size[p1]
	} else {
		uf.parents[p2] = p1
		uf.size[p1] += uf.size[p2]
	}
}
