package xgo

type Pair[A any, B any] struct {
	First  A `json:"first"`
	Second B `json:"second"`
}

type Triple[A any, B any, C any] struct {
	First  A `json:"first"`
	Second B `json:"second"`
	Third  C `json:"third"`
}

type Quadruple[A any, B any, C any, D any] struct {
	First  A `json:"first"`
	Second B `json:"second"`
	Third  C `json:"third"`
	Fourth D `json:"fourth"`
}

type Pentuple[A any, B any, C any, D any, E any] struct {
	First  A `json:"first"`
	Second B `json:"second"`
	Third  C `json:"third"`
	Fourth D `json:"fourth"`
	Fifth  E `json:"fifth"`
}

type Hextuple[A any, B any, C any, D any, E any, F any] struct {
	First  A `json:"first"`
	Second B `json:"second"`
	Third  C `json:"third"`
	Fourth D `json:"fourth"`
	Fifth  E `json:"fifth"`
	Sixth  F `json:"sixth"`
}

type Septuple[A any, B any, C any, D any, E any, F any, G any] struct {
	First   A `json:"first"`
	Second  B `json:"second"`
	Third   C `json:"third"`
	Fourth  D `json:"fourth"`
	Fifth   E `json:"fifth"`
	Sixth   F `json:"sixth"`
	Seventh G `json:"seventh"`
}

type Octuple[A any, B any, C any, D any, E any, F any, G any, H any] struct {
	First   A `json:"first"`
	Second  B `json:"second"`
	Third   C `json:"third"`
	Fourth  D `json:"fourth"`
	Fifth   E `json:"fifth"`
	Sixth   F `json:"sixth"`
	Seventh G `json:"seventh"`
	Eighth  H `json:"eighth"`
}

type Nonuple[A any, B any, C any, D any, E any, F any, G any, H any, I any] struct {
	First   A `json:"first"`
	Second  B `json:"second"`
	Third   C `json:"third"`
	Fourth  D `json:"fourth"`
	Fifth   E `json:"fifth"`
	Sixth   F `json:"sixth"`
	Seventh G `json:"seventh"`
	Eighth  H `json:"eighth"`
	Ninth   I `json:"ninth"`
}

type Decuple[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any] struct {
	First   A `json:"first"`
	Second  B `json:"second"`
	Third   C `json:"third"`
	Fourth  D `json:"fourth"`
	Fifth   E `json:"fifth"`
	Sixth   F `json:"sixth"`
	Seventh G `json:"seventh"`
	Eighth  H `json:"eighth"`
	Ninth   I `json:"ninth"`
	Tenth   J `json:"tenth"`
}

type Undecuple[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any] struct {
	First    A `json:"first"`
	Second   B `json:"second"`
	Third    C `json:"third"`
	Fourth   D `json:"fourth"`
	Fifth    E `json:"fifth"`
	Sixth    F `json:"sixth"`
	Seventh  G `json:"seventh"`
	Eighth   H `json:"eighth"`
	Ninth    I `json:"ninth"`
	Tenth    J `json:"tenth"`
	Eleventh K `json:"eleventh"`
}

type Duodecuple[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any] struct {
	First    A `json:"first"`
	Second   B `json:"second"`
	Third    C `json:"third"`
	Fourth   D `json:"fourth"`
	Fifth    E `json:"fifth"`
	Sixth    F `json:"sixth"`
	Seventh  G `json:"seventh"`
	Eighth   H `json:"eighth"`
	Ninth    I `json:"ninth"`
	Tenth    J `json:"tenth"`
	Eleventh K `json:"eleventh"`
	Twelfth  L `json:"twelfth"`
}

type T2[A any, B any] = Pair[A, B]
type T3[A any, B any, C any] = Triple[A, B, C]
type T4[A any, B any, C any, D any] = Quadruple[A, B, C, D]
type T5[A any, B any, C any, D any, E any] = Pentuple[A, B, C, D, E]
type T6[A any, B any, C any, D any, E any, F any] = Hextuple[A, B, C, D, E, F]
type T7[A any, B any, C any, D any, E any, F any, G any] = Septuple[A, B, C, D, E, F, G]
type T8[A any, B any, C any, D any, E any, F any, G any, H any] = Octuple[A, B, C, D, E, F, G, H]
type T9[A any, B any, C any, D any, E any, F any, G any, H any, I any] = Nonuple[A, B, C, D, E, F, G, H, I]
type T10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any] = Decuple[A, B, C, D, E, F, G, H, I, J]
type T11[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any] = Undecuple[A, B, C, D, E, F, G, H, I, J, K]
type T12[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any] = Duodecuple[A, B, C, D, E, F, G, H, I, J, K, L]

