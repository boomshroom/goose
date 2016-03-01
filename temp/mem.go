package main

import (
	"fmt"
	//"math/rand"
)

var mem Mem = make(Mem, 64)

type Mem []int

// Hack
func RandMapEntry(m map[int]int)int{
	for i := range m {
		return i
	}
	panic("dead code")
}

func main(){
	mem[0] = 1
	i := mem.Malloc(5)
	//fmt.Println(mem)
	mem.Malloc(3)
	//fmt.Println(mem)
	mem.Free(i)
	//fmt.Println(mem)
	s := mem.Malloc(32)
	//fmt.Println(mem)
	sub := mem[s: s+32]
	sub.Malloc(8)
	fmt.Println(mem)
	mem.Malloc(6)
	fmt.Println(mem)

	/*vars := make(map[int]int)

	for i:=0; i< 4; i++{
		if rand.Intn(2) == 0 && len(vars)!=0{
			v := RandMapEntry(vars)
			size := vars[v]
			Free(v)
			delete(vars, v)
			for j := v; j<v+size; j++{
				mem[j] = 0
			}
			fmt.Println("free ", mem, size, v)
		}else{
			size := rand.Intn(32)
			v := Malloc(size)
			vars[v] = size
			for j := 0; j<size; j++{
				mem[j+v] = j
			}
			fmt.Println("alloc", mem, size, v)
		}
	}*/
}

func (m Mem)Malloc(size int)int{
	if size==0{
		//panic("Allocating size 0")
		return 0
	}
	freeMark:=0
	free:=m[freeMark]
	i:=free
	setFreeMark:=true
	//fmt.Println(free)
	for ;;i++{
		if i>= len(m){
			panic(fmt.Errorf("Not enough space for %v bytes", size))
		}
		if i==size+free{
			//fmt.Println(size, free, i)
			free=i+1
			break
		}else if m[i] != 0 && i == size+free-1{
			m[freeMark]=m[i]
			m[i]=0
			setFreeMark=false
		}else if m[i] != 0{
			if size == 6 {
				fmt.Println(i)
			}
			freeMark=i
			free = m[i]
			i = free
		}
	}
	//fmt.Println(freeMark, free, i, setFreeMark)
	m[i-size]=size
	if setFreeMark{
		m[freeMark]=free
	}
	return i-size+1
}

func (m Mem)Free(ptr int){
	if ptr == 0{
		return
	}else if ptr < 2{
		panic("Atempting to freeing important data")
	} 
	size:= m[ptr-1]
	
	freeMark:=0
	for{
		free:=m[freeMark]
		//fmt.Println(free, ptr, size)
		if free==ptr || free==ptr-1{
			//fmt.Println(free, ptr)
			panic("Free and ptr the same!")
			return
		}
		/*if freeMark==ptr-2{

			temp := freeMark
			freeMark=free
			free=m[freeMark]
			fmt.Println(freeMark)

			if freeMark==ptr+size{
				m[temp]=0
			}

			break
		}else*/ if free<ptr{
			temp:=free
			for ;m[temp]==0;temp++ {
			}
			freeMark=temp
			//fmt.Println("Temp: ",temp)
		}else{
			if free != ptr+size {
				m[ptr+size-1]=free
			}
			m[freeMark]=ptr-1
			break
		}
	}
	
	for i:=ptr-1;i<ptr+size-1;i++{
		mem[i]=0
	}
}