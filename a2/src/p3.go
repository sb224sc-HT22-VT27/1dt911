package src

import (
	"fmt"
	"runtime"
)

func P3() {
	fmt.Println("Assignment 2 Summary")
	fmt.Println("\nMachine Info:")
	fmt.Printf("\tOS: %s\n", runtime.GOOS)
	fmt.Printf("\tArchitecture: %s\n", runtime.GOARCH)
	fmt.Printf("\tCPU Cores: %d\n", runtime.NumCPU())
	fmt.Printf("\tGo Version: %s\n", runtime.Version())
	fmt.Println("\nImplemented Problems:")
	fmt.Println("\t- Part A: Locking Strategies (Sequential, Coarse-grained, Fine-grained, Optimistic)")
	fmt.Println("\t- Part B: Concurrent Quicksort using channels")
}
