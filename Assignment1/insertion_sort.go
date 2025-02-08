package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Insertion_sort(arr []int) (array []int) {
	for i := 0; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
	return arr
}

func run_insertion_sorts() (arr []time.Duration, n_times []int) {
	n_array := [8]int{50, 100, 250, 500, 750, 1000, 2000, 5000}
	n_time_used := make([]time.Duration, 8)

	for i, n := range n_array {
		arr := make([]int, n)

		for arrI := range arr {
			arr[arrI] = rand.Intn(n)
		}
		fmt.Printf("Unsorted array: %v\n", arr)

		start_time := time.Now()
		sortedArray := Insertion_sort(arr)
		time_since_start := time.Since(start_time)
		fmt.Printf("Sorted array: %v", sortedArray)
		fmt.Printf("Go Insertion Sort Time: %.4f ms\n", time_since_start.Seconds()*1000)
		n_time_used[i] = time_since_start
	}
	return n_time_used, n_array[:]
}

func main() {
	// size := 1000
	// arr := make([]int, size)

	// for i := range arr {
	// 	arr[i] = rand.Intn(size)
	// }
	// fmt.Printf("Unsorted array: %v\n", arr)

	// start_time := time.Now()
	// sortedArray := Insertion_sort(arr)
	// time_since_start := time.Since(start_time)
	// fmt.Printf("Sorted array: %v", sortedArray)
	// fmt.Printf("Go Insertion Sort Time: %.4f ms\n", time_since_start.Seconds()*1000)
	time_elapsed, n_runs := run_insertion_sorts()
	for i, elapsed := range time_elapsed {
		fmt.Printf("Go Insertion Sort Time: %.4f ms with n=%v\n", elapsed.Seconds()*1000, n_runs[i])
	}
}
