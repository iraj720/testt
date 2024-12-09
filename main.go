package main

import "fmt"

func main() {
	arr := []int32{5, 1, 2, 3, 4}

	x := int32(4)

	index := mergeSearch(arr, int32(len(arr)/2))

	findIn1 := binarySearch(arr[:index], x)
	findIn2 := binarySearch(arr[index:], x)

	fmt.Println(index)
	fmt.Println(findIn1)
	fmt.Println(findIn1 || findIn2)
}

func mergeSearch(arr []int32, searchIndex int32) int32 {
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return searchIndex
		}
	}

	last := int32(len(arr) - 1)
	first := arr[0]

	searchIndexTemp := int32(len(arr) / 2)

	for {
		if arr[searchIndexTemp] > first {
			return mergeSearch(arr[searchIndexTemp:], searchIndexTemp)
		} else if arr[searchIndexTemp] <= last {
			return mergeSearch(arr[:searchIndexTemp+1], searchIndexTemp)
		}
	}
}

func binarySearch(arr []int32, x int32) bool {
	if len(arr) == 1 {
		return x == arr[0]
	}

	searchIndex := len(arr) / 2
	for {
		if arr[searchIndex] > x {
			return binarySearch(arr[:searchIndex], x)
		} else if arr[searchIndex] <= x {
			return binarySearch(arr[searchIndex:], x)
		}
	}
}
