package ipam

// Interface manages the allocation of items out of a range. Interface
// should be threadsafe.
type Interface interface {
	Allocate(int) (bool, error)
	AllocateNext() (int, bool, error)
	Release(int) error
	ForEach(func(int))

	// For testing
	Has(int) bool

	// For testing
	Free() int
}

type AllocatorFactory func(max int, rangeSpec string) (Interface, error)
