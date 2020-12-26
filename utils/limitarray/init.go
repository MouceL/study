package limitarray

type Limitarray struct {
	size int64
	data []interface{}
	tail int64
}

func NewLimitarray(size int64) *Limitarray {
	return &Limitarray{
		size:  size,
		data: make([]interface{},size),
		tail:  -1,
	}
}



func (array *Limitarray)GetData() []interface{}{
	return array.data
}



