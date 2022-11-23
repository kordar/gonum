package vector

import "gonum/src/common"

// Similar 判断两个向量是否相似
func Similar[K common.Key, V common.Value](a *BaseVector[K, V], b *BaseVector[K, V]) bool {
	return a.Size() == b.Size()
}

// IsZero 是否零向量
func IsZero[K common.Key, V common.Value](v *BaseVector[K, V]) bool {
	return len(v.Data()) == 0
}

// Equal 判断向量是否相等
func Equal[K common.Key, V common.Value](a *BaseVector[K, V], b *BaseVector[K, V]) bool {
	if !Similar(a, b) {
		return false
	}
	if len(a.Data()) != len(b.Data()) {
		return false
	}
	t := a.Data()
	for k, v := range b.Data() {
		if v != t[k] {
			return false
		}
	}
	return true
}
