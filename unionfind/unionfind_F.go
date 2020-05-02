package unionfind

type UnionFind_F struct {
	parents []int
}

func NewUnionFind_F(capacity int) *UnionFind_F {
	uf := &UnionFind_F{
		parents: make([]int, capacity),
	}

	for index := 0; index < len(uf.parents); index++ {
		uf.parents[index] = index
	}
	return uf
}

func (uf *UnionFind_F) IsSame(v1, v2 int) bool {
	if uf.checkV(v1) || uf.checkV(v2) || v1 == v2 {
		return false
	}
	return uf.Find(v1) == uf.Find(v2)
}

func (uf *UnionFind_F) checkV(v int) bool {
	if v >= len(uf.parents) || v < 0 {
		return true
	}
	return false
}

func (uf *UnionFind_F) Find(v int) int {
	if uf.checkV(v) {
		return -1
	}

	return uf.parents[v]
}

func (uf *UnionFind_F) Union(v1, v2 int) {
	if uf.checkV(v1) || uf.checkV(v2) || v1 == v2 {
		return
	}

	p1 := uf.Find(v1)
	p2 := uf.Find(v2)
	if p1 == p2 {
		return
	}

	for index := 0; index < len(uf.parents); index++ {
		if uf.parents[index] == p1 {
			uf.parents[index] = p2
		}
	}
}