func NewPair[A any, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{
		First:  a,
		Second: b,
	}
}

func NewT2[A any, B any](a A, b B) T2[A, B] {
	return T2[A, B]{
		First:  a,
		Second: b,
	}
}

func NewTriple[A any, B any, C any](a A, b B, c C) Triple[A, B, C] {
	return Triple[A, B, C]{
		First:  a,
		Second: b,
		Third:  c,
	}
}

func NewT3[A any, B any, C any](a A, b B, c C) T3[A, B, C] {
	return T3[A, B, C]{
		First:  a,
		Second: b,
		Third:  c,
	}
}

func NewQuadruple[A any, B any, C any, D any](a A, b B, c C, d D) Quadruple[A, B, C, D] {
	return Quadruple[A, B, C, D]{
		First:  a,
		Second: b,
		Third:  c,
		Fourth: d,
	}
}

func NewT4[A any, B any, C any, D any](a A, b B, c C, d D) T4[A, B, C, D] {
	return T4[A, B, C, D]{
		First:  a,
		Second: b,
		Third:  c,
		Fourth: d,
	}
}

func NewPentuple[A any, B any, C any, D any, E any](a A, b B, c C, d D, e E) Pentuple[A, B, C, D, E] {
	return Pentuple[A, B, C, D, E]{
		First:  a,
		Second: b,
		Third:  c,
		Fourth: d,
		Fifth:  e,
	}
}

func NewT5[A any, B any, C any, D any, E any](a A, b B, c C, d D, e E) T5[A, B, C, D, E] {
	return T5[A, B, C, D, E]{
		First:  a,
		Second: b,
		Third:  c,
		Fourth: d,
		Fifth:  e,
	}
}

func NewHextuple[A any, B any, C any, D any, E any, F any](a A, b B, c C, d D, e E, f F) Hextuple[A, B, C, D, E, F] {
	return Hextuple[A, B, C, D, E, F]{
		First:  a,
		Second: b,
		Third:  c,
		Fourth: d,
		Fifth:  e,
		Sixth:  f,
	}
}

func NewT6[A any, B any, C any, D any, E any, F any](a A, b B, c C, d D, e E, f F) T6[A, B, C, D, E, F] {
	return T6[A, B, C, D, E, F]{
		First:  a,
		Second: b,
		Third:  c,
		Fourth: d,
		Fifth:  e,
		Sixth:  f,
	}
}

func NewSeptuple[A any, B any, C any, D any, E any, F any, G any](a A, b B, c C, d D, e E, f F, g G) Septuple[A, B, C, D, E, F, G] {
	return Septuple[A, B, C, D, E, F, G]{
		First:   a,
		Second:  b,
		Third:   c,
		Fourth:  d,
		Fifth:   e,
		Sixth:   f,
		Seventh: g,
	}
}

func NewT7[A any, B any, C any, D any, E any, F any, G any](a A, b B, c C, d D, e E, f F, g G) T7[A, B, C, D, E, F, G] {
	return T7[A, B, C, D, E, F, G]{
		First:   a,
		Second:  b,
		Third:   c,
		Fourth:  d,
		Fifth:   e,
		Sixth:   f,
		Seventh: g,
	}
}

func NewOctuple[A any, B any, C any, D any, E any, F any, G any, H any](a A, b B, c C, d D, e E, f F, g G, h H) Octuple[A, B, C, D, E, F, G, H] {
	return Octuple[A, B, C, D, E, F, G, H]{
		First:   a,
		Second:  b,
		Third:   c,
		Fourth:  d,
		Fifth:   e,
		Sixth:   f,
		Seventh: g,
		Eighth:  h,
	}
}

