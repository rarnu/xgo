package xgo

import "fmt"

type XSet[T any] []T

func NewSet[T any]() XSet[T] {
	return XSet[T]{}
}

func NewSetWithList[T any](list []T) XSet[T] {
	s := XSet[T]{}
	s.AddAll(list...)
	return s
}

func NewSetWithItems[T comparable](items ...T) XSet[T] {
	s := XSet[T]{}
	s.AddAll(items...)
	return s
}

func SetOf[T comparable](items ...T) XSet[T] {
	s := XSet[T]{}
	s.AddAll(items...)
	return s
}

func EmptySet[T comparable]() XSet[T] {
	return XSet[T]{}
}

func (x XSet[T]) Size() XInt {
	return XInt(len(x))
}

func (x *XSet[T]) Add(item T) error {
	if !x.Contains(item) {
		*x = append(*x, item)
		return nil
	} else {
		return fmt.Errorf("%v already exists in set", item)
	}
}

// AddAll 添加多个元素
func (x *XSet[T]) AddAll(items ...T) {
	for _, item := range items {
		_ = x.Add(item)
	}
}

// Delete 删除指定Key元素
func (x *XSet[T]) Delete(item T) error {
	idx := ListIndexOf(*x, item)
	if idx != -1 {
		*x = append((*x)[:idx], (*x)[idx+1:]...)
		return nil
	} else {
		return fmt.Errorf("%v not exists in set", item)
	}
}

// Contains 判断key是否存在
func (x XSet[T]) Contains(item T) bool {
	return ListContains(x, item)
}

// Clear 重置
func (x *XSet[T]) Clear() {
	*x = XSet[T]{}
}

func (x XSet[T]) ToList() XList[T] {
	return NewListWithList(x)
}
