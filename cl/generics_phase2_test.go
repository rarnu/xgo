/*
 * Copyright (c) 2025 The XGo Authors (xgo.dev). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cl_test

import (
	"testing"
)

// Test pure XGo syntax for generic type definitions
func TestGenericTypeDefXGo(t *testing.T) {
	// Note: 'any' is printed as 'interface{}' by the Go AST printer
	gopClTest(t, `package main

type MyType[T any] struct {
	Value T
}

type MyType2[K comparable, V any] struct {
	Key K
	Value V
}
`, `package main

type MyType[T interface{}] struct {
	Value T
}
type MyType2[K comparable, V interface{}] struct {
	Key   K
	Value V
}
`)
}

// Test generic constraint in XGo
func TestGenericConstraintXGo(t *testing.T) {
	gopClTest(t, `package main

func Equal[T comparable](x, y T) bool {
	return x == y
}
`, `package main

func Equal[T comparable](x T, y T) bool {
	return x == y
}
`)
}

// Test using generic type in function body
func TestGenericTypeInBodyXGo(t *testing.T) {
	gopClTest(t, `package main

func CreatePair[T any](x T) [2]T {
	var result [2]T
	result[0] = x
	result[1] = x
	return result
}
`, `package main

func CreatePair[T interface{}](x T) [2]T {
	var result [2]T
	result[0] = x
	result[1] = x
	return result
}
`)
}

func TestGenericMultipleTypeParams(t *testing.T) {
	gopClTest(t, `package main

func Combine[K comparable, V any](k K, v V) map[K]V {
	var result map[K]V
	return result
}
`, `package main

func Combine[K comparable, V interface{}](k K, v V) map[K]V {
	var result map[K]V
	return result
}
`)
}

// Test composite literal instantiation with generic type
func TestGenericCompositeLiteralXGo(t *testing.T) {
	gopClTest(t, `package main

type Pair[K comparable, V any] struct {
	Key K
	Value V
}

func NewPair[K comparable, V any](k K, v V) Pair[K, V] {
	return Pair[K, V]{Key: k, Value: v}
}
`, `package main

type Pair[K comparable, V interface{}] struct {
	Key   K
	Value V
}

func NewPair[K comparable, V interface{}](k K, v V) Pair[K, V] {
	return Pair[K, V]{Key: k, Value: v}
}
`)
}

func TestGenericStructMethodsXGO(t *testing.T) {
	gopClTest(t, `package main

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

func (p *Pair[K, V]) GetKey() K {
	return p.Key
}

func (p *Pair[K, V]) GetValue() V {
	return p.Value
}
`, `package main

type Pair[K comparable, V interface{}] struct {
	Key   K
	Value V
}

func (p *Pair[K, V]) GetKey() K {
	return p.Key
}
func (p *Pair[K, V]) GetValue() V {
	return p.Value
}
`)
}

func TestGenericTypeConstraintDefineXGo(t *testing.T) {
	gopClTest(t, `package main

type Number interface {
	~int | int64 | ~uint | float64
}

func Add[T Number](x, y T) T {
	return x + y
}
`, `package main

type Number interface {
	~int | int64 | ~uint | float64
}

func Add[T Number](x T, y T) T {
	return x + y
}
`)
}

func TestGenericArrayTypeConstraintDefineXGo(t *testing.T) {
	gopClTest(t, `package main

type ArrayType[E any] interface {
	~[]E
}

func At[T ArrayType[E], E any](x T, i int) E {
	return x[i]
}
`, `package main

type ArrayType[E interface{}] interface {
	~[]E
}

func At[T ArrayType[E], E any](x T, i int) E {
	return x[i]
}
`)
}

func TestGenericMethodConstraintDefineXGo(t *testing.T) {
	gopClTest(t, `package main

func Add[T interface{ ~int | int64 | ~uint | float64 }](x, y T) T {
	return x + y
}
`, `package main

func Add[T interface {
	~int | int64 | ~uint | float64
}](x T, y T) T {
	return x + y
}
`)
}

func TestGenericMethodConstraintDefineXGO(t *testing.T) {
	gopClTest(t, `package main

func At[T interface{ ~[]E }, E any](x T, i int) E {
	return x[i]
}

var	AtInt = At[[]int]

func Sample() {
	i := AtInt([1,2,3,4,5], 2)
	println(i)
}
`, `package main

import "fmt"

func At[T interface {
	~[]E
}, E interface{}](x T, i int) E {
	return x[i]
}

var AtInt = At[[]int]

func Sample() {
	i := AtInt([]int{1, 2, 3, 4, 5}, 2)
	fmt.Println(i)
}
`)
}
