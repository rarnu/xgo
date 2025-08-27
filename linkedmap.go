package xgo

type LinkedMap[K comparable, V comparable] struct {
	Data    map[K]V
	KeyList []K
}

func NewLinkedMap[K comparable, V comparable]() LinkedMap[K, V] {
	return LinkedMap[K, V]{
		Data:    make(map[K]V),
		KeyList: []K{},
	}
}

func (m LinkedMap[K, V]) Size() int {
	return len(m.KeyList)
}

func (m *LinkedMap[K, V]) Put(k K, v V) {
	if !ListContains(m.KeyList, k) {
		m.KeyList = append(m.KeyList, k)
	}
	m.Data[k] = v
}

func (m *LinkedMap[K, V]) PutPair(item Pair[K, V]) {
	if !ListContains(m.KeyList, item.First) {
		m.KeyList = append(m.KeyList, item.First)
	}
	m.Data[item.First] = item.Second
}

func (m *LinkedMap[K, V]) PutPairs(item ...Pair[K, V]) {
	for _, pair := range item {
		m.PutPair(pair)
	}
}

func (m LinkedMap[K, V]) Get(k K) V {
	return m.Data[k]
}

func (m LinkedMap[K, V]) GetOrDef(k K, def V) V {
	if v, ok := m.Data[k]; ok {
		return v
	} else {
		return def
	}
}

func (m *LinkedMap[K, V]) Delete(k K) {
	idx := ListIndexOf(m.KeyList, k)
	if idx != -1 {
		m.KeyList = append(m.KeyList[:idx], m.KeyList[idx+1:]...)
		delete(m.Data, k)
	}
}

func (m *LinkedMap[K, V]) Clear() {
	m.KeyList = []K{}
	m.Data = make(map[K]V)
}

func (m LinkedMap[K, V]) Keys() []K {
	return m.KeyList
}

func (m LinkedMap[K, V]) GetKey(index XInt) K {
	return m.KeyList[index]
}

func (m LinkedMap[K, V]) GetValue(index XInt) V {
	key := m.KeyList[index]
	return m.Data[key]
}

func (m LinkedMap[K, V]) ForEach(action func(K, V)) {
	for _, key := range m.KeyList {
		action(key, m.Data[key])
	}
}

func (m LinkedMap[K, V]) ForEachIndexed(action func(int, K, V)) {
	for idx, key := range m.KeyList {
		action(idx, key, m.Data[key])
	}
}

func (m LinkedMap[K, V]) Filter(f func(K, V) bool) LinkedMap[K, V] {
	result := NewLinkedMap[K, V]()
	for _, key := range m.KeyList {
		if f(key, m.Data[key]) {
			result.Put(key, m.Data[key])
		}
	}
	return result
}

func (m LinkedMap[K, V]) FilterIndexed(f func(int, K, V) bool) LinkedMap[K, V] {
	result := NewLinkedMap[K, V]()
	for idx, key := range m.KeyList {
		if f(idx, key, m.Data[key]) {
			result.Put(key, m.Data[key])
		}
	}
	return result
}

func (m LinkedMap[K, V]) FilterNot(f func(K, V) bool) LinkedMap[K, V] {
	result := NewLinkedMap[K, V]()
	for _, key := range m.KeyList {
		if !f(key, m.Data[key]) {
			result.Put(key, m.Data[key])
		}
	}
	return result
}

func (m LinkedMap[K, V]) FilterNotIndexed(f func(int, K, V) bool) LinkedMap[K, V] {
	result := NewLinkedMap[K, V]()
	for idx, key := range m.KeyList {
		if !f(idx, key, m.Data[key]) {
			result.Put(key, m.Data[key])
		}
	}
	return result
}

func (m LinkedMap[K, V]) FilterKeys(f func(K) bool) LinkedMap[K, V] {
	result := NewLinkedMap[K, V]()
	for _, key := range m.KeyList {
		if f(key) {
			result.Put(key, m.Data[key])
		}
	}
	return result
}

func (m LinkedMap[K, V]) FilterKeysIndexed(f func(int, K) bool) LinkedMap[K, V] {
	result := NewLinkedMap[K, V]()
	for idx, key := range m.KeyList {
		if f(idx, key) {
			result.Put(key, m.Data[key])
		}
	}
	return result
}

func (m LinkedMap[K, V]) FilterValues(f func(V) bool) LinkedMap[K, V] {
	result := NewLinkedMap[K, V]()
	for _, key := range m.KeyList {
		if f(m.Data[key]) {
			result.Put(key, m.Data[key])
		}
	}
	return result
}

func (m LinkedMap[K, V]) FilterValuesIndexed(f func(int, V) bool) LinkedMap[K, V] {
	result := NewLinkedMap[K, V]()
	for idx, key := range m.KeyList {
		if f(idx, m.Data[key]) {
			result.Put(key, m.Data[key])
		}
	}
	return result
}

func (m LinkedMap[K, V]) FilterTo(dest *LinkedMap[K, V], f func(K, V) bool) LinkedMap[K, V] {
	result := NewLinkedMap[K, V]()
	for _, key := range m.KeyList {
		if f(key, m.Data[key]) {
			dest.Put(key, m.Data[key])
			result.Put(key, m.Data[key])
		}
	}
	return result
}

