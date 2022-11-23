package common

import (
	"fmt"
)

// Arrangement 数学方法计算排列数(从n中取m个数)
func Arrangement[K Key](n K, m K) K {
	return Factorial(n) / Factorial(n-m)
}

// Combination 数学方法计算组合数(从n中取m个数)
func Combination[K Key](n K, m K) K {
	return Factorial(n) / (Factorial(n-m) * Factorial(m))
}

// Factorial 阶乘
func Factorial[K Key](n K) K {
	result := K(1)
	for i := K(2); i <= n; i++ {
		result *= i
	}
	return result
}

func CombinationList[K Key](n K, m K) [][]K {
	result := CombinationResult(n, m)
	data := make([][]K, len(result))
	for i, item := range result {
		d := make([]K, 0)
		for j, k := range item {
			if k == 1 {
				d = append(d, K(j+1))
			}
		}
		data[i] = d
	}
	return data
}

// CombinationResult 组合算法(从nums中取出m个数)
func CombinationResult[K Key](n K, m K) [][]K {
	if m < 1 || m > n {
		fmt.Println("Illegal argument. Param m must between 1 and len(nums).")
		return [][]K{}
	}
	// 保存最终结果的数组，总数直接通过数学公式计算
	result := make([][]K, 0, Combination(n, m))
	// 保存每一个组合的索引的数组，1表示选中，0表示未选中
	indexs := make([]K, n)
	for i := K(0); i < n; i++ {
		if i < m {
			indexs[i] = 1
		} else {
			indexs[i] = 0
		}
	}
	// 第一个结果
	result = addTo(result, indexs)
	for {
		find := false
		// 每次循环将第一次出现的 1 0 改为 0 1，同时将左侧的1移动到最左侧
		for i := K(0); i < n-1; i++ {
			if indexs[i] == 1 && indexs[i+1] == 0 {
				find = true
				indexs[i], indexs[i+1] = 0, 1
				if i > 1 {
					moveOneToLeft(indexs[:i])
				}
				result = addTo(result, indexs)
				break
			}
		}
		// 本次循环没有找到 1 0 ，说明已经取到了最后一种情况
		if !find {
			break
		}
	}
	return result
}

//将ele复制后添加到arr中，返回新的数组
func addTo[K Key](arr [][]K, ele []K) [][]K {
	newEle := make([]K, len(ele))
	copy(newEle, ele)
	arr = append(arr, newEle)
	return arr
}

func moveOneToLeft[K Key](leftNums []K) {
	//计算有几个1
	sum := 0
	for i := 0; i < len(leftNums); i++ {
		if leftNums[i] == 1 {
			sum++
		}
	}
	//将前sum个改为1，之后的改为0
	for i := 0; i < len(leftNums); i++ {
		if i < sum {
			leftNums[i] = 1
		} else {
			leftNums[i] = 0
		}
	}
}

// FindNumsByIndexs 根据索引号数组得到元素数组
func FindNumsByIndexs(nums []int, indexs [][]int) [][]int {
	if len(indexs) == 0 {
		return [][]int{}
	}
	result := make([][]int, len(indexs))
	for i, v := range indexs {
		line := make([]int, 0)
		for j, v2 := range v {
			if v2 == 1 {
				line = append(line, nums[j])
			}
		}
		result[i] = line
	}
	return result
}

// ArrangementResult 组合结果
func ArrangementResult[K Key](n K, m K) [][]K {
	//组合结果
	c := CombinationResult(n, m)
	//保存最终排列结果
	result := make([][]K, 0)
	//遍历组合结果，对每一项进行全排列
	for _, v := range c {
		p := ArrangementAll(v)
		result = append(result, p...)
	}
	return result
}

// ArrangementAll n个数全排列
// 如输入[1 2 3]，则返回[123 132 213 231 312 321]
func ArrangementAll[K Key](nums []K) [][]K {
	COUNT := len(nums)
	//检查
	if COUNT == 0 || COUNT > 10 {
		panic("Illegal argument. nums size must between 1 and 9.")
	}
	//如果只有一个数，则直接返回
	if COUNT == 1 {
		return [][]K{nums}
	}
	//否则，将最后一个数插入到前面的排列数中的所有位置
	return insertItem(ArrangementAll(nums[:COUNT-1]), nums[COUNT-1])
}

func insertItem[K Key](res [][]K, insertNum K) [][]K {
	//保存结果的slice
	result := make([][]K, len(res)*(len(res[0])+1))
	index := 0
	for _, v := range res {
		for i := 0; i < len(v); i++ {
			//在v的每一个元素前面插入新元素
			result[index] = insertToSlice(v, K(i), insertNum)
			index++
		}
		//在v最后面插入新元素
		result[index] = append(v, insertNum)
		index++
	}
	return result
}

// 将元素value插入到数组nums中索引为index的位置
func insertToSlice[K Key](nums []K, index K, value K) []K {
	result := make([]K, len(nums)+1)
	copy(result[:index], nums[:index])
	result[index] = value
	copy(result[index+1:], nums[index:])
	return result
}
