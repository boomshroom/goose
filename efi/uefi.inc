;*********************************************************************
;*                                                                   *
;*          UEFI library for fasm by bzt, Public Domain              *
;*                                                                   *
;*********************************************************************
 
; include x86asm.net's efi.inc
;include 'efi.inc'
struc int8 {
  . db ?
}
struc int16 {
  align 2
  . dw ?
}
struc int32 {
  align 4
  . dd ?
}
struc int64 {
  align 8
  . dq ?
}
struc intn {
  align 8
  . dq ?
}
struc dptr {
  align 8
  . dq ?
}
 
;symbols
 
EFIERR = 0x8000000000000000
EFI_SUCCESS			= 0
EFI_LOAD_ERROR			= EFIERR or 1
EFI_INVALID_PARAMETER		= EFIERR or 2
EFI_UNSUPPORTED 		= EFIERR or 3
EFI_BAD_BUFFER_SIZE		= EFIERR or 4
EFI_BUFFER_TOO_SMALL		= EFIERR or 5
EFI_NOT_READY			= EFIERR or 6
EFI_DEVICE_ERROR		= EFIERR or 7
EFI_WRITE_PROTECTED		= EFIERR or 8
EFI_OUT_OF_RESOURCES		= EFIERR or 9
EFI_VOLUME_CORRUPTED		= EFIERR or 10
EFI_VOLUME_FULL 		= EFIERR or 11
EFI_NO_MEDIA			= EFIERR or 12
EFI_MEDIA_CHANGED		= EFIERR or 13
EFI_NOT_FOUND			= EFIERR or 14
EFI_ACCESS_DENIED		= EFIERR or 15
EFI_NO_RESPONSE 		= EFIERR or 16
EFI_NO_MAPPING			= EFIERR or 17
EFI_TIMEOUT			= EFIERR or 18
EFI_NOT_STARTED 		= EFIERR or 19
EFI_ALREADY_STARTED		= EFIERR or 20
EFI_ABORTED			= EFIERR or 21
EFI_ICMP_ERROR			= EFIERR or 22
EFI_TFTP_ERROR			= EFIERR or 23
EFI_PROTOCOL_ERROR		= EFIERR or 24
 
;helper macro for definition of relative structure member offsets
 
macro struct name
{
  virtual at 0
    name name
  end virtual
}
 
;structures
EFI_SYSTEM_TABLE_SIGNATURE	equ	49h,42h,49h,20h,53h,59h,53h,54h
struc EFI_TABLE_HEADER {
 .Signature    		int64
 .Revision     		int32
 .HeaderSize   		int32
 .CRC32        		int32
 .Reserved     		int32
}
struct EFI_TABLE_HEADER
 
struc EFI_SYSTEM_TABLE {
 .Hdr		        EFI_TABLE_HEADER
 .FirmwareVendor        dptr
 .FirmwareRevision      int32
 .ConsoleInHandle       dptr
 .ConIn 	        dptr
 .ConsoleOutHandle      dptr
 .ConOut	        dptr
 .StandardErrorHandle   dptr
 .StdErr	        dptr
 .RuntimeServices       dptr
 .BootServices	        dptr
 .NumberOfTableEntries  intn
 .ConfigurationTable    dptr
}
struct EFI_SYSTEM_TABLE
 
struc SIMPLE_TEXT_OUTPUT_INTERFACE {
 .Reset 	    	dptr
 .OutputString	    	dptr
 .TestString	    	dptr
 .QueryMode	    	dptr
 .SetMode	    	dptr
 .SetAttribute	    	dptr
 .ClearScreen	    	dptr
 .SetCursorPosition 	dptr
 .EnableCursor	    	dptr
 .Mode		    	dptr
}
struct SIMPLE_TEXT_OUTPUT_INTERFACE
 
;---include ends
 
struc SIMPLE_INPUT_INTERFACE {
 .Reset			dptr
 .ReadKeyStroke		dptr
 .WaitForKey		dptr
}
struct SIMPLE_INPUT_INTERFACE
 
struc EFI_BOOT_SERVICES_TABLE {
 .Hdr		       	EFI_TABLE_HEADER
 .RaisePriority		dptr
 .RestorePriority	dptr
 .AllocatePages		dptr
 .FreePages		dptr
 .GetMemoryMap		dptr
 .AllocatePool		dptr
 .FreePool		dptr
 .CreateEvent		dptr
 .SetTimer		dptr
 .WaitForEvent		dptr
 .SignalEvent		dptr
 .CloseEvent		dptr
 .CheckEvent		dptr
 .InstallProtocolInterface dptr
 .ReInstallProtocolInterface dptr
 .UnInstallProtocolInterface dptr
 .HandleProtocol	dptr
 .Void			dptr
 .RegisterProtocolNotify dptr
 .LocateHandle		dptr
 .LocateDevicePath	dptr
 .InstallConfigurationTable dptr
 .ImageLoad		dptr
 .ImageStart		dptr
 .Exit			dptr
 .ImageUnLoad		dptr
 .ExitBootServices	dptr
 .GetNextMonotonicCount	dptr
 .Stall			dptr
 .SetWatchdogTimer	dptr
 .ConnectController	dptr
 .DisConnectController	dptr
 .OpenProtocol		dptr
 .CloseProtocol		dptr
 .OpenProtocolInformation dptr
 .ProtocolsPerHandle	dptr
 .LocateHandleBuffer	dptr
 .LocateProtocol	dptr
 .InstallMultipleProtocolInterfaces dptr
 .UnInstallMultipleProtocolInterfaces dptr
 .CalculateCrc32	dptr
 .CopyMem		dptr
 .SetMem		dptr
}
struct EFI_BOOT_SERVICES_TABLE
 
