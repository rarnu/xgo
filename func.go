package xgo

import (
	"errors"
	"reflect"
)

func ListForEach[T any](list []T, action func(T)) {
	for _, item := range list {
		action(item)
	}
}

func ListForEachIndexed[T any](list []T, action func(int, T)) {
	for i, item := range list {
		action(i, item)
	}
}

func MapForEach[K comparable, V any](m map[K]V, action func(K, V)) {
	for k, v := range m {
		action(k, v)
	}
}

func ListFilter[T any](list []T, predicate func(T) bool) []T {
	var dest []T
	return ListFilterTo(list, &dest, predicate)
}

func ListFilterTo[T any](list []T, dest *[]T, predicate func(T) bool) []T {
	var n []T
	for _, e := range list {
		if predicate(e) {
			*dest = append(*dest, e)
			n = append(n, e)
		}
	}
	return n
}

func ListFilterNot[T any](list []T, predicate func(T) bool) []T {
	var n []T
	return ListFilterNotTo(list, &n, predicate)
}

func ListFilterNotTo[T any](list []T, dest *[]T, predicate func(T) bool) []T {
	var n []T
	for _, e := range list {
		if !predicate(e) {
			*dest = append(*dest, e)
			n = append(n, e)
		}
	}
	return n
}

func ListFilterIndexed[T any](list []T, predicate func(int, T) bool) []T {
	var n []T
	return ListFilterIndexedTo(list, &n, predicate)
}

func ListFilterIndexedTo[T any](list []T, dest *[]T, predicate func(int, T) bool) []T {
	var n []T
	for i, e := range list {
		if predicate(i, e) {
			*dest = append(*dest, e)
			n = append(n, e)
		}
	}
	return n
}

func ListFilterNotIndexed[T any](list []T, predicate func(int, T) bool) []T {
	var n []T
	return ListFilterNotIndexedTo(list, &n, predicate)
}

func ListFilterNotIndexedTo[T any](list []T, dest *[]T, predicate func(int, T) bool) []T {
	var n []T
	for i, e := range list {
		if !predicate(i, e) {
			*dest = append(*dest, e)
			n = append(n, e)
		}
	}
	return n
}

func ListContains[T any](list []T, item T) bool {
	ret := false
	for _, e := range list {
		if reflect.DeepEqual(e, item) {
			ret = true
			break
		}
	}
	return ret
}

func ListDistinct[T any](list []T) []T {
	return SliceDistinct(list)
}

func MapFilter[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V {
	n := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			n[k] = v
		}
	}
	return n
}

func MapFilterNot[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V {
	n := make(map[K]V)
	for k, v := range m {
		if !predicate(k, v) {
			n[k] = v
		}
	}
	return n
}

func MapFilterKeys[K comparable, V any](m map[K]V, predicate func(K) bool) map[K]V {
	n := make(map[K]V)
	for k, v := range m {
		if predicate(k) {
			n[k] = v
		}
	}
	return n
}

func MapFilterValues[K comparable, V any](m map[K]V, predicate func(V) bool) map[K]V {
	n := make(map[K]V)
	for k, v := range m {
		if predicate(v) {
			n[k] = v
		}
	}
	return n
}

func MapFilterTo[K comparable, V any](m map[K]V, dest *map[K]V, predicate func(K, V) bool) map[K]V {
	n := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			(*dest)[k] = v
			n[k] = v
		}
	}
	return n
}

func MapFilterNotTo[K comparable, V any](m map[K]V, dest *map[K]V, predicate func(K, V) bool) map[K]V {
	n := make(map[K]V)
	for k, v := range m {
		if !predicate(k, v) {
			(*dest)[k] = v
			n[k] = v
		}
	}
	return n
}

func MapContains[K comparable, V any](m map[K]V, k K, v V) bool {
	ret := false
	for t, u := range m {
		if t == k && reflect.DeepEqual(u, v) {
			ret = true
			break
		}
	}
	return ret
}

