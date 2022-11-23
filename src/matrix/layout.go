package matrix

import "gonum/src/common"

// BlockMatrix 分块矩阵
type BlockMatrix[K common.Key, V common.Value] struct {
	OffsetRow K // 行偏移值
	OffsetCol K // 列偏移值
	Matrix    *Matrix[K, V]
}

func (block *BlockMatrix[K, V]) GetMatrix() *Matrix[K, V] {
	return block.Matrix
}

type Layout[K common.Key, V common.Value] struct {
	Row      K
	Col      K
	Sections []*BlockMatrix[K, V]
}

func NewLayout[K common.Key, V common.Value](matrix *Matrix[K, V], segRow []K, segCol []K) *Layout[K, V] {
	return SegmentMatrix[K, V](matrix, segRow, segCol)
}

func (l *Layout[K, V]) ToMatrix() *Matrix[K, V] {
	matrix := NewMatrix[K, V](l.Row, l.Col, nil)
	for _, section := range l.Sections {
		for k, v := range section.Matrix.Data() {
			i, j := section.Matrix.GetSerialByNumber(k)
			matrix.Save(section.OffsetRow+i, section.OffsetCol+j, v)
		}
	}
	return matrix
}
