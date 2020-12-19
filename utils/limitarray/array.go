package limitarray

func (array *Limitarray)Insert(obj interface{}) bool {
	if array.tail == array.size {
		return false
	}
	array.tail += 1
	array.arr[array.tail] = obj
	return true
}

// 将 array 数据全部取出
func (array *Limitarray) Flush() []interface{}{
	var dst []interface{}
	copy(array.arr,dst)
	array.tail = -1
	return dst
}