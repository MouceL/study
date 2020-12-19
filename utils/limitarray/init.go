package limitarray

type Limitarray struct {
	size int64
	arr []interface{}
	tail int64
}

func New(size int64) *Limitarray {
	return &Limitarray{
		size:  size,
		arr: make([]interface{},size),
		tail:  -1,
	}
}





