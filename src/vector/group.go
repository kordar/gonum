package vector

import (
	"gonum/src/common"
	"gonum/src/matrix"
	"log"
)

type Group[K common.Key, V common.Value] struct {
	direction common.VectorType
	size      K
	data      []*Vector[K, V]
}

func NewGroup[K common.Key, V common.Value](direction common.VectorType, size K, data []*Vector[K, V]) *Group[K, V] {
	return &Group[K, V]{direction: direction, size: size, data: data}
}

func (g *Group[K, V]) Direction() common.VectorType {
	return g.direction
}

func (g *Group[K, V]) Size() K {
	return g.size
}

func (g *Group[K, V]) Data() []*Vector[K, V] {
	return g.data
}

func (g *Group[K, V]) Add(vec *Vector[K, V]) {
	if vec.Direction() != g.Direction() || vec.Size() != g.Size() {
		log.Panicln("不同类型的向量无法添加")
	}
	g.data = append(g.data, vec)
}

func (g *Group[K, V]) ToMatrix() *matrix.Matrix[K, V] {
	items := make(map[K][]V, len(g.Data()))
	for i, v := range g.Data() {
		data := v.Data()
		item := make([]V, v.Size())
		//log.Println(item, data)
		for j := K(0); j < v.Size(); j++ {
			item[j] = data[j+1]
		}
		//log.Println(item, data)
		items[K(i+1)] = item
	}
	target := matrix.NewMatrix(K(len(g.Data())), g.Size(), items)
	if g.Direction() == common.VERTICAL {
		target.Transpose()
	}
	return target
}

// Relevant 是否线性相关
func (g *Group[K, V]) Relevant() bool {
	m := g.ToMatrix()
	return m.Rank() < m.Row()
}
