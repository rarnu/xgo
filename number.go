package xgo

import (
	"fmt"
	"strconv"
	"unicode"
)

const (
	MinLowSurrogate  = 0xDC00
	MaxLowSurrogate  = 0xDFFF
	MinHighSurrogate = 0xD800
	MaxHighSurrogate = 0xDBFF
)

type XInt int
type XInt64 int64
type XFloat float32
type XFloat64 float64
type XChar rune

func (i XInt) RangeTo(to XInt) XList[XInt] {
	var ret XList[XInt]
	for ii := i; ii <= to; ii++ {
		ret = append(ret, ii)
	}
	return ret
}

func (i XInt) RangeStepTo(to XInt, step XInt) XList[XInt] {
	var ret XList[XInt]
	for ii := i; ii <= to; ii += step {
		ret = append(ret, ii)
	}
	return ret
}

func (i XInt) DownTo(to XInt) XList[XInt] {
	var ret XList[XInt]
	for ii := i; ii >= to; ii-- {
		ret = append(ret, ii)
	}
	return ret
}

func (i XInt) DownStepTo(to XInt, step XInt) XList[XInt] {
	var ret XList[XInt]
	for ii := i; ii >= to; ii -= step {
		ret = append(ret, ii)
	}
	return ret
}

func (i XInt64) RangeTo(to XInt64) XList[XInt64] {
	var ret XList[XInt64]
	for ii := i; ii <= to; ii++ {
		ret = append(ret, ii)
	}
	return ret
}

func (i XInt64) RangeStepTo(to XInt64, step XInt64) XList[XInt64] {
	var ret XList[XInt64]
	for ii := i; ii <= to; ii += step {
		ret = append(ret, ii)
	}
	return ret
}

func (i XInt64) DownTo(to XInt64) XList[XInt64] {
	var ret XList[XInt64]
	for ii := i; ii >= to; ii-- {
		ret = append(ret, ii)
	}
	return ret
}

func (i XInt64) DownStepTo(to XInt64, step XInt64) XList[XInt64] {
	var ret XList[XInt64]
	for ii := i; ii >= to; ii -= step {
		ret = append(ret, ii)
	}
	return ret
}

func (i XChar) RangeTo(to XChar) XList[XChar] {
	var ret XList[XChar]
	for ii := int64(i); ii <= int64(to); ii++ {
		ret = append(ret, XChar(ii))
	}
	return ret
}

func (i XChar) RangeStepTo(to XChar, step XInt64) XList[XChar] {
	var ret XList[XChar]
	for ii := int64(i); ii <= int64(to); ii += int64(step) {
		ret = append(ret, XChar(ii))
	}
	return ret
}

func (i XChar) DownTo(to XChar) XList[XChar] {
	var ret XList[XChar]
	for ii := int64(i); ii >= int64(to); ii-- {
		ret = append(ret, XChar(ii))
	}
	return ret
}

func (i XChar) DownStepTo(to XChar, step XInt64) XList[XChar] {
	var ret XList[XChar]
	for ii := int64(i); ii >= int64(to); ii -= int64(step) {
		ret = append(ret, XChar(ii))
	}
	return ret
}

func (i XInt) ToString() XString {
	return XString(fmt.Sprintf("%d", i))
}

func (i XInt64) ToString() XString {
	return XString(fmt.Sprintf("%d", i))
}

func (i XChar) ToString() XString {
	return XString(i)
}

func (i XChar) Code() XInt {
	return XInt(i)
}

func (i XFloat) ToString() XString {
	return XString(fmt.Sprintf("%f", i))
}

func (i XFloat64) ToString() XString {
	return XString(fmt.Sprintf("%f", i))
}

func (i XInt) RotateLeft(bitCount XInt) XInt {
	return XInt(i<<bitCount | i>>(8-bitCount))
}

func (i XInt) RotateRight(bitCount XInt) XInt {
	return XInt(i>>bitCount | i<<(8-bitCount))
}

func (i *XInt64) RotateLeft(bitCount XInt64) XInt64 {
	return XInt64((*i)<<bitCount | (*i)>>(64-bitCount))
}

func (i XInt64) RotateRight(bitCount XInt64) XInt64 {
	return XInt64(i>>bitCount | i<<(64-bitCount))
}

func (i XInt) ToHex() XString {
	return XString(fmt.Sprintf("%X", i))
}

func (i XInt64) ToHex() XString {
	return XString(fmt.Sprintf("%X", i))
}

func (i XInt) ToOct() XString {
	return XString(strconv.FormatInt(int64(i), 8))
}

func (i XInt64) ToOct() XString {
	return XString(strconv.FormatInt(int64(i), 8))
}

func (i XInt) ToBinary() XString {
	return XString(strconv.FormatInt(int64(i), 2))
}

func (i XInt64) ToBinary() XString {
	return XString(strconv.FormatInt(int64(i), 2))
}

func (i XChar) IsLetter() bool {
	return unicode.IsLetter(rune(i))
}

func (i XChar) IsDigit() bool {
	return unicode.IsDigit(rune(i))
}

func (i XChar) IsLetterOrDigit() bool {
	return unicode.IsLetter(rune(i)) || unicode.IsDigit(rune(i))
}

func (i XChar) IsSymbol() bool {
	return unicode.IsSymbol(rune(i))
}

func (i XChar) IsWhitespace() bool {
	return unicode.IsSpace(rune(i))
}

func (i XChar) ToUpper() XChar {
	return XChar(unicode.ToUpper(rune(i)))
}

func (i XChar) ToLower() XChar {
	return XChar(unicode.ToLower(rune(i)))
}

func (i XChar) ToTitle() XChar {
	return XChar(unicode.ToTitle(rune(i)))
}

func (i XChar) IsUpper() bool {
	return unicode.IsUpper(rune(i))
}

func (i XChar) IsLower() bool {
	return unicode.IsLower(rune(i))
}

func (i XChar) IsTitle() bool {
	return unicode.IsTitle(rune(i))
}

func (i XChar) IsISOControl() bool {
	return unicode.IsControl(rune(i))
}

func (i XChar) IsHighSurrogate() bool {
	return i >= MinHighSurrogate && i < (MaxHighSurrogate+1)
}

func (i XChar) IsLowSurrogate() bool {
	return i >= MinLowSurrogate && i < (MaxLowSurrogate+1)
}
