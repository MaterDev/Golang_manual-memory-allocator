package main

import (
	"fmt"
	"golang_manual-memory-allocator/allocator"
	"log"
	"math/rand"
	// "golang_manual-memory-allocator/allocator"
	// "golang_manual-memory-allocator/visualizer"
)

// Constant Values
const (
	// Total size of simulated memory pool, in bytes
	memorySize = 1024
	// Number of allocation operations to perform
	numOperations = 50
)

func main() {
	// This variable will keep track of all the addresses that have been allocated. Allocate will add, deallocate will remove.
	var allocatedAddresses []int

	// Create a new memory allocator
	mem := allocator.NewMemoryAllocator(memorySize)

	// Perform random allocations and deallocations
		// 30% chance of deallocation after each operation
		// Visualize the memory sate after each operation.
	for i := 0; i < numOperations; i++ {
		if rand.Float32() < 0.7 { // 70% chance allocation
			size := rand.Intn(100) + 1 // Allocate between 1 and 100 bytes
			addr, err := mem.Allocate(size)
			
			if err != nil {
				// If error is not nil, log an error message
				log.Printf("Allocation failed: %v", err)
			} else {
				// Else print to standard logger, how many bytes (size) and at what address it is stored.
				log.Printf("Allocated %d bytes at address %d", size, addr)
				allocatedAddresses = append(allocatedAddresses, addr)
			}
		} else { // 30% chance deallocation
			// Generate a random index from allocated addresses
			index := rand.Intn(len(allocatedAddresses))
			addr := allocatedAddresses[index]

			// Use random address to deallocate from the memoryblock
			err := mem.Deallocate(addr)

			if err != nil {
				log.Printf("Deallocation failed: %v", err)
			} else {
				log.Printf("Deallocated memory at address %d", addr)
				// remove address from the list
				allocatedAddresses = append(allocatedAddresses[:index], allocatedAddresses[index+1:]...)
			}
		}

	
		// TODO: Visualize the memory state after each operation
	}

	// Perform garbage collection.
	collected := mem.GarbageCollect()
	log.Printf("Garbage collection freed %d bytes", collected)
	
	// Final Memory Visualization
	fmt.Println("Final Memory State: 📈")
}