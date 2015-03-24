ENTRY (loader)

kernel_offset = 0xC0000000;

SECTIONS
{

    . = kernel_offset;
    kernel_start = .;

    .text ALIGN (0x1000): AT(ADDR(.text) - kernel_offset){
	*(.entry)
	*(.text)
	*(.rodata*)
    }

    .data ALIGN (0x1000) : AT(ADDR(.data) - kernel_offset)
    {
        *(.data)
    }

    .bss : AT(ADDR(.bss) - kernel_offset)
    {
        _sbss = .;
        *(COMMON)
        *(.bss)
        _ebss = .;
    }
    
    kernel_end = .;
}