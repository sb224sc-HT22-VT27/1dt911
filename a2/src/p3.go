package src

import (
	"fmt"
	"runtime"
)

func P3() {
	fmt.Println("Assignment 2 Summary")
	fmt.Println("====================")
	fmt.Printf("Machine Info:\n")
	fmt.Printf("  OS: %s\n", runtime.GOOS)
	fmt.Printf("  Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("  CPU Cores: %d\n", runtime.NumCPU())
	fmt.Printf("  Go Version: %s\n", runtime.Version())
	fmt.Println("\nImplemented Problems:")
	fmt.Println("  - Part A: Locking Strategies (Sequential, Coarse-grained, Fine-grained, Optimistic)")
	fmt.Println("  - Part B: Concurrent Quicksort using channels")
	fmt.Println("\nFor detailed experiments and diagrams, refer to the report.")
}
