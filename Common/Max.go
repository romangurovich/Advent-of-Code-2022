package common

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](a T, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}
