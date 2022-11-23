package common

import (
	"fmt"
	"log"
	"math"
)

// GetSerialByKey 通过key值计算矩阵下标
func GetSerialByKey[K Key](number K, col K) (K, K) {
	j := number % col
	if j == 0 {
		j = col
	}
	return K(math.Ceil(float64(number) / float64(col))), j
}

// GetKeyBySerial 通过行列下标计算key的值
func GetKeyBySerial[K Key](i K, j K, col K) K {
	return (i-1)*col + j
}

// OutOfBounds 校验是否越界
func OutOfBounds[K Key](i K, j K, row K, col K) {
	if i > row || j > col {
		log.Panicln(fmt.Sprintf("out of bounds, maxRow = %d, maxCol = %d", row, col))
	}
}

// OutOfBoundsMax 校验越界
func OutOfBoundsMax[K Key](i K, max K) {
	if i > max {
		log.Panicln(fmt.Sprintf("out of bounds, max = %d", max))
	}
}

// Min 获取两列中的最小数
func Min[K Key](a K, b K) K {
	if a > b {
		return b
	} else {
		return a
	}
}

// Max 获取两列中的最大数
func Max[K Key](a K, b K) K {
	if a < b {
		return b
	} else {
		return a
	}
}