func MapContainsKey[K comparable, V any](m map[K]V, k K) bool {
	ret := false
	for t := range m {
		if t == k {
			ret = true
			break
		}
	}
	return ret
}

func MapContainsValue[K comparable, V any](m map[K]V, v V) bool {
	ret := false
	for _, u := range m {
		if reflect.DeepEqual(u, v) {
			ret = true
			break
		}
	}
	return ret
}

func ListFind[T any](list []T, predicate func(T) bool) (T, error) {
	var n T
	found := false
	for _, e := range list {
		if predicate(e) {
			found = true
			n = e
			break
		}
	}
	if found {
		return n, nil
	} else {
		return n, errors.New("not found")
	}
}

func ListFindLast[T any](list []T, predicate func(T) bool) (T, error) {
	var n T
	found := false
	for i := len(list) - 1; i >= 0; i-- {
		e := list[i]
		if predicate(e) {
			found = true
			n = e
			break
		}
	}
	if found {
		return n, nil
	} else {
		return n, errors.New("not found")
	}
}

func ListFirst[T any](list []T) (T, error) {
	var n T
	if len(list) == 0 {
		return n, errors.New("not found")
	}
	n = list[0]
	return n, nil
}

func ListLast[T any](list []T) (T, error) {
	var n T
	if len(list) == 0 {
		return n, errors.New("not found")
	}
	n = list[len(list)-1]
	return n, nil
}

func ListIndexOf[T any](list []T, item T) int {
	idx := -1
	for i, e := range list {
		if reflect.DeepEqual(e, item) {
			idx = i
			break
		}
	}
	return idx
}

func ListLastIndexOf[T any](list []T, item T) int {
	idx := -1
	for i := len(list) - 1; i >= 0; i-- {
		e := list[i]
		if reflect.DeepEqual(e, item) {
			idx = i
			break
		}
	}
	return idx
}

func ListIndexOfCondition[T any](list []T, condition func(T) bool) int {
	idx := -1
	for i, e := range list {
		if condition(e) {
			idx = i
			break
		}
	}
	return idx
}

func ListLastIndexOfCondition[T any](list []T, condition func(T) bool) int {
	idx := -1
	for i := len(list) - 1; i >= 0; i-- {
		e := list[i]
		if condition(e) {
			idx = i
			break
		}
	}
	return idx
}

func ListJoinToStringFull[T any](list []T, sep string, prefix string, postfix string, joint func(T) string) string {
	buffer := prefix
	var count = 0
	for _, e := range list {
		count++
		if count > 1 {
			buffer += sep
		}
		buffer += joint(e)
	}
	buffer += postfix
	return buffer
}

func ListJoinToString[T any](list []T, joint func(T) string) string {
	return ListJoinToStringFull(list, ",", "", "", joint)
}

func MapJoinToStringFull[K comparable, V any](m map[K]V, sep string, prefix string, postfix string, joint func(K, V) string) string {
	buffer := prefix
	var count = 0
	for k, v := range m {
		count++
		if count > 1 {
			buffer += sep
		}
		buffer += joint(k, v)
	}
	buffer += postfix
	return buffer
}

func MapJoinToString[K comparable, V any](m map[K]V, joint func(K, V) string) string {
	return MapJoinToStringFull(m, ",", "", "", joint)
}

func ListAll[T any](list []T, f func(T) bool) bool {
	for _, e := range list {
		if !f(e) {
			return false
		}
	}
	return true
}

func ListAny[T any](list []T, f func(T) bool) bool {
	for _, e := range list {
		if f(e) {
			return true
		}
	}
	return false
}

func ListNone[T any](list []T, f func(T) bool) bool {
	for _, e := range list {
		if f(e) {
			return false
		}
	}
	return true
}

func ListCount[T any](list []T, f func(T) bool) int {
	num := 0
	for _, e := range list {
		if f(e) {
			num++
		}
	}
	return num
}