func NewT8[A any, B any, C any, D any, E any, F any, G any, H any](a A, b B, c C, d D, e E, f F, g G, h H) T8[A, B, C, D, E, F, G, H] {
	return T8[A, B, C, D, E, F, G, H]{
		First:   a,
		Second:  b,
		Third:   c,
		Fourth:  d,
		Fifth:   e,
		Sixth:   f,
		Seventh: g,
		Eighth:  h,
	}
}

func NewNonuple[A any, B any, C any, D any, E any, F any, G any, H any, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) Nonuple[A, B, C, D, E, F, G, H, I] {
	return Nonuple[A, B, C, D, E, F, G, H, I]{
		First:   a,
		Second:  b,
		Third:   c,
		Fourth:  d,
		Fifth:   e,
		Sixth:   f,
		Seventh: g,
		Eighth:  h,
		Ninth:   i,
	}
}

func NewT9[A any, B any, C any, D any, E any, F any, G any, H any, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) T9[A, B, C, D, E, F, G, H, I] {
	return T9[A, B, C, D, E, F, G, H, I]{
		First:   a,
		Second:  b,
		Third:   c,
		Fourth:  d,
		Fifth:   e,
		Sixth:   f,
		Seventh: g,
		Eighth:  h,
		Ninth:   i,
	}
}

func NewDecuple[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](a A, b B, c C, d D, e E, f F, g G, h H, i I, j J) Decuple[A, B, C, D, E, F, G, H, I, J] {
	return Decuple[A, B, C, D, E, F, G, H, I, J]{
		First:   a,
		Second:  b,
		Third:   c,
		Fourth:  d,
		Fifth:   e,
		Sixth:   f,
		Seventh: g,
		Eighth:  h,
		Ninth:   i,
		Tenth:   j,
	}
}

func NewT10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](a A, b B, c C, d D, e E, f F, g G, h H, i I, j J) T10[A, B, C, D, E, F, G, H, I, J] {
	return T10[A, B, C, D, E, F, G, H, I, J]{
		First:   a,
		Second:  b,
		Third:   c,
		Fourth:  d,
		Fifth:   e,
		Sixth:   f,
		Seventh: g,
		Eighth:  h,
		Ninth:   i,
		Tenth:   j,
	}
}

func NewUndecuple[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any](a A, b B, c C, d D, e E, f F, g G, h H, i I, j J, k K) Undecuple[A, B, C, D, E, F, G, H, I, J, K] {
	return Undecuple[A, B, C, D, E, F, G, H, I, J, K]{
		First:    a,
		Second:   b,
		Third:    c,
		Fourth:   d,
		Fifth:    e,
		Sixth:    f,
		Seventh:  g,
		Eighth:   h,
		Ninth:    i,
		Tenth:    j,
		Eleventh: k,
	}
}

func NewT11[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any](a A, b B, c C, d D, e E, f F, g G, h H, i I, j J, k K) T11[A, B, C, D, E, F, G, H, I, J, K] {
	return T11[A, B, C, D, E, F, G, H, I, J, K]{
		First:    a,
		Second:   b,
		Third:    c,
		Fourth:   d,
		Fifth:    e,
		Sixth:    f,
		Seventh:  g,
		Eighth:   h,
		Ninth:    i,
		Tenth:    j,
		Eleventh: k,
	}
}

func NewDuodecuple[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any](a A, b B, c C, d D, e E, f F, g G, h H, i I, j J, k K, l L) Duodecuple[A, B, C, D, E, F, G, H, I, J, K, L] {
	return Duodecuple[A, B, C, D, E, F, G, H, I, J, K, L]{
		First:    a,
		Second:   b,
		Third:    c,
		Fourth:   d,
		Fifth:    e,
		Sixth:    f,
		Seventh:  g,
		Eighth:   h,
		Ninth:    i,
		Tenth:    j,
		Eleventh: k,
		Twelfth:  l,
	}
}

func NewT12[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any](a A, b B, c C, d D, e E, f F, g G, h H, i I, j J, k K, l L) T12[A, B, C, D, E, F, G, H, I, J, K, L] {
	return T12[A, B, C, D, E, F, G, H, I, J, K, L]{
		First:    a,
		Second:   b,
		Third:    c,
		Fourth:   d,
		Fifth:    e,
		Sixth:    f,
		Seventh:  g,
		Eighth:   h,
		Ninth:    i,
		Tenth:    j,
		Eleventh: k,
		Twelfth:  l,
	}
}
