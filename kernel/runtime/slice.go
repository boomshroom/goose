package runtime

import "unsafe"

type Slice struct {
	Values          unsafe.Pointer
	Count, Capacity int
}
