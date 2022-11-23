package common

import (
	"log"
	"testing"
)

func TestArrangement(t *testing.T) {
	arrangement := Arrangement[int](3, 2)
	log.Println("A3-2=", arrangement)
	log.Println("C3-2=", Combination(3, 2))
}

func TestConsecutiveElements(t *testing.T) {
	elements := CombinationResult[int](4, 3)
	log.Println(elements)
	log.Println(Combination[int](4, 3))
	log.Println(CombinationList[int](4, 3))
}
