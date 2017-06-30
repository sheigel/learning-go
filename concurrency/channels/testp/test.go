package testp


type Val int

func (i Val) True() bool {
	return i == 0
}

