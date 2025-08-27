package xgo

type XList[T any] []T

func NewList[T any]() XList[T] {
	return []T{}
}

func NewListWithList[T any](list []T) XList[T] {
	return list
}

func NewListWithItems[T any](items ...T) XList[T] {
	return items
}

func ListOf[T any](items ...T) XList[T] {
	return items
}

func EmptyList[T any]() XList[T] {
	return []T{}
}

func (x *XList[T]) Add(item T) XInt {
	idx := len(*x)
	*x = append(*x, item)
	return XInt(idx)
}

func (x *XList[T]) AddAll(item ...T) {
	*x = append(*x, item...)
}

func (x *XList[T]) AddList(list []T) {
	*x = append(*x, list...)
}

func (x *XList[T]) Insert(index XInt, item T) XInt {
	*x = append((*x)[:index], append([]T{item}, (*x)[index:]...)...)
	return index
}

func (x *XList[T]) Delete(index XInt) T {
	item := (*x)[index]
	*x = append((*x)[:index], (*x)[index+1:]...)
	return item
}

func (x *XList[T]) Remove(item T) XInt {
	idx := ListIndexOf(*x, item)
	if idx != -1 {
		x.Delete(XInt(idx))
	}
	return XInt(idx)
}

func (x *XList[T]) Clear() {
	*x = []T{}
}

func (x XList[T]) IsEmpty() bool {
	return len(x) == 0
}

func (x XList[T]) IsNotEmpty() bool {
	return len(x) > 0
}

func (x XList[T]) Size() int {
	return len(x)
}

func (x XList[T]) ForEach(action func(T)) {
	ListForEach(x, action)
}

func (x XList[T]) ForEachIndexed(action func(int, T)) {
	ListForEachIndexed(x, action)
}

func (x XList[T]) Distinct() XList[T] {
	return ListDistinct(x)
}

func (x XList[T]) Filter(predicate func(T) bool) XList[T] {
	return ListFilter(x, predicate)
}

func (x XList[T]) FilterNot(predicate func(T) bool) XList[T] {
	return ListFilterNot(x, predicate)
}

func (x XList[T]) FilterIndexed(predicate func(int, T) bool) XList[T] {
	return ListFilterIndexed(x, predicate)
}

func (x XList[T]) FilterNotIndexed(predicate func(int, T) bool) XList[T] {
	return ListFilterNotIndexed(x, predicate)
}

func (x XList[T]) FilterTo(dest *[]T, predicate func(T) bool) XList[T] {
	return ListFilterTo(x, dest, predicate)
}

func (x XList[T]) FilterNotTo(dest *[]T, predicate func(T) bool) XList[T] {
	return ListFilterNotTo(x, dest, predicate)
}

func (x XList[T]) FilterIndexedTo(dest *[]T, predicate func(int, T) bool) XList[T] {
	return ListFilterIndexedTo(x, dest, predicate)
}

func (x XList[T]) FilterNotIndexedTo(dest *[]T, predicate func(int, T) bool) XList[T] {
	return ListFilterNotIndexedTo(x, dest, predicate)
}

func (x XList[T]) Contains(item T) bool {
	return ListContains(x, item)
}

func (x XList[T]) Find(predicate func(T) bool) (T, error) {
	return ListFind(x, predicate)
}

func (x XList[T]) FindLast(predicate func(T) bool) (T, error) {
	return ListFindLast(x, predicate)
}

func (x XList[T]) First() (T, error) {
	return ListFirst(x)
}

func (x XList[T]) Last() (T, error) {
	return ListLast(x)
}

func (x XList[T]) IndexOf(item T) XInt {
	return XInt(ListIndexOf(x, item))
}

func (x XList[T]) LastIndexOf(item T) XInt {
	return XInt(ListLastIndexOf(x, item))
}

func (x XList[T]) IndexOfCondition(condition func(T) bool) XInt {
	return XInt(ListIndexOfCondition(x, condition))
}

func (x XList[T]) LastIndexOfCondition(condition func(T) bool) XInt {
	return XInt(ListLastIndexOfCondition(x, condition))
}

func (x XList[T]) JoinToString(joint func(T) string) XString {
	return XString(ListJoinToString(x, joint))
}

func (x XList[T]) JoinToStringFull(sep XString, prefix XString, postfix XString, joint func(T) string) XString {
	return XString(ListJoinToStringFull(x, string(sep), string(prefix), string(postfix), joint))
}

func (x XList[T]) All(predicate func(T) bool) bool {
	return ListAll(x, predicate)
}

func (x XList[T]) Any(predicate func(T) bool) bool {
	return ListAny(x, predicate)
}

func (x XList[T]) None(predicate func(T) bool) bool {
	return ListNone(x, predicate)
}

func (x XList[T]) Count(predicate func(T) bool) int {
	return ListCount(x, predicate)
}

func (x XList[T]) SubList(fromIndex XInt, toIndex XInt) XList[T] {
	return SliceSubList(x, int(fromIndex), int(toIndex))
}

func (x XList[T]) Take(n XInt) XList[T] {
	return ListTake(x, int(n))
}

func (x XList[T]) TakeLast(n XInt) XList[T] {
	return ListTakeLast(x, int(n))
}

func (x XList[T]) TakeWhile(n XInt, f func(T) bool) XList[T] {
	return ListTakeWhile(x, int(n), f)
}

func (x XList[T]) TakeLastWhile(n XInt, f func(T) bool) XList[T] {
	return ListTakeLastWhile(x, int(n), f)
}

func (x XList[T]) Drop(n XInt) XList[T] {
	return ListDrop(x, int(n))
}

func (x XList[T]) DropLast(n XInt) XList[T] {
	return ListDropLast(x, int(n))
}

func (x XList[T]) DropWhile(n XInt, f func(T) bool) XList[T] {
	return ListDropWhile(x, int(n), f)
}

func (x XList[T]) DropLastWhile(n XInt, f func(T) bool) XList[T] {
	return ListDropLastWhile(x, int(n), f)
}

func (x XList[T]) Partition(partition XInt) [][]T {
	return ListPartition(x, int(partition))
}

func (x XList[T]) PartitionWithCal(f func(int) int) [][]T {
	return ListPartitionWithCal(x, f)
}

func (x XList[T]) Plus(n []T) XList[T] {
	return ListPlus(x, n)
}

func (x XList[T]) Minus(n []T) XList[T] {
	return ListMinus(x, n)
}

func (x XList[T]) Equals(n []T) bool {
	return ListEquals(x, n)
}

func (x XList[T]) ToSet() XSet[T] {
	return NewSetWithList(x)
}
