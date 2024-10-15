# GoLang Manual Memory Allocator (Simulator)

![cover](./images/cover.png)

## Overview

This project is a Go-based simulation of memory allocation and management processes. It provides a hands-on way to understand how computer memory is managed at a low level, including concepts like memory allocation, deallocation, fragmentation, and garbage collection.

## What is Memory Allocation?

Memory allocation is the process of reserving a portion of computer memory for program use. In most modern programming languages, this process is handled automatically by the runtime environment or operating system. However, understanding how it works "under the hood" is crucial for writing efficient and bug-free code.

This simulator demonstrates several key concepts:

1. **Allocation**: Reserving a block of memory for use.
2. **Deallocation**: Freeing up previously allocated memory when it's no longer needed.
3. **Fragmentation**: The phenomenon where free memory is split into small, non-contiguous blocks.
4. **Coalescing**: The process of merging adjacent free blocks to reduce fragmentation.
5. **Garbage Collection**: A simplified simulation of automatic memory recovery.

## Project Structure

```bash
go-manual-memory-allocator/
├── main.go
├── allocator/
│   └── allocator.go
└── visualizer/
    └── visualizer.go
```

## Prerequisites

- Go 1.18 or later

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/go-manual-memory-allocator.git
   cd go-manual-memory-allocator
   ```

2. Ensure you have Go installed on your system.

## Running the Simulator

To run the memory allocation simulator:

```bash
go run main.go
```

This will start the simulation, performing a series of random memory allocations and deallocations. The output will show:

- Allocations: The size and address of each allocated block.
- Deallocations: The address of each deallocated block.
- Memory state visualization after each operation.
- A final garbage collection step and its results.

## Understanding the Output

- `#` represents allocated memory blocks.
- `-` represents free memory blocks.
- Each operation (allocation or deallocation) is followed by a visual representation of the memory state.

Example output:

```bash
Allocated 50 bytes at address 0
[##################################################--------------------------------------------------]
Legend: # = Allocated, - = Free

Deallocated memory at address 0
[--------------------------------------------------##################################################]
Legend: # = Allocated, - = Free
```

## Key Components

1. **MemoryAllocator**: Manages the memory pool and allocation/deallocation operations.
2. **Allocate**: Finds and reserves a free block of sufficient size.
3. **Deallocate**: Marks a previously allocated block as free and attempts to merge adjacent free blocks.
4. **Coalesce**: Merges adjacent free blocks to combat fragmentation.
5. **GarbageCollect**: Simulates a simple garbage collection process.

## Customization

You can modify the following parameters in `main.go` to experiment with different scenarios:

- `memorySize`: Total size of the simulated memory pool.
- `numOperations`: Number of allocation/deallocation operations to perform.
- Allocation probability: Adjust the `0.7` value in `rand.Float32() < 0.7` to change the ratio of allocations to deallocations.
