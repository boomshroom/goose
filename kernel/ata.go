package ata

import "asm"

// Heavily borrowed from Trinix 
// github.com/Bloodmanovski/Trinix

type Controller struct{
	Base uint32
	Num uint8
	Drives [2]Drive
}
var Controllers [2]Controller = {{Bus1, 0}, {Bus2, 1}}

type Drive struct{
	slave bool
	blocks uint
	data [256]uint16
}

const (
	Bus1 = 0x1F0
	Bus2 = 0x170
)

const (
	Data uint16 = iota
	FeaturesError
	SectCount
	Partial1
	Partial2
	Partial3
	DriveSelect
	Command
)

const (
	Identitfy uint8 = 0xEC
	Read = 0x20
	Write = 0x30
)

func init(){
	(&Controllers[0]).identity(false)
	(&Controllers[0]).identity(true)
	(&Controllers[1]).identity(false)
	(&Controllers[1]).identity(true)
}

func (c *Controller)identity(slave bool){
	if slave {
		asm.OutportB(DriveSelect, 0xB0)
	}else{
		asm.OutportB(DriveSelect, 0xA0)
	}

	asm.OutportB(Command, Identitfy)
	ret := asm.InportB(Command)
	if ret==0{
		return
	}

	for (ret&0x88 != 0x08) && (ret&1 !=1){
		ret = asm.InportB(Command)
	}

	if ret&1==1{
		return
	}
	var drive *Drive
	if slave{
		drive = &c.Drives[1]
	}else{
		drive = &c.Drives[0]
	}

	for i := range drive.data{
		drive.data[i] = asm.Inport16(Data)
	}
	drive.blocks = uint32(drive.data[61])<<16 | uint32(drive.data[60])
	drive.slave = slave
}