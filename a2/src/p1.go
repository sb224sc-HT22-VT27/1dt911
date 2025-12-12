package src

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	minInt = math.MinInt
	maxInt = math.MaxInt
)

// Node represents a node in the linked list
type Node struct {
	key  int
	next *Node
	mu   sync.Mutex // For fine-grained locking
}

// LinkedList represents a sorted linked list
type LinkedList struct {
	head *Node
	mu   sync.Mutex // For coarse-grained locking
}

// NewLinkedList creates a new empty linked list with sentinel nodes
func NewLinkedList() *LinkedList {
	// Head sentinel with minimum value
	head := &Node{key: minInt}
	// Tail sentinel with maximum value
	tail := &Node{key: maxInt}
	head.next = tail
	return &LinkedList{head: head}
}

// find returns the predecessor node and the node with the given key
func (ll *LinkedList) find(key int) (*Node, *Node) {
	pred := ll.head
	curr := pred.next
	for curr.key < key {
		pred = curr
		curr = curr.next
	}
	return pred, curr
}

// Add adds a node with the given key (sequential version)
func (ll *LinkedList) Add(key int) bool {
	pred, curr := ll.find(key)
	if curr.key == key {
		return false // Already exists
	}
	node := &Node{key: key, next: curr}
	pred.next = node
	return true
}

// Remove removes a node with the given key (sequential version)
func (ll *LinkedList) Remove(key int) bool {
	pred, curr := ll.find(key)
	if curr.key != key {
		return false // Not found
	}
	pred.next = curr.next
	return true
}

// Contains checks if a node with the given key exists
func (ll *LinkedList) Contains(key int) bool {
	_, curr := ll.find(key)
	return curr.key == key
}

// Count returns the number of nodes in the list
func (ll *LinkedList) Count() int {
	count := 0
	curr := ll.head.next
	for curr != nil && curr.key != maxInt {
		count++
		curr = curr.next
	}
	return count
}

// StrictlyIncreasing checks if the list is strictly increasing
func (ll *LinkedList) StrictlyIncreasing() bool {
	curr := ll.head.next
	for curr != nil && curr.next != nil && curr.key != maxInt {
		if curr.key >= curr.next.key {
			return false
		}
		curr = curr.next
	}
	return true
}

// CoarseGrainedList implements coarse-grained locking
type CoarseGrainedList struct {
	*LinkedList
}

func NewCoarseGrainedList() *CoarseGrainedList {
	return &CoarseGrainedList{NewLinkedList()}
}

func (ll *CoarseGrainedList) Add(key int) bool {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	return ll.LinkedList.Add(key)
}

func (ll *CoarseGrainedList) Remove(key int) bool {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	return ll.LinkedList.Remove(key)
}

func (ll *CoarseGrainedList) Contains(key int) bool {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	return ll.LinkedList.Contains(key)
}

// FineGrainedList implements fine-grained (hand-over-hand) locking
type FineGrainedList struct {
	head *Node
}

func NewFineGrainedList() *FineGrainedList {
	head := &Node{key: minInt}
	tail := &Node{key: maxInt}
	head.next = tail
	return &FineGrainedList{head: head}
}

func (ll *FineGrainedList) Add(key int) bool {
	ll.head.mu.Lock()
	pred := ll.head
	curr := pred.next
	curr.mu.Lock()
	
	for curr.key < key {
		pred.mu.Unlock()
		pred = curr
		curr = curr.next
		curr.mu.Lock()
	}
	
	defer pred.mu.Unlock()
	defer curr.mu.Unlock()
	
	if curr.key == key {
		return false
	}
	
	node := &Node{key: key, next: curr}
	pred.next = node
	return true
}

func (ll *FineGrainedList) Remove(key int) bool {
	ll.head.mu.Lock()
	pred := ll.head
	curr := pred.next
	curr.mu.Lock()
	
	for curr.key < key {
		pred.mu.Unlock()
		pred = curr
		curr = curr.next
		curr.mu.Lock()
	}
	
	defer pred.mu.Unlock()
	defer curr.mu.Unlock()
	
	if curr.key != key {
		return false
	}
	
	pred.next = curr.next
	return true
}

func (ll *FineGrainedList) Contains(key int) bool {
	ll.head.mu.Lock()
	pred := ll.head
	curr := pred.next
	curr.mu.Lock()
	
	for curr.key < key {
		pred.mu.Unlock()
		pred = curr
		curr = curr.next
		curr.mu.Lock()
	}
	
	pred.mu.Unlock()
	curr.mu.Unlock()
	
	return curr.key == key
}

