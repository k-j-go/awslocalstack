package funcs

import (
	"fmt"
	underscore "github.com/ahl5esoft/golang-underscore"
	"testing"
)

type car struct {
	name string
}

type CarIterator interface {
	Next() (value string, ok bool)
}

type Collection[T car] struct {
	index int
	List  []T
}

func (collection *Collection[T]) Next() (value *T, ok bool) {
	collection.index++
	if collection.index >= len(collection.List) {
		return nil, false
	}
	return &collection.List[collection.index], true
}

func newSlice[T car](s []T) *Collection[T] {
	return &Collection[T]{-1, s}
}

func TestCars_Func_Way(t *testing.T) {
	values := NewFromSlice([]car{car{"toyota"}}).Map(func(x car) car {
		return x
	}).Filter(func(x car) bool {
		check := func(c car) bool {
			if c.name == "toyota" {
				return true
			} else {
				return false
			}
		}

		return check(x)
	}).values

	for _, x := range values {
		fmt.Printf("%v \n", x)
	}
}

func Test_Under_Score(t *testing.T) {
	var res []int
	underscore.Chain([]int{1, 2}).Aggregate(
		func(memo []int, n, _ int) []int {
			memo = append(memo, n)
			memo = append(memo, n+10)
			return memo
		},
		make([]int, 0),
	).Value(&res)
}