func (m LinkedMap[K, V]) FilterIndexedTo(dest *LinkedMap[K, V], f func(int, K, V) bool) LinkedMap[K, V] {
	result := NewLinkedMap[K, V]()
	for idx, key := range m.KeyList {
		if f(idx, key, m.Data[key]) {
			dest.Put(key, m.Data[key])
			result.Put(key, m.Data[key])
		}
	}
	return result
}

func (m LinkedMap[K, V]) FilterNotTo(dest *LinkedMap[K, V], f func(K, V) bool) LinkedMap[K, V] {
	result := NewLinkedMap[K, V]()
	for _, key := range m.KeyList {
		if !f(key, m.Data[key]) {
			dest.Put(key, m.Data[key])
			result.Put(key, m.Data[key])
		}
	}
	return result
}

func (m LinkedMap[K, V]) FilterNotIndexedTo(dest *LinkedMap[K, V], f func(int, K, V) bool) LinkedMap[K, V] {
	result := NewLinkedMap[K, V]()
	for idx, key := range m.KeyList {
		if !f(idx, key, m.Data[key]) {
			dest.Put(key, m.Data[key])
			result.Put(key, m.Data[key])
		}
	}
	return result
}

func (m LinkedMap[K, V]) Contains(k K, v V) bool {
	return MapContains(m.Data, k, v)
}

func (m LinkedMap[K, V]) ContainsKey(k K) bool {
	return MapContainsKey(m.Data, k)
}

func (m LinkedMap[K, V]) ContainsValue(v V) bool {
	return MapContainsValue(m.Data, v)
}

func (m LinkedMap[K, V]) JoinToString(f func(K, V) string) XString {
	return m.JoinToStringFull(",", "", "", f)
}

func (m LinkedMap[K, V]) JoinToStringFull(sep XString, prefix XString, postfix XString, f func(K, V) string) XString {
	buffer := prefix
	var count = 0
	for _, key := range m.KeyList {
		count++
		if count > 1 {
			buffer += sep
		}
		buffer += XString(f(key, m.Data[key]))
	}
	buffer += postfix
	return buffer
}

func (m LinkedMap[K, V]) All(f func(K, V) bool) bool {
	for _, key := range m.KeyList {
		if !f(key, m.Data[key]) {
			return false
		}
	}
	return true
}

func (m LinkedMap[K, V]) Any(f func(K, V) bool) bool {
	for _, key := range m.KeyList {
		if f(key, m.Data[key]) {
			return true
		}
	}
	return false
}

func (m LinkedMap[K, V]) None(f func(K, V) bool) bool {
	for _, key := range m.KeyList {
		if f(key, m.Data[key]) {
			return false
		}
	}
	return true
}

func (m LinkedMap[K, V]) Count(f func(K, V) bool) XInt {
	num := 0
	for _, key := range m.KeyList {
		if f(key, m.Data[key]) {
			num++
		}
	}
	return XInt(num)
}

func (m LinkedMap[K, V]) AllKey(f func(K) bool) bool {
	for _, key := range m.KeyList {
		if !f(key) {
			return false
		}
	}
	return true
}

func (m LinkedMap[K, V]) AnyKey(f func(K) bool) bool {
	for _, key := range m.KeyList {
		if f(key) {
			return true
		}
	}
	return false
}

func (m LinkedMap[K, V]) NoneKey(f func(K) bool) bool {
	for _, key := range m.KeyList {
		if f(key) {
			return false
		}
	}
	return true
}

func (m LinkedMap[K, V]) CountKey(f func(K) bool) XInt {
	num := 0
	for _, key := range m.KeyList {
		if f(key) {
			num++
		}
	}
	return XInt(num)
}

func (m LinkedMap[K, V]) AllValue(f func(V) bool) bool {
	for _, key := range m.KeyList {
		if !f(m.Data[key]) {
			return false
		}
	}
	return true
}

func (m LinkedMap[K, V]) AnyValue(f func(V) bool) bool {
	for _, key := range m.KeyList {
		if f(m.Data[key]) {
			return true
		}
	}
	return false
}

func (m LinkedMap[K, V]) NoneValue(f func(V) bool) bool {
	for _, key := range m.KeyList {
		if f(m.Data[key]) {
			return false
		}
	}
	return true
}

func (m LinkedMap[K, V]) CountValue(f func(V) bool) XInt {
	num := 0
	for _, key := range m.KeyList {
		if f(m.Data[key]) {
			num++
		}
	}
	return XInt(num)
}

func (m LinkedMap[K, V]) ToList() XList[Pair[K, V]] {
	var n XList[Pair[K, V]]
	for _, key := range m.KeyList {
		n = append(n, NewPair(key, m.Data[key]))
	}
	return n
}

func (m LinkedMap[K, V]) Plus(n LinkedMap[K, V]) LinkedMap[K, V] {
	r := NewLinkedMap[K, V]()
	for _, key := range m.KeyList {
		r.Put(key, m.Data[key])
	}
	n.ForEach(func(k K, v V) {
		r.Put(k, v)
	})
	return r
}

func (m LinkedMap[K, V]) Minus(n LinkedMap[K, V]) LinkedMap[K, V] {
	r := NewLinkedMap[K, V]()
	for _, key := range m.KeyList {
		if _, ok := n.Data[key]; !ok {
			r.Put(key, m.Data[key])
		}
	}
	return r
}
