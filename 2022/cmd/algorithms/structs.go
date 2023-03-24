package algorithms

import (
	"golang.org/x/exp/constraints"
)

type PriorityQueue[T constraints.Ordered] struct {
	heap []T
}