func (ll *FineGrainedList) Count() int {
	count := 0
	ll.head.mu.Lock()
	curr := ll.head.next
	curr.mu.Lock()
	ll.head.mu.Unlock()
	
	for curr != nil && curr.key != maxInt {
		count++
		next := curr.next
		if next != nil {
			next.mu.Lock()
		}
		curr.mu.Unlock()
		curr = next
	}
	
	if curr != nil {
		curr.mu.Unlock()
	}
	
	return count
}

// OptimisticList implements optimistic locking with atomic operations
type OptimisticList struct {
	head *Node
	mu   sync.Mutex // Optional: can be disabled for testing
	useLocking bool
}

func NewOptimisticList(useLocking bool) *OptimisticList {
	head := &Node{key: minInt}
	tail := &Node{key: maxInt}
	head.next = tail
	return &OptimisticList{head: head, useLocking: useLocking}
}

func (ll *OptimisticList) validate(pred, curr *Node) bool {
	return atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&pred.next))) == unsafe.Pointer(curr)
}

func (ll *OptimisticList) Add(key int) bool {
	for {
		pred := ll.head
		curr := pred.next
		
		for curr.key < key {
			pred = curr
			curr = curr.next
		}
		
		if ll.useLocking {
			pred.mu.Lock()
			curr.mu.Lock()
			
			if ll.validate(pred, curr) {
				defer pred.mu.Unlock()
				defer curr.mu.Unlock()
				
				if curr.key == key {
					return false
				}
				
				node := &Node{key: key, next: curr}
				pred.next = node
				return true
			}
			
			pred.mu.Unlock()
			curr.mu.Unlock()
		} else {
			// WARNING: Without locking, this has intentional race conditions for demonstration
			// Multiple goroutines can simultaneously modify pred.next causing data corruption
			if curr.key == key {
				return false
			}
			node := &Node{key: key, next: curr}
			pred.next = node
			return true
		}
	}
}

func (ll *OptimisticList) Contains(key int) bool {
	pred := ll.head
	curr := pred.next
	
	for curr.key < key {
		pred = curr
		curr = curr.next
	}
	
	if ll.useLocking {
		pred.mu.Lock()
		curr.mu.Lock()
		defer pred.mu.Unlock()
		defer curr.mu.Unlock()
		
		if ll.validate(pred, curr) {
			return curr.key == key
		}
		return false
	}
	
	return curr.key == key
}

func (ll *OptimisticList) Count() int {
	count := 0
	curr := ll.head.next
	for curr != nil && curr.key != maxInt {
		count++
		curr = curr.next
	}
	return count
}

// Experiment functions

func runSequentialExperiment() {
	fmt.Println("\n--- A.1: Sequential Linked List ---")
	
	ll := NewLinkedList()
	numNodes := 5000
	
	// Single goroutine
	start := time.Now()
	for i := 0; i < numNodes; i++ {
		ll.Add(rand.Intn(1000000))
	}
	elapsed := time.Since(start)
	
	fmt.Printf("Single goroutine: Added %d nodes in %v\n", ll.Count(), elapsed)
	fmt.Printf("Strictly increasing: %v\n", ll.StrictlyIncreasing())
	
	// Multiple goroutines (will have race conditions)
	ll2 := NewLinkedList()
	var wg sync.WaitGroup
	numGoroutines := 4
	nodesPerGoroutine := numNodes / numGoroutines
	
	start = time.Now()
	for g := 0; g < numGoroutines; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < nodesPerGoroutine; i++ {
				ll2.Add(rand.Intn(1000000))
			}
		}()
	}
	wg.Wait()
	elapsed = time.Since(start)
	
	fmt.Printf("Multiple goroutines (%d): Added %d nodes in %v (may have race conditions)\n", 
		numGoroutines, ll2.Count(), elapsed)
}

func runCoarseGrainedExperiment() {
	fmt.Println("\n--- A.2.1 & A.2.2: Coarse-Grained Locking ---")
	
	numCores := runtime.NumCPU()
	numNodes := 10000 // Reduced for faster testing
	maxValue := 1000000
	
	fmt.Printf("Testing with %d nodes, %d cores\n", numNodes, numCores)
	
	// Random values experiment
	fmt.Println("\nRandom values:")
	goroutineCounts := []int{1, 2, 4, 8, 16, 32}
	if numCores > 1 {
		maxG := 32 * numCores
		if maxG > 32 {
			goroutineCounts = append(goroutineCounts, 64, 128)
		}
	}
	for _, g := range goroutineCounts {
		runCoarseTest(g, numNodes, maxValue, true)
	}
	
	// Sequential values experiment
	fmt.Println("\nSequential values:")
	for _, g := range goroutineCounts {
		runCoarseTest(g, numNodes, maxValue, false)
	}
}

