package unionfind

type UnionFind_UR struct {
	parents []int
	rank    []int
}

func NewUnionFind_UR(capacity int) *UnionFind_UR {
	uf := &UnionFind_UR{
		parents: make([]int, capacity),
	}
	uf.rank=make([]int,capacity)
	for index := 0; index < len(uf.parents); index++ {
		uf.parents[index] = index
		uf.rank[index] = 1
	}
	return uf
}

func (uf *UnionFind_UR) IsSame(v1, v2 int) bool {
	if uf.checkV(v1) || uf.checkV(v2) || v1 == v2 {
		return false
	}
	return uf.Find(v1) == uf.Find(v2)
}

func (uf *UnionFind_UR) checkV(v int) bool {
	if v >= len(uf.parents) || v < 0 {
		return true
	}
	return false
}

func (uf *UnionFind_UR) Find(v int) int {
	if uf.checkV(v) {
		return -1
	}

	for v != uf.parents[v] {
		v = uf.parents[v]
	}

	return v
}

func (uf *UnionFind_UR) Union(v1, v2 int) {
	if uf.checkV(v1) || uf.checkV(v2) || v1 == v2 {
		return
	}

	p1 := uf.Find(v1)
	p2 := uf.Find(v2)
	if p1 == p2 {
		return
	}

	if uf.rank[p1] > uf.rank[p2] {
		uf.parents[p2] = p1
	} else if uf.rank[p1] < uf.rank[p2] {
		uf.parents[p1] = p2
	} else {
		uf.parents[p1] = p2
		uf.rank[p2]++
	}
}
