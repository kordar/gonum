package common

type Key interface {
	~int
}

type Value interface {
	~int | ~float32 | ~float64
}

type VectorType string

const (
	VERTICAL   VectorType = "vertical"
	HORIZONTAL VectorType = "horizontal"
)
