package allocator

import (
	"bytes"
	"errors"
)

/*
	- NewMemoryAllocator: Creates a new memory allocator with a fixed-size memory pool.
	- Allocate: Finds a free block of memory and allocates it.
	- Deallocate: Frees a previously allocated block of memory.
	- coalesce: Merges adjacent free blocks to reduce fragmentation.
	- GarbageCollect: Simulates a simple garbage collection process.
	- GetBlocks and String: Helper methods for visualization and debugging.
*/

// Struct for block of memory
type MemoryBlock struct {
	Address int
	Size    int
	Free    bool
}

// Struct for MemoryAllocator
type MemoryAllocator struct {
	Memory []byte
	Blocks []*MemoryBlock // A slice of MemoryBlock pointers
}

// Function to creates and returns a pointer to a new MemoryAllocator of a given size
func NewMemoryAllocator(size int) *MemoryAllocator {
	// When & is used in front of a struct literal, it creates the struct and returns a pointer to it in one step.
	return &MemoryAllocator{
		Memory: make([]byte, size),
		Blocks: []*MemoryBlock{{Address: 0, Size: size, Free: true}},
	}
}

// TODO: Method for MemoryAllocator that allocates block of memory of the given size.
func (ma *MemoryAllocator) Allocate(size int) (int, error) {
	for i, block := range ma.Blocks {
		// If block is free and large enough for the size needed.
		if block.Free && block.Size >= size {
			if block.Size > size {
				// Split the block
				newBlock := &MemoryBlock{
					Address: block.Address + size,
					Size:    block.Size - size,
					Free:    true,
				}
				ma.Blocks = append(ma.Blocks[:i+1], append([]*MemoryBlock{newBlock}, ma.Blocks[i+1:]...)...)
				block.Size = size

				block.Free = false
				return block.Address, nil
			}
		}
	}
	return 0, errors.New("no free memory block large enough.")

}

// TODO: Method for MemoryAllocator to free memory blocks at a given address.

// TODO: Method for MemoryAllocator to coalesce (merge) adjacent free blocks.

// TODO: Method for MemoryAllocator to simulate garbage collection process

// -------------- Visualizer

// TODO: Method for MemoryAllocator to return copy of memory blocks for visualziation.

// TODO: Method for MemoryAllocator to provide a string representation of the MemoryAllocator
