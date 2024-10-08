// make vs new
// a := make([]int, 10)
// b := make([]int, 0, 10) >> make(type, length, capacity).
// a := make([]int, 10, 10)

// Why we use context:


// nil vs empty slice
> when we use a := make([]int, 0(l), 10(c)): It will initialise a slice with 0 elements.
> When we use a := new([]int) it will return a pointer to slice []int, with value as nil. In go pointers can have nil value type
