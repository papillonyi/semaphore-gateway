package bubblesort

import "fmt"

type Rect struct {
	x, y          float64
	width, height float64
}

type Integer int

func (a *Integer) Less(b Integer) bool {
	return *a < b
}

type LessAdder interface {
	Less(b Integer) bool
	//Add(b Integer)
}

func BubbleSort(values []int) {
	reca := &Rect{}

	fmt.Print(reca)
	for i := 0; i < len(values)-1; i++ {
		flag := true

		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}

		if flag {
			fmt.Print(values[i])
			break
		}

	}
}

func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}

		if j >= p {
			values[p] = values[j]
			p = j
		}

		if values[i] <= temp && i <= p {
			i++
		}

		if i <= p {
			values[p] = values[i]
			p = i
		}
	}

	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}

}

func QuickSort(value []int) {
	quickSort(value, 0, len(value)-1)
}
