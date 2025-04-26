package collections

import "errors"

// ?copy, cut, delete, extend, insert, push, pop的实现

// copy 简单的一维切片的复制实现, 覆盖操作
// newSlice和oldSlice的len必须相同，否则报错
func copy(oldSlice []int, newSlice []int) error {
	if oldSlice == nil || newSlice == nil {
		return errors.New("oldSlice or newSlice is nil")
	}

	if len(oldSlice) != len(newSlice) {
		return errors.New("oldSlice length is not equall to newSlice length")
	}

	for i := 0; i < len(oldSlice); i++ {
		newSlice[i] = oldSlice[i]
	}

	return nil
}

// extend 将新切片中的元素追加到旧切片中
func extend(srcSlice []int, newSlice []int) (error, []int) {
	if srcSlice == nil || newSlice == nil {
		return errors.New("slice is nil"), nil
	}

	srcSlice = append(srcSlice, newSlice...)

	return nil, srcSlice
}

// DeleteEleFromSlice 根据值从切片中删除元素(根据元素值), 删除元素后的新切片长度等于旧切片长度
func DeleteEleInSlice[T uint64 | int64 | int | int8 | int16](target T, s []T) []T {
	newS := make([]T, len(s))
	newSIndex := 0

	for _, v := range s {
		if v == target {
			continue
		}
		newS[newSIndex] = v
		newSIndex++
	}

	return newS
}
