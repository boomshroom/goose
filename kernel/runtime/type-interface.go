package runtime

import "unsafe"

type Interface struct {
	// Pointer to array of methods beginning with typeDesc,
	//actual methods are unnesesary at the moment.
	Methods **typeDesc
	Obj     unsafe.Pointer
}

func TypeHashInterface(val *Interface, _ uintptr) uintptr {
	if val.Methods == nil {
		return 0
	}
	desc := *val.Methods
	size := desc.size
	f := &desc.hashfn
	fn := *(*func(unsafe.Pointer, uintptr) uintptr)(unsafe.Pointer(&f))
	if isPointer(desc) {
		return fn(unsafe.Pointer(&val.Obj), size)
	} else {
		return fn(val.Obj, size)
	}
}

func TypeEqualInterface(v1, v2 *Interface, _ uintptr) bool {
	if v1.Methods == nil || v2.Methods == nil {
		return v1.Methods == v2.Methods
	}
	v1Desc := *v1.Methods
	v2Desc := *v2.Methods
	if !typeDescEqual(v1Desc, v2Desc) {
		return false
	}
	if isPointer(v1Desc) {
		return v1.Obj == v2.Obj
	} else {
		f := &v1Desc.equalfn
		fn := *(*func(unsafe.Pointer, unsafe.Pointer, uintptr) bool)(unsafe.Pointer(&f))
		return fn(v1.Obj, v2.Obj, v1Desc.size)
	}
}

func isPointer(t *typeDesc) bool {
	return t.code&0x1f == 22 || t.code&0x1f == 26
}

func typeDescEqual(td1, td2 *typeDesc) bool {
	if td1 == td2 {
		return true
	}
	if td1 == nil || td2 == nil {
		return false
	}
	if td1.code != td2.code || td1.hash != td2.hash {
		return false
	}
	if td1.uncommon != nil && td1.uncommon.name != nil {
		if td2.uncommon == nil || td2.uncommon.name == nil {
			return false
		}
		return PtrStringsEqual(td1.uncommon.name, td2.uncommon.name) && PtrStringsEqual(td1.uncommon.pkg, td2.uncommon.pkg)
	}
	if td2.uncommon != nil && td2.uncommon.name != nil {
		return false
	}
	return PtrStringsEqual(td1.reflection, td2.reflection)
}
