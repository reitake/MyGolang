package main

import (
	"fmt"
)

type Sorter interface { // 声明Sorter接口类型，及其包含的方法
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type IntArray []int // 声明类型

// 为类型添加接口的方法
func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

func main() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := IntArray(data)
	fmt.Println("Unsorted:", a)
	var aI Sorter
	aI = a
	Sort(aI)
	fmt.Println("Sorted:", a)

}
