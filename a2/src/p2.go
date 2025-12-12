package src

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	// concurrentThreshold is the minimum array size for concurrent sorting
	concurrentThreshold = 10000
	// maxDepth is the maximum recursion depth for concurrent sorting
	maxDepth = 4
)

// serialQuicksort is a standard quicksort implementation
func serialQuicksort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	
	pivot := arr[len(arr)/2]
	left := 0
	right := len(arr) - 1
	
	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	
	if right > 0 {
		serialQuicksort(arr[:right+1])
	}
	if left < len(arr) {
		serialQuicksort(arr[left:])
	}
}

// partition partitions the array around a pivot, returning left, middle (equal), and right partitions
func partition(arr []int) ([]int, []int, []int) {
	if len(arr) == 0 {
		return nil, nil, nil
	}
	
	pivot := arr[len(arr)/2]
	var left, middle, right []int
	
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v == pivot {
			middle = append(middle, v)
		} else {
			right = append(right, v)
		}
	}
	
	return left, middle, right
}

// concurrentQuicksort implements quicksort using goroutines and channels
func concurrentQuicksort(arr []int, depth int) []int {
	if len(arr) <= 1 {
		return arr
	}
	
	// Use serial sort for small arrays or deep recursion to avoid goroutine overhead
	if len(arr) < concurrentThreshold || depth > maxDepth {
		serialQuicksort(arr)
		return arr
	}
	
	left, middle, right := partition(arr)
	
	// Use channels to synchronize goroutines
	leftChan := make(chan []int, 1)
	rightChan := make(chan []int, 1)
	
	// Sort left partition in a goroutine
	go func() {
		leftChan <- concurrentQuicksort(left, depth+1)
	}()
	
	// Sort right partition in a goroutine
	go func() {
		rightChan <- concurrentQuicksort(right, depth+1)
	}()
	
	// Wait for results via channels
	sortedLeft := <-leftChan
	sortedRight := <-rightChan
	
	// Combine results: left + middle (pivot elements) + right
	result := make([]int, 0, len(arr))
	result = append(result, sortedLeft...)
	result = append(result, middle...)
	result = append(result, sortedRight...)
	
	return result
}

// isSorted checks if an array is sorted
func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func P2() {
	fmt.Println("Part B: Concurrent Quicksort")
	
	sizes := []int{10000, 50000, 100000, 500000}
	
	for _, size := range sizes {
		fmt.Printf("\n--- Array size: %d ---\n", size)
		
		// Generate random array
		arr := make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = rand.Intn(1000000)
		}
		
		// Test serial quicksort
		arr1 := make([]int, len(arr))
		copy(arr1, arr)
		
		start := time.Now()
		serialQuicksort(arr1)
		serialTime := time.Since(start)
		
		fmt.Printf("Serial Quicksort:     %v (sorted: %v)\n", serialTime, isSorted(arr1))
		
		// Test concurrent quicksort
		arr2 := make([]int, len(arr))
		copy(arr2, arr)
		
		start = time.Now()
		arr2 = concurrentQuicksort(arr2, 0)
		concurrentTime := time.Since(start)
		
		fmt.Printf("Concurrent Quicksort: %v (sorted: %v)\n", concurrentTime, isSorted(arr2))
		
		// Test Go's built-in sort for comparison
		arr3 := make([]int, len(arr))
		copy(arr3, arr)
		
		start = time.Now()
		sort.Ints(arr3)
		builtinTime := time.Since(start)
		
		fmt.Printf("Go sort.Ints:         %v (sorted: %v)\n", builtinTime, isSorted(arr3))
		
		// Calculate speedup
		if serialTime > concurrentTime {
			speedup := float64(serialTime) / float64(concurrentTime)
			fmt.Printf("Speedup: %.2fx faster\n", speedup)
		} else {
			slowdown := float64(concurrentTime) / float64(serialTime)
			fmt.Printf("Speedup: %.2fx slower\n", slowdown)
		}
	}
}
