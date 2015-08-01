package regs

type Regs struct{
	GS ,FS, ES, DS uint32
	EDI, ESI, EBP, ESP, EBX, EDX, ECX, EAX uint32
	IntNo, ErrCode uint32
	EIP, CS, EFlags, UserESP, SS uint32
}