func MapAll[K comparable, V any](m map[K]V, f func(K, V) bool) bool {
	for k, v := range m {
		if !f(k, v) {
			return false
		}
	}
	return true
}

func MapAny[K comparable, V any](m map[K]V, f func(K, V) bool) bool {
	for k, v := range m {
		if f(k, v) {
			return true
		}
	}
	return false
}

func MapNone[K comparable, V any](m map[K]V, f func(K, V) bool) bool {
	for k, v := range m {
		if f(k, v) {
			return false
		}
	}
	return true
}

func MapCount[K comparable, V any](m map[K]V, f func(K, V) bool) int {
	num := 0
	for k, v := range m {
		if f(k, v) {
			num++
		}
	}
	return num
}

func MapAllKey[K comparable, V any](m map[K]V, f func(K) bool) bool {
	for k := range m {
		if !f(k) {
			return false
		}
	}
	return true
}

func MapAnyKey[K comparable, V any](m map[K]V, f func(K) bool) bool {
	for k := range m {
		if f(k) {
			return true
		}
	}
	return false
}

func MapNoneKey[K comparable, V any](m map[K]V, f func(K) bool) bool {
	for k := range m {
		if f(k) {
			return false
		}
	}
	return true
}

func MapCountKey[K comparable, V any](m map[K]V, f func(K) bool) int {
	num := 0
	for k := range m {
		if f(k) {
			num++
		}
	}
	return num
}

func MapAllValue[K comparable, V any](m map[K]V, f func(V) bool) bool {
	for _, v := range m {
		if !f(v) {
			return false
		}
	}
	return true
}

func MapAnyValue[K comparable, V any](m map[K]V, f func(V) bool) bool {
	for _, v := range m {
		if f(v) {
			return true
		}
	}
	return false
}

func MapNoneValue[K comparable, V any](m map[K]V, f func(V) bool) bool {
	for _, v := range m {
		if f(v) {
			return false
		}
	}
	return true
}

func MapCountValue[K comparable, V any](m map[K]V, f func(V) bool) int {
	num := 0
	for _, v := range m {
		if f(v) {
			num++
		}
	}
	return num
}

func ListTake[T any](list []T, n int) []T {
	var t []T
	if n <= 0 {
		return t
	}
	if n >= len(list) {
		return list
	}
	return SliceSubList(list, 0, n)
}

func ListTakeLast[T any](list []T, n int) []T {
	var t []T
	if n <= 0 {
		return t
	}
	if n >= len(list) {
		return list
	}
	return SliceSubList(list, len(list)-n, len(list))
}

func ListTakeWhile[T any](list []T, n int, predicate func(T) bool) []T {
	if n <= 0 {
		return nil
	}
	t := make([]T, n)
	for i := 0; i < len(list); i++ {
		if len(t) < n && predicate(list[i]) {
			t = append(t, list[i])
		}
	}
	return t
}

func ListTakeLastWhile[T any](list []T, n int, predicate func(T) bool) []T {
	if n <= 0 {
		return nil
	}
	t := make([]T, n)
	for i := len(list) - 1; i >= 0; i++ {
		if len(t) < n && predicate(list[i]) {
			t = append(t, list[i])
		}
	}
	return t
}

func ListDrop[T any](list []T, n int) []T {
	if n < 0 {
		return list
	}
	return list[n:]
}

func ListDropLast[T any](list []T, n int) []T {
	if n < 0 {
		return nil
	}
	if n == 0 {
		return list
	}
	return list[:n]
}

func ListDropWhile[T any](list []T, n int, predicate func(T) bool) []T {
	var t []T
	var dropCount = 0
	for i := 0; i < len(list); i++ {
		if dropCount < n {
			if predicate(list[i]) {
				// dropped!
				dropCount++
			} else {
				t = append(t, list[i])
			}
		} else {
			t = append(t, list[i])
		}
	}
	return t
}

