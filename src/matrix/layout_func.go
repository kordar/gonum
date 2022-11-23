package matrix

import (
	"fmt"
	"github.com/kordar/gonum/src/common"
	"log"
	"sort"
)

type Segment[K common.Key] struct {
	Offset K // 偏移值
	Top    K // 上边界
	Bottom K // 下边界
	Value  K // 分割值
}

func removeDuplicateElement[K common.Key](data []K) []K {
	result := make([]K, 0, len(data))
	temp := map[K]struct{}{}
	for _, item := range data {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func SplitSegment[K common.Key](s []K, max K) []Segment[K] {
	segments := make([]Segment[K], 0)
	if s == nil || len(s) == 0 {
		segments = append(segments, Segment[K]{Top: 0, Offset: 0, Value: max, Bottom: max})
		return segments
	}
	s = append(s, max)
	s = removeDuplicateElement(s)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	// 循环已排序列表，分别计算分割的上边界、下边界
	for i, k := range s {
		if k > max {
			log.Panicln("分割值不能大于最大边界值")
		}
		segment := Segment[K]{Bottom: k, Value: k}
		if i == 0 {
			// 首元素，上边界一点为0，下边界一定为k
			segment.Top = 0
			segment.Offset = 0
		} else {
			segment.Top = s[i-1]
			segment.Offset = s[i-1]
		}
		segments = append(segments, segment)
	}

	sort.Slice(segments, func(i, j int) bool {
		return segments[i].Value < segments[j].Value
	})

	return segments
}

func SearchSegment[K common.Key](i K, segments []Segment[K]) Segment[K] {
	var seg Segment[K]
	for _, segment := range segments {
		if i <= segment.Value {
			seg = segment
			break
		}
	}
	return seg
}

func SegmentMatrix[K common.Key, V common.Value](matrix *Matrix[K, V], segRow []K, segCol []K) *Layout[K, V] {
	layout := &Layout[K, V]{}
	layout.Row = matrix.Row()
	layout.Col = matrix.Col()

	segmentsOfRow := SplitSegment(segRow, matrix.Row())
	segmentsOfCol := SplitSegment(segCol, matrix.Col())

	// num := (len(segRow) + 1) * (len(segCol) + 1) // 计算切换后的矩阵数量
	mm := make(map[string]*BlockMatrix[K, V])
	layout.Sections = make([]*BlockMatrix[K, V], 0)

	for k, v := range matrix.Data() {
		i, j := matrix.GetSerialByNumber(k)
		iOfSeg := SearchSegment(i, segmentsOfRow)
		jOfSeg := SearchSegment(j, segmentsOfCol)
		offsetRow, ti, maxRow := iOfSeg.Offset, i-iOfSeg.Offset, iOfSeg.Bottom-iOfSeg.Top
		offsetCol, tj, maxCol := jOfSeg.Offset, j-jOfSeg.Offset, jOfSeg.Bottom-jOfSeg.Top
		//log.Println(i, "-", j, "=", v, ",", offsetRow, "->", offsetCol, ",", ti, ":", tj, "|", maxRow, maxCol)
		kk := fmt.Sprintf("%v_%v", offsetRow, offsetCol)
		b := mm[kk]
		if b == nil {
			b = &BlockMatrix[K, V]{offsetRow, offsetCol, NewMatrix[K, V](maxRow, maxCol, nil)}
			mm[kk] = b
		}
		b.Matrix.Save(ti, tj, v)
	}

	for _, b := range mm {
		layout.Sections = append(layout.Sections, b)
	}

	return layout
}

func PrintBlockMatrix[K common.Key, V common.Value](block *BlockMatrix[K, V]) {
	fmt.Println("***************************")
	fmt.Println(fmt.Sprintf("** row:%v, col:%v, top:%v, left: %v", block.Matrix.Row(), block.Matrix.Col(), block.OffsetRow, block.OffsetCol))
	fmt.Println("***************************")
	Print(block.Matrix, 1, 1, block.Matrix.Row(), block.Matrix.Col())
}
