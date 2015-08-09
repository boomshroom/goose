package runtime

import "unsafe"

type typeDesc struct{
	code uint8
	align uint8
	fieldAlign uint8
	size uintptr
	hash uint32
	hashfn uintptr
	equalfn uintptr
	gc *[4]uintptr
	reflection *string
	uncommon uintptr
	ptrToThis *typeDesc
	zero *uint64
}

type fieldAlign struct{
	c byte
	p *struct{}
}

var reflection = "unsafe.Pointer"
var zero uint64 = 0
var gc = [...]uintptr{unsafe.Sizeof(&struct{}{}), 2, 0, 0}

var UnsafePointerDesc = typeDesc{
	57,
	uint8(unsafe.Alignof(&struct{}{})),
	uint8(unsafe.Offsetof(fieldAlign{}.p) - 1),
	unsafe.Sizeof(&struct{}{}),
	78501163,
	0, // calculated at runtime because of go's func handling
	0, // calculated at runtime because of go's func handling
	&gc,
	&reflection,
	0,
	nil,
	&zero,
}

//extern _get_ptr_desc
func getUnsafePointerDesc()*typeDesc

func init(){
	hashDummy := TypeHashIdentity
	equalDummy := TypeEqualIdentity
	*getUnsafePointerDesc() = typeDesc{
		57,
		uint8(unsafe.Alignof(&struct{}{})),
		uint8(unsafe.Offsetof(fieldAlign{}.p) - 1),
		unsafe.Sizeof(&struct{}{}),
		78501163,
		**(**uintptr)(unsafe.Pointer(&hashDummy)),
		**(**uintptr)(unsafe.Pointer(&equalDummy)),
		&gc,
		&reflection,
		0,
		nil,
		&zero,
	}
}