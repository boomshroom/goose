package main

var mem [32]int

func main(){
	 mem[0]=1
	 i:=Malloc(8)
	 Malloc(2)
	 Free(i)
}

func Malloc(size int)int{
	if size==0{
		//panic("Allocating size 0")
		return 0
	}
	freeMark:=0
	free:=mem[freeMark]
	i:=free
	setFreeMark:=true
	//fmt.Println(free)
	for ;;i++{
		if mem[i] != 0 && i == size+free-1{
			mem[freeMark]=mem[i]
			mem[i]=0
			setFreeMark=false
		}else if mem[i] != 0{
			freeMark=i
			free = mem[i]
			i = free
		}else if i==size+free{
			free=i+1
			break
		}
	}
	//fmt.Println(freeMark, free, i, setFreeMark)
	mem[i-size]=size
	if setFreeMark{
		mem[freeMark]=free
	}
	return i-size+1
}

func Free(ptr int){
	if ptr < 2{
		//panic("Atempting to freeing important data") 
		return
	}
	size:= mem[ptr-1]
	
	freeMark:=0
	for{
		free:=mem[freeMark]
		//fmt.Println(free, ptr, size)
		if free==ptr || free==ptr-1{
			//fmt.Println(free, ptr)
			//panic("Free and ptr the same!")
			return
		}
		if freeMark==ptr-2{
			temp:= freeMark
			freeMark=free
			free=mem[freeMark]
			//fmt.Println(freeMark)
			if freeMark==ptr+size{
				mem[temp]=0
			}
			break
		}else if free<ptr{
			temp:=free
			for ;mem[temp]==0;temp++ {
			}
			freeMark=temp
			//fmt.Println("Temp: ",temp)
		}else{
			mem[ptr+size-1]=free
			mem[freeMark]=ptr-1
			break
		}
	}
	
	for i:=ptr-1;i<ptr+size-1;i++{
		mem[i]=0
	}
}