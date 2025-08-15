package main

import "fmt"

func Arrayminmax() {
	array := [5]int{20, 30, 15, 19, 10}

	min := array[0]
	max := array[0]

	for i := 0; i < len(array); i++ {
		if min > array[i] {
			min = array[i]
		}
	}
	for j := 0; j < len(array); j++ {
		if max < array[j] {
			max = array[j]
		}
	}

	fmt.Println("Minimum value in the array : ", min)
	fmt.Println("Maximum value in the array : ", max)

}