func ListDropLastWhile[T any](list []T, n int, predicate func(T) bool) []T {
	var t []T
	var dropCount = 0
	for i := len(list) - 1; i >= 0; i-- {
		if dropCount < n {
			if predicate(list[i]) {
				// dropped!
				dropCount++
			} else {
				t = append(t, list[i])
			}
		} else {
			t = append(t, list[i])
		}
	}
	// reverse
	var r []T
	for i := len(t) - 1; i >= 0; i-- {
		r = append(r, t[i])
	}
	return t
}

func ListPartition[T any](list []T, partition int) [][]T {
	return ListPartitionWithCal(list, func(int) int {
		return partition
	})
}

func ListPartitionWithCal[T any](list []T, f func(int) int) [][]T {
	var array [][]T

	length := len(list)
	if length == 0 {
		return array
	}

	partition := f(length)
	if partition <= 0 {
		array = append(array, list)
		return array
	}
	//list下标
	n := 0
	partitionLen := length / partition
	for i := 0; i < partition; i++ {
		var subList []T
		for j := 0; j < partitionLen; j++ {
			subList = append(subList, list[n])
			n++
		}
		//list有多余
		if i == partition-1 {
			for k := n; k < length; k++ {
				subList = append(subList, list[k])
			}
		}
		array = append(array, subList)
	}
	return array
}

func ListPlus[T any](list []T, n []T) []T {
	var t []T
	for _, e := range list {
		t = append(t, e)
	}
	for _, e := range n {
		if !ListContains(t, e) {
			t = append(t, e)
		}
	}
	return t
}

func ListMinus[T any](list []T, n []T) []T {
	var t []T
	for _, e := range list {
		if !ListContains(n, e) {
			t = append(t, e)
		}
	}
	return t
}

func MapPlus[K comparable, V any](m map[K]V, n map[K]V) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		r[k] = v
	}
	for k, v := range n {
		r[k] = v
	}
	return r
}

func MapMinus[K comparable, V any](m map[K]V, n map[K]V) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		if _, ok := n[k]; !ok {
			r[k] = v
		}
	}
	return r
}

func ListEquals[T any](leftList []T, rightList []T) bool {
	if leftList == nil && rightList == nil {
		return true
	}
	if leftList == nil || rightList == nil {
		return false
	}
	if len(leftList) != len(rightList) {
		return false
	}
	l := ListDistinct(append(leftList, rightList...))
	return len(l) == len(leftList)
}

// MapEquals 比较两个map是否相同
func MapEquals[K comparable, V any](leftMap map[K]V, rightMap map[K]V) bool {
	if leftMap == nil && rightMap == nil {
		return true
	}
	if leftMap == nil || rightMap == nil {
		return false
	}
	if len(leftMap) != len(rightMap) {
		return false
	}

	for key, leftValue := range leftMap {
		rightValue, exist := rightMap[key]
		if !exist || !reflect.DeepEqual(rightValue, leftValue) {
			return false
		}
	}
	return true
}

func ListAssociate[T any, K comparable, V any](list []T, transform func(T) Pair[K, V]) map[K]V {
	r := make(map[K]V)
	return ListAssociateTo(list, &r, transform)
}

func ListAssociateTo[T any, K comparable, V any](list []T, destination *map[K]V, transform func(T) Pair[K, V]) map[K]V {
	for _, e := range list {
		item := transform(e)
		(*destination)[item.First] = item.Second
	}
	return *destination
}

func ListAssociateBy[T any, K comparable](list []T, keySelector func(T) K) map[K]T {
	r := make(map[K]T)
	for _, e := range list {
		r[keySelector(e)] = e
	}
	return r
}

func ListAssociateByAndValue[T any, V any, K comparable](list []T, keySelector func(T) K, valueTransform func(T) V) map[K]V {
	r := make(map[K]V)
	for _, e := range list {
		r[keySelector(e)] = valueTransform(e)
	}
	return r
}

