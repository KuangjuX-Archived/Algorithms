package main

import "fmt"

func Merge(a []int, l int, mid int, r int){
	temp := make([]int , r-l+5);
	len := 0
	i := l
	j := mid + 1
	for i <= mid && j <= r{
		if a[i] <= a[j]{
			temp[len] = a[i]
			len += 1
			i += 1
		}else{
			temp[len] = a[j]
			len += 1
			j += 1
		}
	}

	for i <= mid {
		temp[len] = a[i]
		len += 1
		i += 1
	}

	for j <= r {
		temp[len] = a[j]
		len += 1
		j += 1
	}


	for m, k := 0, l; k <= r; m, k = m+1, k+1{
		a[k] = temp[m]
	}

	
}

func MergeSort(a []int, l int, r int){
	if l >= r{
		return
	}

	mid := (l + r) / 2
	MergeSort(a, l, mid)
	MergeSort(a, mid+1, r)

	Merge(a, l, mid, r)
}

func main()  {
	a := []int{3, 1, 2, 5, 6, 43, 4}
	// a := []int{6, 7, 10, 8, 4, 5, 3}
	MergeSort(a, 0, len(a)-1)

	fmt.Println(a)
}

