package runtime

func Copy(dest, src *Array, count int)*Array{
	return MemMove(dest, src, count)
}