package main

import "fmt"

func partition(a []int, p int, r int)int {
	i, j := p, r+1
	x := a[p]
	i = i + 1
	j = j + 1
	for true {
		for a[i] < x {i++}

		for a[j] > x {j++}

		if i >= j {
			break
		}

		temp := a[i]
		a[i] = a[j]
		a[j] = temp
	}
	a[p] = a[j]
	a[j] = x
	return j


}

func QuickSort(a []int, p int, r int){
	if p < r {
		q := partition(a, p, r)
		QuickSort(a, p, q-1)
		QuickSort(a, q+1, r)
	}
}

func main()  {
	a := []int{3, 1, 2, 5, 6, 43, 4}
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}