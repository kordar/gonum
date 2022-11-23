package det

import (
	"fmt"
	"github.com/kordar/gonum/src/common"
)

// Save 通过行列保存数据
func Save[K common.Key, V common.Value](Determinant *Determinant[K, V], i K, j K, v V) *Determinant[K, V] {
	target := Determinant.Clone()
	target.Save(i, j, v)
	return target
}

// SaveLine 保存一行
func SaveLine[K common.Key, V common.Value](Determinant *Determinant[K, V], i K, v ...V) *Determinant[K, V] {
	target := Determinant.Clone()
	target.SaveLine(i, v...)
	return target
}

// SaveLines 保存多行
func SaveLines[K common.Key, V common.Value](Determinant *Determinant[K, V], data map[K][]V) *Determinant[K, V] {
	target := Determinant.Clone()
	target.SaveLines(data)
	return target
}

// Get 通过行列坐标获取值
func Get[K common.Key, V common.Value](Determinant *Determinant[K, V], i K, j K) V {
	return Determinant.Get(i, j)
}

func PrintAll[K common.Key, V common.Value](Determinant *Determinant[K, V]) {
	fmt.Println("***************************")
	fmt.Println(fmt.Sprintf("** determinant, row:%v, col:%v", Determinant.Row(), Determinant.Col()))
	fmt.Println("***************************")
	Print(Determinant, 1, 1, Determinant.Row(), Determinant.Col())
}

func Print[K common.Key, V common.Value](Determinant *Determinant[K, V], sRow K, sCol K, eRow K, eCol K) {
	common.OutOfBounds(sRow, sCol, Determinant.Row(), Determinant.Col())
	common.OutOfBounds(eRow, eCol, Determinant.Row(), Determinant.Col())
	for r := sRow; r <= eRow; r++ {
		for c := sCol; c <= eCol; c++ {
			fmt.Printf("%v\t", Determinant.Get(r, c))
		}
		fmt.Print("\n")
	}
}
