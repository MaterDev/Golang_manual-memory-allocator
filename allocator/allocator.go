package allocator

import (
	"errors"
	"fmt"
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

// Method for MemoryAllocator that allocates block of memory of the given size.
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
	return 0, errors.New("no free memory block large enough")

}

// Method for MemoryAllocator to free memory blocks at a given address.
func (ma *MemoryAllocator) Deallocate(address int) error {
	for i, block := range ma.Blocks {
		// fmt.Printf("In for scope of Deallocate()")
		if block.Address == address {
			if block.Free {
				return errors.New("memory already free")
			}
			// Free the block
			block.Free = true
			ma.coalesce(i)
			return nil
		}

	}
	return errors.New("invalid address")
}

// Method for MemoryAllocator to coalesce (merge) adjacent free blocks (merging into left).
func (ma *MemoryAllocator) coalesce(index int) {

	// ! Step 1
	// Coalesce with next block if free
	// If this index is not the last block AND the block after it is Free
	if index < len(ma.Blocks)-1 && ma.Blocks[index+1].Free {
		// Will add the size of the next block into the index block.
		ma.Blocks[index].Size += ma.Blocks[index+1].Size
		// Remove the next block from the slice:
		// ma.Blocks[:index+1] keeps all blocks upto and including the current index
		// ma.Blocks[index+2...:] will create a new slice with the items from index+2 at the end.
		// the ... after the slice wll unpack these as individual arguments.
		ma.Blocks = append(ma.Blocks[:index+1], ma.Blocks[index+2:]...)
	}

	// ! Step 2
	// Coalesce with previous block if free
	// If this index is after the 0th AND the block before it is free
	if index > 0 && ma.Blocks[index-1].Free {
		// Adds the size of the current index to the previous index
		ma.Blocks[index-1].Size += ma.Blocks[index].Size
		// Will slice up to and not including the current index
		// Will unpack and append all the blocks from after the first index.
		ma.Blocks = append(ma.Blocks[:index], ma.Blocks[index+1:]...)
	}

}

// Method for MemoryAllocator to simulate garbage collection process
func (ma *MemoryAllocator) GarbageCollect() int {
	collected := 0

	// For each block in the MemoryAllocator's slice of []MemoryBlocks,

	// If not free, will check if the memory block at the current address is zeroed out.
	// If zeroed out
	// will set Free to true
	// Will add the block's size to collected
	// Will return total collected.

	for _, block := range ma.Blocks {
		// Check if the block is not-free
		if !block.Free {
			// If not free, will check if the memory block at the current address is zeroed out.
			// (Essentially it's in use or not. This is a simulated process, assuming if the byte is zero, the block is no longer in use.)
			if ma.Memory[block.Address] == 0 {
				block.Free = true
				collected += block.Size
			}

		}
	}
	return collected
}

// -------------- Visualizer

// Method for MemoryAllocator to return copy of memory blocks for visualziation.
func (ma *MemoryAllocator) GetBlocks() []MemoryBlock {
	// Make a new slice of MemoryBlock with a size that is the same the current instance of MemoryAllocator's blocks.
	blocks := make([]MemoryBlock, len(ma.Blocks))

	// For all blocks in the Memory Allocator's blocks
	for i, block := range ma.Blocks {
		// Set the value of blocks current index to a pointer for the corresponding block in MemoryAllocator's MemoryBlock
		blocks[i] = *block // *When you assign a struct to a new variable, you are creating a copy of that struct.
	}

	// Basically, returning dereferenced copy of the original slice.
	return blocks
}

// Method for MemoryAllocator to provide a string representation of the MemoryAllocator
func (ma *MemoryAllocator) String() string {
	var s string
	
	// For each block in the  MA Blocks, 
		// Start with a status of "Allocated"
			// If the block is free, change the status to "Free"
		// After, concatenate a new message to s, which has the information about the current block
	for _, block := range ma.Blocks {
		status := "Allocated"
		if block.Free {
			status = "Free"
		}
		s += fmt.Sprintf("Address: %d, Size: %d, Status: %s\n", block.Address, block.Size, status)
	}

	// Return formatted string representation of MemoryAllocator
	return s
}
