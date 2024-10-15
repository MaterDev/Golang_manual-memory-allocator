package visualizer

import (
	"strings"
	"golang_manual-memory-allocator/allocator"
)

// Creates a string representation of the memory state
func VisualizeMemory(ma *allocator.MemoryAllocator) string {
	blocks := ma.GetBlocks()
	
	// Will use string builder to programmatically create a string that can be displayed.
	var sb strings.Builder

	sb.WriteString("Memory State:\n")
	sb.WriteString("[")

	// Programmatically building a string
	// Mapping memory into a string representation
	// For each block will represent as - or #, line by line
	for _, block := range blocks {
		// For either representation, will write as many as match the block size.
		if block.Free {
			sb.WriteString(strings.Repeat("-", block.Size))
		} else {
			sb.WriteString(strings.Repeat("#", block.Size))
		}

	}

	sb.WriteString("]\n")
	sb.WriteString("Legend: # = Allocated, - = Free\n")

	return sb.String()
}