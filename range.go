package xgo

type IntRange struct {
	Start XInt
	End   XInt
}

type IntRange64 struct {
	Start XInt64
	End   XInt64
}

func MakeIntRange(AStart XInt, AEnd XInt) IntRange {
	return IntRange{
		Start: AStart,
		End:   AEnd,
	}
}

func MakeIntRange64(AStart XInt64, AEnd XInt64) IntRange64 {
	return IntRange64{
		Start: AStart,
		End:   AEnd,
	}
}