func ListAssociateByTo[T any, K comparable](list []T, destination *map[K]T, keySelector func(T) K) map[K]T {
	for _, e := range list {
		(*destination)[keySelector(e)] = e
	}
	return *destination
}

func ListAssociateByAndValueTo[T, V any, K comparable](list []T, destination *map[K]V, keySelector func(T) K, valueTransform func(T) V) map[K]V {
	for _, e := range list {
		(*destination)[keySelector(e)] = valueTransform(e)
	}
	return *destination
}

func ListAssociateWith[T comparable, V any](list []T, valueSelector func(T) V) map[T]V {
	destination := make(map[T]V)
	for _, e := range list {
		destination[e] = valueSelector(e)
	}
	return destination
}

func ListAssociateWithTo[T comparable, V any](list []T, destination *map[T]V, valueSelector func(T) V) map[T]V {
	for _, e := range list {
		(*destination)[e] = valueSelector(e)
	}
	return *destination
}

func ListFlatMap[T any, R any](list []T, f func(T) []R) []R {
	var r []R
	for _, e := range list {
		rlist := f(e)
		for _, rl := range rlist {
			r = append(r, rl)
		}
	}
	return r
}

func ListFlatMapIndexed[T any, R any](list []T, f func(int, T) []R) []R {
	var r []R
	for i, e := range list {
		rlist := f(i, e)
		for _, rl := range rlist {
			r = append(r, rl)
		}
	}
	return r
}

func ListFlatten[T any](list [][]T) []T {
	var r []T
	for _, e := range list {
		rlist := e
		for _, rl := range rlist {
			r = append(r, rl)
		}
	}
	return r
}

func ListFlatMapTo[T any, R any](list []T, dest *[]R, f func(T) []R) []R {
	var r []R
	for _, e := range list {
		rlist := f(e)
		for _, rl := range rlist {
			*dest = append(*dest, rl)
			r = append(r, rl)
		}
	}
	return r
}

func ListFlatMapIndexedTo[T any, R any](list []T, dest *[]R, f func(int, T) []R) []R {
	var r []R
	for i, e := range list {
		rlist := f(i, e)
		for _, rl := range rlist {
			*dest = append(*dest, rl)
			r = append(r, rl)
		}
	}
	return r
}

func MapFlatMap[K comparable, V any, R any](m map[K]V, f func(K, V) []R) []R {
	var r []R
	for k, v := range m {
		rlist := f(k, v)
		for _, rl := range rlist {
			r = append(r, rl)
		}
	}
	return r
}

func MapFlatMapTo[K comparable, V any, R any](m map[K]V, dest *[]R, f func(K, V) []R) []R {
	var r []R
	for k, v := range m {
		rlist := f(k, v)
		for _, rl := range rlist {
			*dest = append(*dest, rl)
			r = append(r, rl)
		}
	}
	return r
}

func ListGroupBy[T any, K comparable](list []T, keySelector func(T) K) (destination map[K][]T) {
	dest := make(map[K][]T)
	return ListGroupByTo(list, &dest, keySelector)
}

func ListGroupByTransform[T any, K comparable, V any](list []T, keySelector func(T) K, trans func(T) V) map[K][]V {
	dest := make(map[K][]V)
	return ListGroupByTransformTo(list, &dest, keySelector, trans)
}

func ListGroupByTo[T any, K comparable](list []T, dest *map[K][]T, keySelector func(T) K) (destination map[K][]T) {
	r := make(map[K][]T)
	if *dest == nil {
		return r
	}
	for _, e := range list {
		key := keySelector(e)
		sl, _ := r[key]
		sl = append(sl, e)
		r[key] = sl
		sll, _ := (*dest)[key]
		sll = append(sll, e)
		(*dest)[key] = sll
	}
	return r
}

