package xgo

import "errors"

type XMap[K comparable, V any] map[K]V

func NewMap[K comparable, V any]() XMap[K, V] {
	return make(XMap[K, V])
}

func NewMapWithMap[K comparable, V any](ma map[K]V) XMap[K, V] {
	return ma
}

func NewMapWithPairs[K comparable, V any](pairs ...Pair[K, V]) XMap[K, V] {
	m := make(map[K]V)
	for _, item := range pairs {
		m[item.First] = item.Second
	}
	return m
}

func MapOf[K comparable, V any](pairs ...Pair[K, V]) XMap[K, V] {
	return NewMapWithPairs(pairs...)
}

func (x XMap[K, V]) IsEmpty() bool {
	return len(x) == 0
}

func (x XMap[K, V]) IsNotEmpty() bool {
	return len(x) != 0
}

func (x XMap[K, V]) Size() XInt {
	return XInt(len(x))
}

func (x *XMap[K, V]) Put(k K, v V) {
	(*x)[k] = v
}

func (x *XMap[K, V]) PutPair(item Pair[K, V]) {
	(*x)[item.First] = item.Second
}

func (x *XMap[K, V]) PutAllPairs(item ...Pair[K, V]) {
	for _, e := range item {
		x.PutPair(e)
	}
}

func (x XMap[K, V]) Get(k K) (V, error) {
	if v, ok := x[k]; ok {
		return v, nil
	} else {
		return v, errors.New("not found")
	}
}

func (x XMap[K, V]) GetOrDef(k K, def V) V {
	if v, ok := x[k]; ok {
		return v
	} else {
		return def
	}
}

func (x *XMap[K, V]) Delete(k K) {
	delete(*x, k)
}

func (x *XMap[K, V]) Clear() {
	*x = make(XMap[K, V])
}

func (x XMap[K, V]) ForEach(action func(K, V)) {
	MapForEach(x, action)
}

func (x XMap[K, V]) Filter(predicate func(K, V) bool) XMap[K, V] {
	return MapFilter(x, predicate)
}

func (x XMap[K, V]) FilterNot(predicate func(K, V) bool) XMap[K, V] {
	return MapFilterNot(x, predicate)
}

func (x XMap[K, V]) FilterKeys(predicate func(K) bool) XMap[K, V] {
	return MapFilterKeys(x, predicate)
}

func (x XMap[K, V]) FilterValues(predicate func(V) bool) XMap[K, V] {
	return MapFilterValues(x, predicate)
}

func (x XMap[K, V]) FilterTo(dest *map[K]V, predicate func(K, V) bool) XMap[K, V] {
	return MapFilterTo(x, dest, predicate)
}

func (x XMap[K, V]) FilterNotTo(dest *map[K]V, predicate func(K, V) bool) XMap[K, V] {
	return MapFilterNotTo(x, dest, predicate)
}

func (x XMap[K, V]) Contains(k K, v V) bool {
	return MapContains(x, k, v)
}

func (x XMap[K, V]) ContainsKey(k K) bool {
	return MapContainsKey(x, k)
}

func (x XMap[K, V]) ContainsValue(v V) bool {
	return MapContainsValue(x, v)
}

func (x XMap[K, V]) JoinToString(f func(K, V) string) XString {
	return XString(MapJoinToString(x, f))
}

func (x XMap[K, V]) JoinToStringFull(sep XString, prefix XString, postfix XString, joint func(K, V) string) XString {
	return XString(MapJoinToStringFull(x, string(sep), string(prefix), string(postfix), joint))
}

func (x XMap[K, V]) All(predicate func(K, V) bool) bool {
	return MapAll(x, predicate)
}

func (x XMap[K, V]) Any(predicate func(K, V) bool) bool {
	return MapAny(x, predicate)
}

func (x XMap[K, V]) None(predicate func(K, V) bool) bool {
	return MapNone(x, predicate)
}

func (x XMap[K, V]) Count(predicate func(K, V) bool) XInt {
	return XInt(MapCount(x, predicate))
}

func (x XMap[K, V]) AllKey(predicate func(K) bool) bool {
	return MapAllKey(x, predicate)
}

func (x XMap[K, V]) AnyKey(predicate func(K) bool) bool {
	return MapAnyKey(x, predicate)
}

func (x XMap[K, V]) NoneKey(predicate func(K) bool) bool {
	return MapNoneKey(x, predicate)
}

func (x XMap[K, V]) CountKey(predicate func(K) bool) XInt {
	return XInt(MapCountKey(x, predicate))
}

func (x XMap[K, V]) AllValue(predicate func(V) bool) bool {
	return MapAllValue(x, predicate)
}

func (x XMap[K, V]) AnyValue(predicate func(V) bool) bool {
	return MapAnyValue(x, predicate)
}

func (x XMap[K, V]) NoneValue(predicate func(V) bool) bool {
	return MapNoneValue(x, predicate)
}

func (x XMap[K, V]) CountValue(predicate func(V) bool) XInt {
	return XInt(MapCountValue(x, predicate))
}

func (x XMap[K, V]) ToList() []Pair[K, V] {
	var n []Pair[K, V]
	for k, v := range x {
		n = append(n, NewPair(k, v))
	}
	return n
}

func (x XMap[K, V]) Plus(n map[K]V) XMap[K, V] {
	return MapPlus(x, n)
}

func (x XMap[K, V]) Minus(n map[K]V) XMap[K, V] {
	return MapMinus(x, n)
}

func (x XMap[K, V]) Equals(n map[K]V) bool {
	return MapEquals(x, n)
}

func (x XMap[K, V]) Keys() XList[K] {
	i := 0
	keys := make([]K, len(x))
	for k := range x {
		keys[i] = k
		i++
	}
	return keys
}

func (x XMap[K, V]) Values() XList[V] {
	i := 0
	vals := make([]V, len(x))
	for _, v := range x {
		vals[i] = v
		i++
	}
	return vals
}

func (x XMap[K, V]) Entries() XSet[Pair[K, V]] {
	ret := XSet[Pair[K, V]]{}
	for k, v := range x {
		_ = ret.Add(Pair[K, V]{
			First:  k,
			Second: v,
		})
	}
	return ret
}
