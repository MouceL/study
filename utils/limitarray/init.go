package limitarray

type Limitarray struct {
	size int64
	Arr []interface{}
	tail int64
}

func NewLimitarray(size int64) *Limitarray {
	return &Limitarray{
		size:  size,
		Arr: make([]interface{},size),
		tail:  -1,
	}
}