func ListGroupByTransformTo[T any, K comparable, V any](list []T, dest *map[K][]V, keySelector func(T) K, trans func(T) V) map[K][]V {
	r := make(map[K][]V)
	for _, e := range list {
		key := keySelector(e)
		value := trans(e)
		sl, _ := r[key]
		sl = append(sl, value)
		r[key] = sl
		sll, _ := (*dest)[key]
		sll = append(sll, value)
		(*dest)[key] = sll
	}
	return r
}

func ListMap[T any, R any](list []T, f func(T) R) []R {
	var n []R
	for _, e := range list {
		n = append(n, f(e))
	}
	return n
}

func ListMapIndexed[T any, R any](list []T, f func(int, T) R) []R {
	var n []R
	for i, e := range list {
		n = append(n, f(i, e))
	}
	return n
}

func ListMapTo[T any, R any](list []T, dest *[]R, f func(T) R) []R {
	var n []R
	for _, e := range list {
		item := f(e)
		*dest = append(*dest, item)
		n = append(n, item)
	}
	return n
}

func ListMapIndexedTo[T any, R any](list []T, dest *[]R, f func(int, T) R) []R {
	var n []R
	for i, e := range list {
		item := f(i, e)
		*dest = append(*dest, item)
		n = append(n, item)
	}
	return n
}

func MapMap[K comparable, V any, R any](m map[K]V, f func(K, V) R) []R {
	var n []R
	for k, v := range m {
		n = append(n, f(k, v))
	}
	return n
}

func MapMapTo[K comparable, V any, R any](m map[K]V, dest *[]R, f func(K, V) R) []R {
	var n []R
	for k, v := range m {
		item := f(k, v)
		*dest = append(*dest, item)
		n = append(n, item)
	}
	return n
}

func ListReduce[S any, T any](list []T, init func(T) S, f func(S, T) S) S {
	accumulator := init(list[0])
	for _, e := range list[1:] {
		accumulator = f(accumulator, e)
	}
	return accumulator
}

func ListReduceIndexed[S any, T any](list []T, init func(int, T) S, f func(int, S, T) S) S {
	accumulator := init(0, list[0])
	for i, e := range list[1:] {
		accumulator = f(i, accumulator, e)
	}
	return accumulator
}

func ListToMap[K comparable, V any](list []Pair[K, V]) map[K]V {
	m := make(map[K]V)
	for _, item := range list {
		m[item.First] = item.Second
	}
	return m
}

func MapToList[K comparable, V any](m map[K]V) []Pair[K, V] {
	var n []Pair[K, V]
	for k, v := range m {
		n = append(n, NewPair(k, v))
	}
	return n
}

func SliceDistinct[T any](list []T) []T {
	result := NewList[T]()
	for _, k := range list {
		if !result.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

func SliceDistinctTo[T any, V comparable](list []T, valueTransform func(T) V) []T {
	m := SliceTo(list, valueTransform)
	var result []T
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func SliceSubList[T any](list []T, fromIndex int, toIndex int) []T {
	m := map[int]int{fromIndex: toIndex, toIndex: fromIndex}
	start := fromIndex
	if fromIndex < start {
		start = toIndex
	}
	end := m[start]

	if start < 0 {
		start = 0
	}
	if start == end {
		return []T{}
	}
	if end > len(list) {
		return list[start:]
	}
	return list[start:end]
}

func SliceSub[T any](list []T, r IntRange) []T {
	return SliceSubList(list, int(r.Start), int(r.End))
}

func SliceContains[T any, K comparable](list []T, predicate func(T) K, key K) bool {
	m := SliceTo(list, predicate)
	_, ok := m[key]
	return ok
}

func SliceTo[T any, K comparable](list []T, valueTransform func(T) K) map[K]T {
	m := make(map[K]T)
	for _, e := range list {
		m[valueTransform(e)] = e
	}
	return m
}

func InSlice[T comparable](list []T, val T) bool {
	return ListIndexOf(list, val) >= 0
}