func runCoarseTest(numGoroutines, numNodes, maxValue int, random bool) {
	ll := NewCoarseGrainedList()
	var wg sync.WaitGroup
	nodesPerGoroutine := numNodes / numGoroutines
	
	start := time.Now()
	for g := 0; g < numGoroutines; g++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < nodesPerGoroutine; i++ {
				var key int
				if random {
					key = rand.Intn(maxValue)
				} else {
					key = id*nodesPerGoroutine + i
				}
				ll.Add(key)
			}
		}(g)
	}
	wg.Wait()
	elapsed := time.Since(start)
	
	fmt.Printf("Goroutines: %3d, Time: %v, Nodes: %d\n", 
		numGoroutines, elapsed, ll.Count())
}

func runFineGrainedExperiment() {
	fmt.Println("\n--- A.2.3: Fine-Grained Locking ---")
	
	numCores := runtime.NumCPU()
	numNodes := 10000 // Reduced for faster testing
	maxValue := 1000000
	
	goroutineCounts := []int{1, 2, 4, 8, 16, 32}
	if numCores > 1 {
		maxG := 32 * numCores
		if maxG > 32 {
			goroutineCounts = append(goroutineCounts, 64, 128)
		}
	}
	
	// Random values
	fmt.Println("\nRandom values:")
	for _, g := range goroutineCounts {
		runFineTest(g, numNodes, maxValue, true)
	}
	
	// Sequential values
	fmt.Println("\nSequential values:")
	for _, g := range goroutineCounts {
		runFineTest(g, numNodes, maxValue, false)
	}
}

func runFineTest(numGoroutines, numNodes, maxValue int, random bool) {
	ll := NewFineGrainedList()
	var wg sync.WaitGroup
	nodesPerGoroutine := numNodes / numGoroutines
	
	start := time.Now()
	for g := 0; g < numGoroutines; g++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < nodesPerGoroutine; i++ {
				var key int
				if random {
					key = rand.Intn(maxValue)
				} else {
					key = id*nodesPerGoroutine + i
				}
				ll.Add(key)
			}
		}(g)
	}
	wg.Wait()
	elapsed := time.Since(start)
	
	fmt.Printf("Goroutines: %3d, Time: %v, Nodes: %d\n", 
		numGoroutines, elapsed, ll.Count())
}

func runOptimisticExperiment() {
	fmt.Println("\n--- A.3: Optimistic Locking ---")
	
	numCores := runtime.NumCPU()
	numNodes := 10000 // Reduced for faster testing
	maxValue := 1000000
	
	goroutineCounts := []int{1, 2, 4, 8, 16, 32}
	if numCores > 1 {
		maxG := 32 * numCores
		if maxG > 32 {
			goroutineCounts = append(goroutineCounts, 64, 128)
		}
	}
	
	// With locking
	fmt.Println("\nWith locking enabled - Random values:")
	for _, g := range goroutineCounts {
		runOptimisticTest(g, numNodes, maxValue, true, true)
	}
	
	fmt.Println("\nWith locking enabled - Sequential values:")
	for _, g := range goroutineCounts {
		runOptimisticTest(g, numNodes, maxValue, false, true)
	}
	
	// Without locking (demonstration)
	fmt.Println("\nWith locking disabled - Random values (demonstration with race conditions):")
	runOptimisticTest(4, 1000, maxValue, true, false)
}

func runOptimisticTest(numGoroutines, numNodes, maxValue int, random, useLocking bool) {
	ll := NewOptimisticList(useLocking)
	var wg sync.WaitGroup
	nodesPerGoroutine := numNodes / numGoroutines
	
	start := time.Now()
	for g := 0; g < numGoroutines; g++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < nodesPerGoroutine; i++ {
				var key int
				if random {
					key = rand.Intn(maxValue)
				} else {
					key = id*nodesPerGoroutine + i
				}
				ll.Add(key)
			}
		}(g)
	}
	wg.Wait()
	elapsed := time.Since(start)
	
	lockStatus := "enabled"
	if !useLocking {
		lockStatus = "disabled"
	}
	fmt.Printf("Goroutines: %3d, Time: %v, Nodes: %d (locking: %s)\n", 
		numGoroutines, elapsed, ll.Count(), lockStatus)
}

func P1() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("Part A: Locking Strategies for Linked Lists")
	fmt.Printf("Machine: %d cores\n", runtime.NumCPU())
	
	runSequentialExperiment()
	runCoarseGrainedExperiment()
	runFineGrainedExperiment()
	runOptimisticExperiment()
}
