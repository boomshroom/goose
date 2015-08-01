package segment

type Seg64 uint64
type Seg128 [2]uint64

func Compose(s1, s2 Seg64)Seg128{
	return Seg128{uint64(s1), uint64(s2)}
}
func (s Seg128)Decompose()(s1, s2 Seg64){
	return Seg64(s[0]), Seg64(s[1])
}

type SegType uint8

const(
	LDT SegType = 0x2
	TSSAvail = 0x9
	TSSBusy = 0xB
	Call = 0xC
	Interupt = 0xE
	Trap = 0xF
)

type TablePtrPacked [5]uint16
type TablePtr struct{
	Size uintptr
	Ptr uintptr
}
func (p TablePtr)Pack()(ret TablePtrPacked){
	ret[0] = uint16(p.Size)
	ret[1] = uint16(p.Ptr)
	ret[2] = uint16(p.Ptr >> 16)
	ret[3] = uint16(p.Ptr >> 32)
	ret[4] = uint16(p.Ptr >> 48)
	return
}

type CodeDataDesc struct{
	Code, User bool
}

func (d CodeDataDesc)Pack()(ret Seg64){
	ret = (1 << 47) | (1 << 44)
	if d.Code{
		ret |= (1 << 43) | (1 << 53)
	}
	if d.User{
		ret |= 3 << 45
	}
	return 
}

type SystemDesc struct{
	Base uint64
	Limit uint32
	Type SegType
	Granularity, User bool
}

func (d SystemDesc)Pack()(ret Seg128){
	ret[0] = (1 << 47) | ((uint64(d.Type) & 0xF) << 40)
	if d.User{
		ret[0] |= 3 << 45
	}
	if d.Granularity {
		ret[0] |= 1 << 55
	}
	ret[1] = (d.Base >> 32 ) & 0xFFFFFFFF
	ret[0] |= uint64(d.Limit) & 0xFFFF
	ret[0] |= (d.Base & 0xFFFFFF) << 16
	ret[0] |= (uint64(d.Limit) & 0xF0000) << 32
	ret[0] |= (d.Base & 0xFF000000) << 32
	return 
}

type GateDesc struct{
	IST uint8
	Type SegType
	Offset uintptr
	Selector uint16
	User bool
}

func (d GateDesc)Pack()(ret Seg128){
	ret[0] = (1 << 47) | ((uint64(d.Type) & 0xF) << 40)
	if d.User{
		ret[0] |= 3 << 45
	}
	ret[1] = (uint64(d.Offset) >> 32 ) & 0xFFFFFFFF
	ret[0] |= uint64(d.Offset) & 0xFFFF
	ret[0] |= uint64(d.Selector) << 16
	ret[0] |= (uint64(d.Offset) & 0xFFFF0000) << 32
	if d.Type == Interupt || d.Type == Trap{
		ret[0] |= (uint64(d.IST) & 7) << 32
	}
	return 
}

type TSSPacked [27]uint32

type TSS struct{
	IOMapOffset uint16
	RSP [3]uint64
	IST [7]uint64
}

func (s TSSPacked)Unpack()(ret TSS){
	for i := 0; i < 3; i++ {
		ret.RSP[i] = uint64(s[i*2 + 1]) | uint64(s[i*2 + 2]) << 32
	}
	for i := 0; i < 7; i++ {
		ret.RSP[i] = uint64(s[i*2 + 9]) | uint64(s[i*2 + 10]) << 32
	}
	ret.IOMapOffset = uint16(s[26] >> 48)
	return
}

func (s TSS)Pack()(ret TSSPacked){
	for i := 0; i < 3; i++ {
		ret[i*2 + 1] = uint32(s.RSP[i])>>32
		ret[i*2 + 2] = uint32(s.RSP[i])
	}
	for i := 0; i < 7; i++ {
		ret[i*2 + 9] = uint32(s.IST[i])>>32
		ret[i*2 + 10] = uint32(s.IST[i])
	}
	ret[26] = uint32(s.IOMapOffset) << 16
	return
}

type SelectorPacked uint16
type Selector struct{
	Index uint16
	User bool
}

func (s Selector)Pack() (ret SelectorPacked){
	ret = SelectorPacked(s.Index) << 3
	if s.User{
		ret |= 3
	}
	return
}