struc EFI_RUNTIME_SERVICES_TABLE {
 .Hdr		       	EFI_TABLE_HEADER
 .GetTime		dptr
 .SetTime		dptr
 .GetWakeUpTime		dptr
 .SetWakeUpTime		dptr
 .SetVirtualAddressMap	dptr
 .ConvertPointer	dptr
 .GetVariable		dptr
 .GetNextVariableName	dptr
 .SetVariable		dptr
 .GetNextHighMonoCount	dptr
 .ResetSystem		dptr
}
struct EFI_RUNTIME_SERVICES_TABLE
 
struc EFI_TIME {
 .Year			int16
 .Month			int8
 .Day			int8
 .Hour			int8
 .Minute		int8
 .Second		int8
 .Pad1			int8
 .Nanosecond		int32
 .TimeZone		int16
 .Daylight		int8
 .Pad2			int8
 .sizeof		rb 1
}
struct EFI_TIME
 
EFI_LOADED_IMAGE_PROTOCOL_UUID equ 0A1h,31h,1bh,5bh,62h,95h,0d2h,11h,8Eh,3Fh,0h,0A0h,0C9h,69h,72h,3Bh
struc EFI_LOADED_IMAGE_PROTOCOL {
 .Revision		int32
 .ParentHandle		int64
 .SystemTable		dptr
 .DeviceHandle		int64
 .FilePath		dptr
 .Reserved		int64
 .LoadOptionsSize	int32
 .ImageBase		dptr
 .ImageSize		int64
 .ImageCodeType		int32
 .ImageDataType		int32
 .UnLoad		dptr
}
struct EFI_LOADED_IMAGE_PROTOCOL
 
EFI_BLOCK_IO_PROTOCOL_UUID equ 21h,5bh,4eh,96h,59h,64h,0d2h,11h,8eh,39h,00h,0a0h,0c9h,69h,72h,3bh
struc EFI_BLOCK_IO_PROTOCOL {
 .Revision		int64
 .Media			dptr
 .Reset			dptr
 .ReadBlocks		dptr
 .WriteBlocks		dptr
 .FlushBlocks		dptr
}
struct EFI_BLOCK_IO_PROTOCOL
 
struc EFI_BLOCK_IO_MEDIA {
 .MediaId		int32
 .RemovableMedia	int8
 .MediaPresent		int8
 .LogicalPartition	int8
 .ReadOnly		int8
 .WriteCaching		int8
 .BlockSize		int32
 .IoAlign		int32
 .LastBlock		int64
}
struct EFI_BLOCK_IO_MEDIA
 
EFI_GRAPHICS_OUTPUT_PROTOCOL_UUID equ 0deh, 0a9h, 42h,90h,0dch,023h,38h,04ah,96h,0fbh,7ah,0deh,0d0h,80h,51h,6ah
struc EFI_GRAPHICS_OUTPUT_PROTOCOL {
 .QueryMode		dptr
 .SetMode		dptr
 .Blt			dptr
 .Mode			dptr
}
struct EFI_GRAPHICS_OUTPUT_PROTOCOL
 
struc EFI_GRAPHICS_OUTPUT_PROTOCOL_MODE {
 .MaxMode		int32
 .CurrentMode		int32
 .ModeInfo		dptr
 .SizeOfModeInfo	intn
 .FrameBufferBase	dptr
 .FrameBufferSize	intn
}
struct EFI_GRAPHICS_OUTPUT_PROTOCOL_MODE
 
struc EFI_GRAPHICS_OUTPUT_MODE_INFORMATION {
 .Version		int32
 .HorizontalResolution	int32
 .VerticalResolution	int32
 .PixelFormat		int32
 .RedMask		int32
 .GreenMask		int32
 .BlueMask		int32
 .Reserved		int32
 .PixelsPerScanline	int32
}
struct EFI_GRAPHICS_OUTPUT_MODE_INFORMATION
;---macros to make life easier---
;call it early, after entry point is the best
macro InitializeLib
{
		clc
		or			rdx, rdx
		jz			.badout
		cmp			dword [rdx], 20494249h
		je			@f
.badout:	  xor			rcx, rcx
		xor			rdx, rdx
		stc
@@:		mov			[efi_handler], rcx		; ImageHandle
		mov			[efi_ptr], rdx    		; pointer to SystemTable
}
 
