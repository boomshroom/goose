ENTRY (loader)

kernel_offset = 0x100000000;

SECTIONS
{

    . = kernel_offset;
    kernel_start = .;

    .text ALIGN (0x1000): AT(ADDR(.text)){
	*(.entry)
	*(.text)
	*(.rodata*)
    }

    .data ALIGN (0x1000) : AT(ADDR(.data))
    {
        *(.data)
    }

    .bss : AT(ADDR(.bss))
    {
        _sbss = .;
        *(COMMON)
        *(.bss)
        _ebss = .;
    }
    
    kernel_end = .;
}