package limitarray

const (
	TAIL = -1
)

func (array *Limitarray)Insert(obj interface{}) bool {
	if array.tail == array.size {
		return false
	}
	array.tail += 1
	array.Arr[array.tail] = obj
	return true
}

// 将 array 数据全部取出, 这里用copy 可能有性能问题
/*
func (array *Limitarray) Flush() []interface{}{
	var dst []interface{}
	copy(array.Arr,dst)
	array.tail = TAIL
	return dst
}
*/


func (array *Limitarray) Reset() {
	array.tail = TAIL
}


func (array *Limitarray) IsFull() bool {
	if array.tail == array.size {
		return true
	}
	return false
}
