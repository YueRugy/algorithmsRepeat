package unionfind

type UnionFind_U struct {
	parents []int
}

func NewUnionFind_U(capacity int) *UnionFind_U {
	uf := &UnionFind_U{
		parents: make([]int, capacity),
	}

	for index := 0; index < len(uf.parents); index++ {
		uf.parents[index] = index
	}
	return uf
}

func (uf *UnionFind_U) IsSame(v1, v2 int) bool {
	if uf.checkV(v1) || uf.checkV(v2) || v1 == v2 {
		return false
	}
	return uf.Find(v1) == uf.Find(v2)
}

func (uf *UnionFind_U) checkV(v int) bool {
	if v >= len(uf.parents) || v < 0 {
		return true
	}
	return false
}

func (uf *UnionFind_U) Find(v int) int {
	if uf.checkV(v) {
		return -1
	}

	for v != uf.parents[v] {
		v = uf.parents[v]
	}

	return v
}

func (uf *UnionFind_U) Union(v1, v2 int) {
	if uf.checkV(v1) || uf.checkV(v2) || v1 == v2 {
		return
	}

	p1 := uf.Find(v1)
	p2 := uf.Find(v2)
	if p1 == p2 {
		return
	}

	uf.parents[p1] = p2
}
