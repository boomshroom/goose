package tables

type ACPI struct{} // TODO

// The information table structure defined by Pure64 to reside at 0x5000
type InfoTable struct{
	ACPI *ACPI
	BSP uint32
	_ uint32

	CPUSpeed uint16
	CoresActive uint16
	CoresDetect uint16
	_ uint64

	RAMAmount uint32
	_ uint64

	IOAPICCount uint8
	_ uint64

	HPET uint64
	_ uint64

	LAPIC uint64
	IOAPIC uint64

	VideoBase uint64
	VideoX, VideoY uint16
	VideoDepth uint8

	APICID uint8
}