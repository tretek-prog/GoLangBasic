package main

import (
	"fmt"
)

func vstavka(a []int) {
	var n = len(a)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
			j = j - 1
		}
	}
}

func main() {
	a := []int{1,4,-4,6,1,0} // сюда записываем срез для сортировки
	vstavka(a)
	fmt.Println(a)
}
