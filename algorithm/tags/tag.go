package tags

type Array int

type LUT int

const (
	_                Array = iota
	Scan                   // Scan array
	Partition              // 2-way partition or 3-way partition in quick sort.
	DoublePointer          // Scan array by double pointers.
	CollisionPointer       // Scan array by the collision pointers.
	SlidingWindow          // A sliding window to find the optimal solution.
)

const (
	_ LUT = iota
	LookupTable
)