;invoke an UEFI function
macro uefi_call_wrapper			interface,function,arg1,arg2,arg3,arg4,arg5,arg6,arg7,arg8,arg9,arg10,arg11
{
numarg = 0
if ~ arg11 eq
 numarg = numarg + 1
 if ~ arg11 eq rdi
		mov			rdi, arg11
 end if
end if
if ~ arg10 eq
 numarg = numarg + 1
 if ~ arg10 eq rsi
		mov			rsi, arg10
 end if
end if
if ~ arg9 eq
 numarg = numarg + 1
 if ~ arg9 eq r14
		mov			r14, arg9
 end if
end if
if ~ arg8 eq
 numarg = numarg + 1
 if ~ arg8 eq r13
		mov			r13, arg8
 end if
end if
if ~ arg7 eq
 numarg = numarg + 1
 if ~ arg7 eq r12
		mov			r12, arg7
 end if
end if
if ~ arg6 eq
 numarg = numarg + 1
 if ~ arg6 eq r11
		mov			r11, arg6
 end if
end if
if ~ arg5 eq
 numarg = numarg + 1
 if ~ arg5 eq r10
		mov			r10, arg5
 end if
end if
if ~ arg4 eq
 numarg = numarg + 1
 if ~ arg4 eq r9
		mov			r9, arg4
 end if
end if
if ~ arg3 eq
 numarg = numarg + 1
 if ~ arg3 eq r8
		mov			r8, arg3
 end if
end if
if ~ arg2 eq
 numarg = numarg + 1
 if ~ arg2 eq rdx
		mov			rdx, arg2
 end if
end if
if ~ arg1 eq
 numarg = numarg + 1
 if ~ arg1 eq rcx
  if ~ arg1 in <ConsoleInHandle,ConIn,ConsoleOutHandle,ConOut,StandardErrorHandle,StdErr,RuntimeServices,BootServices>
		mov			rcx, arg1
  end if
 end if
end if
		xor			rax, rax
		mov			al, numarg
if interface in <ConsoleInHandle,ConIn,ConsoleOutHandle,ConOut,StandardErrorHandle,StdErr,RuntimeServices,BootServices>
		mov			rbx, [efi_ptr]
		mov			rbx, [rbx + EFI_SYSTEM_TABLE.#interface]
else
 if ~ interface eq rbx
		mov			rbx, interface
 end if
end if
if arg1 in <ConsoleInHandle,ConIn,ConsoleOutHandle,ConOut,StandardErrorHandle,StdErr,RuntimeServices,BootServices>
		mov			rcx, rbx
end if
if defined SIMPLE_INPUT_INTERFACE.#function
		mov			rbx, [rbx + SIMPLE_INPUT_INTERFACE.#function]
else
 if defined SIMPLE_TEXT_OUTPUT_INTERFACE.#function
		mov			rbx, [rbx + SIMPLE_TEXT_OUTPUT_INTERFACE.#function]
 else
  if defined EFI_BOOT_SERVICES_TABLE.#function
		mov			rbx, [rbx + EFI_BOOT_SERVICES_TABLE.#function]
  else
   if defined EFI_RUNTIME_SERVICES_TABLE.#function
		mov			rbx, [rbx + EFI_RUNTIME_SERVICES_TABLE.#function]
   else
    if defined EFI_GRAPHICS_OUTPUT_PROTOCOL.#function
		mov			rbx, [rbx + EFI_GRAPHICS_OUTPUT_PROTOCOL.#function]
    else
     if defined EFI_GRAPHICS_OUTPUT_PROTOCOL_MODE.#function
		mov			rbx, [rbx + EFI_GRAPHICS_OUTPUT_PROTOCOL_MODE.#function]
     else
		mov			rbx, [rbx + function]
     end if
    end if
   end if
  end if
 end if
end if
		call			uefifunc
}
 
;*********************************************************************
;*                       Library functions                           *
;*********************************************************************
 
section '.text' code executable readable
 
uefifunc:	;save stack pointer
		mov			qword [uefi_rsptmp], rsp
		;set up new aligned stack
		and			esp, 0FFFFFFF0h
		;alignment check on arguments
		bt			eax, 0
		jnc			@f
		push			rax
		;arguments
@@:		cmp			al, 11
		jb			@f
		push			rdi
@@:		cmp			al, 10
		jb			@f
		push			rsi
@@:		cmp			al, 9
		jb			@f
		push			r14
@@:		cmp			al, 8
		jb			@f
		push			r13
@@:		cmp			al, 7
		jb			@f
		push			r12
@@:		cmp			al, 6
		jb			@f
		push			r11
@@:		cmp			al, 5
		jb			@f
		push			r10
@@:		;space for
		;r9
		;r8
		;rdx
		;rcx
		sub			rsp, 4*8
		;call function
		call			rbx
		;restore old stack
		mov			rsp, qword [uefi_rsptmp]
		ret
 
section '.data' data readable writeable
efi_handler:	dq			0
efi_ptr:	dq			0
uefi_rsptmp:	dq			0