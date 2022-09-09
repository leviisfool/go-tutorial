package main

import (
	"math"
)

const L = 30

type BitTree struct {
	chidren [2]*BitTree
	min     int
}

func (bt *BitTree) insert(val int) {
	node := bt
	if node.min > val {
		node.min = val
	}
	for i := L - 1; i >= 0; i-- {
		b := (val >> i) & 1
		if node.chidren[b] == nil {
			node.chidren[b] = &BitTree{min: val}
		}
		node = node.chidren[b]
		if node.min > val {
			node.min = val
		}
	}
}

func (bt *BitTree) maxXOR(val, limit int) int {
	node := bt

	if limit < node.min {
		return -1
	}

	ans := 0
	for i := L - 1; i >= 0; i-- {
		b := (val >> i) & 1
		xor := b ^ 1&1

		if node.chidren[xor] != nil && node.chidren[xor].min <= limit {
			ans = (ans << (i + 1)) | xor
			node = node.chidren[xor]
		} else {
			ans = (ans << (i + 1)) | b
			node = node.chidren[b]
		}
	}
	return ans ^ val
}

func maximizeXor(nums []int, queries [][]int) []int {
	bitTree := &BitTree{min: math.MaxInt32}
	for _, n := range nums {
		bitTree.insert(n)
	}
	var ans = make([]int, len(queries))

	for i, q := range queries {
		ans[i] = bitTree.maxXOR(q[0], q[1])
	}
	return ans
}

func main() {
	ans := maximizeXor([]int{5, 2, 4, 6, 6, 3}, [][]int{[]int{12, 4}, []int{8, 1}, []int{6, 3}})
	for _, v := range ans {
		println(v)
	}
}
