package main

import (
	"fmt"
)
func merge (a []int, b []int) (c []int) {   //Функция слияния двух срезов
	//Срезы должны быть отсортированы или по возрастанию или по убыванию
	n := len(a)
	m := len(b)
	i := 0
	j := 0
	k := 0

	for i < n && j < m {
		if a[i] <= b[j] {
			c = append(c, a[i])
			i++

		}else{
			c = append(c, b[j])
			j++
		}
	}  // Добавляем в срез с оставшиеся элементы среза а и среза b
	if i < n {
		for i < n {
			c = append(c, a[i])
			i++
		}
	}else{
		if j < m {
			for j < m {
				c = append(c, b[j])
				j++
			}
		}
	}
	return c
}
//Рекурсивная функция, которая делит исходный срез на маленькие срезы и отправляет их на сортировку
func splitAndMergeList(a []int) ([]int) {
	n1 := len(a) / 2
	a1 := a[:n1]  //Деление исходного среза на примерно равные
	a2 := a[n1:]

	if len(a1) > 1{  //Если длина 1-го среза больше 1, то делиш дальше
		a1 = splitAndMergeList(a1)
	}
	if len(a2) > 1{  //Если длина 2-го среза больше 1, то делим дальше
		a2 = splitAndMergeList(a2)
	}
	return merge(a1, a2)  //отправка срезов на сортировку
}


func main () {
	a := []int{1, 10, 0, -5, 8, 9, 6} // cюда вводим срез для сортировки
	a = splitAndMergeList(a)
	fmt.Println(a)
}
