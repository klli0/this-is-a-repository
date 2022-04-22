/*
 * @Author: liziwei01
 * @Date: 2022-03-03 19:52:07
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-23 00:39:03
 * @Description: 切片工具
 */
package utils

import (
	"fmt"
	"reflect"
)

// 获取sptr的第一个值, 若sptr为空, 则取defau.
func (u *USlice) GetFirstDefault(sptr interface{}, defau ...interface{}) interface{} {
	s := reflect.ValueOf(sptr)
	// if s is a ptr to a slice or an array, get the val it points to
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}
	// if s is not a normal slice or an array, stop
	if s.Kind() != reflect.Slice && s.Kind() != reflect.Array {
		return nil
	}
	if s.Len() != 0 {
		return s.Index(0).Interface()
	}
	if len(defau) != 0 {
		return defau[0]
	}
	return nil
}

// 切片去重.
// sptr: 切片指针.
func (u *USlice) DeDuplicate(sptr interface{}) {
	s := reflect.ValueOf(sptr)
	// if s is a ptr to a slice or an array, get the val it points to
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}
	// if s is not a normal slice or an array, stop
	if s.Kind() != reflect.Slice && s.Kind() != reflect.Array {
		return
	}
	var newArr = reflect.New(s.Type()).Elem()
	for i := 0; i != s.Len(); i++ {
		repeat := false
		for j := i + 1; j != s.Len(); j++ {
			if reflect.DeepEqual(s.Index(i).Interface(), s.Index(j).Interface()) {
				repeat = true
				break
			}
		}
		// find last distinct non-repetitive element and store
		if !repeat {
			newArr = reflect.Append(newArr, s.Index(i))
		}
	}
	// make sptr points to newArr
	s.Set(newArr)
}

// 翻转切片.
// sptr: 切片指针.
func (u *USlice) Reverse(sptr interface{}) {
	s := reflect.ValueOf(sptr)
	// if s is a ptr to a slice or an array, get the val it points to
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}
	// if s is not a normal slice or an array, stop
	if s.Kind() != reflect.Slice && s.Kind() != reflect.Array {
		return
	}
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		x, y := s.Index(i).Interface(), s.Index(j).Interface()
		s.Index(i).Set(reflect.ValueOf(y))
		s.Index(j).Set(reflect.ValueOf(x))
	}
}

// 元素i在切片sptr中是否存在.
// sptr: 切片或切片指针.
func (u *USlice) In(i interface{}, sptr interface{}) bool {
	if u.Pos(i, sptr) != -1 {
		return true
	}
	return false
}

// 寻找特定元素在Slice中的位置,从左到右返回最先找到的位置.
// beg: 从slice的哪里开始找, 可以留空, 默认0.
// sptr: 切片或切片指针.
func (u *USlice) Pos(ele interface{}, sptr interface{}, beg ...int) int {
	var begin int = 0
	if len(beg) != 0 {
		begin = beg[0]
	}
	s := reflect.ValueOf(sptr)
	// if s is a ptr to a slice or an array, get the val it points to.
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}
	// if s is not a normal slice or an array, stop.
	if s.Kind() != reflect.Slice && s.Kind() != reflect.Array {
		return -1
	}
	for i := begin; i < s.Len(); i++ {
		if s.Index(i).Interface() == ele {
			return i
		}
	}
	return -1
}

func (u *USlice) Remove(sptr interface{}, ele interface{}) error {
	return u.RemoveAt(sptr, u.Pos(ele, sptr))
}

// 删除Slice中特定位置i的元素.
// sptr: 切片指针.
func (u *USlice) RemoveAt(sptr interface{}, idx int) error {
	if ps, ok := sptr.(*[]bool); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]uint); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]uint8); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]uint16); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]uint32); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]uint64); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]int); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]int8); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]int16); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]int32); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]int64); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]float32); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]float64); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]complex64); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]complex128); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]byte); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]string); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]rune); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]uintptr); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else if ps, ok := sptr.(*[]interface{}); ok {
		*ps = append((*ps)[:idx], (*ps)[idx+1:]...)
	} else {
		return fmt.Errorf("[Slice]:SliceRemove unsupported type: %T\n", sptr)
	}
	return nil
}
