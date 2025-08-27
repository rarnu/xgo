package xgo

import (
	"regexp"
	"strconv"
	"strings"
)

type XString string

func (s XString) At(index XInt) XChar {
	return XChar(s[index])
}

func (s XString) Length() XInt {
	return XInt(len(s))
}

func (s XString) Chars() XList[XChar] {
	var list XList[XChar]
	for i := 0; i < len(s); i++ {
		list = append(list, XChar(s[i]))
	}
	return list
}

func (s XString) Count(substr XString) XInt {
	return XInt(strings.Count(string(s), string(substr)))
}

func (s XString) Contains(substr XString) bool {
	return strings.Contains(string(s), string(substr))
}

func (s XString) ContainsAny(chars XString) bool {
	return strings.ContainsAny(string(s), string(chars))
}

func (s XString) ContainsRune(r XChar) bool {
	return strings.ContainsRune(string(s), rune(r))
}

func (s XString) IndexOf(substr XString) XInt {
	return XInt(strings.Index(string(s), string(substr)))
}

func (s XString) LastIndexOf(substr XString) XInt {
	return XInt(strings.LastIndex(string(s), string(substr)))
}

func (s XString) IndexByteOf(c byte) XInt {
	return XInt(strings.IndexByte(string(s), c))
}

func (s XString) IndexOfAny(chars XString) XInt {
	return XInt(strings.IndexAny(string(s), string(chars)))
}

func (s XString) LastIndexOfAny(chars XString) XInt {
	return XInt(strings.LastIndexAny(string(s), string(chars)))
}

func (s XString) LastIndexOfByte(c byte) XInt {
	return XInt(strings.LastIndexByte(string(s), c))
}

func (s XString) SplitN(sep XString, n XInt) XList[XString] {
	ss := strings.SplitN(string(s), string(sep), int(n))
	var ret XList[XString]
	for _, item := range ss {
		ret = append(ret, XString(item))
	}
	return ret
}

func (s XString) SplitAfterN(sep XString, n XInt) XList[XString] {
	ss := strings.SplitAfterN(string(s), string(sep), int(n))
	var ret XList[XString]
	for _, item := range ss {
		ret = append(ret, XString(item))
	}
	return ret
}

func (s XString) Split(sep XString) XList[XString] {
	ss := strings.Split(string(s), string(sep))
	var ret XList[XString]
	for _, item := range ss {
		ret = append(ret, XString(item))
	}
	return ret
}

func (s XString) SplitAfter(sep XString) XList[XString] {
	ss := strings.SplitAfter(string(s), string(sep))
	var ret XList[XString]
	for _, item := range ss {
		ret = append(ret, XString(item))
	}
	return ret
}

func (s XString) Fields() XList[XString] {
	ss := strings.Fields(string(s))
	var ret XList[XString]
	for _, item := range ss {
		ret = append(ret, XString(item))
	}
	return ret
}

func (s XString) FieldsFunc(f func(rune) bool) XList[XString] {
	ss := strings.FieldsFunc(string(s), f)
	var ret XList[XString]
	for _, item := range ss {
		ret = append(ret, XString(item))
	}
	return ret
}

func (s XString) StartsWith(prefix XString) bool {
	return strings.HasPrefix(string(s), string(prefix))
}

func (s XString) EndsWith(suffix XString) bool {
	return strings.HasSuffix(string(s), string(suffix))
}

func (s XString) Repeat(count XInt) XString {
	return XString(strings.Repeat(string(s), int(count)))
}

func (s XString) PadStart(length XInt, padChar XChar) XString {
	if length < 0 || int(length) <= len(s) {
		return s
	}
	var sb strings.Builder
	sb.Grow(int(length))
	for i := 0; i < int(length)-len(s); i++ {
		sb.WriteRune(rune(padChar))
	}
	sb.WriteString(string(s))
	return XString(sb.String())
}

func (s XString) PadEnd(length XInt, padChar XChar) XString {
	if length < 0 || int(length) <= len(s) {
		return s
	}
	var sb strings.Builder
	sb.Grow(int(length))
	sb.WriteString(string(s))
	for i := 0; i < int(length)-len(s); i++ {
		sb.WriteRune(rune(padChar))
	}
	return XString(sb.String())
}

func (s XString) TrimLeftFunc(f func(rune) bool) XString {
	return XString(strings.TrimLeftFunc(string(s), f))
}

func (s XString) TrimRightFunc(f func(rune) bool) XString {
	return XString(strings.TrimRightFunc(string(s), f))
}

func (s XString) TrimFunc(f func(rune) bool) XString {
	return XString(strings.TrimFunc(string(s), f))
}

func (s XString) IndexOfFunc(f func(rune) bool) XInt {
	return XInt(strings.IndexFunc(string(s), f))
}

func (s XString) LastIndexOfFunc(f func(rune) bool) XInt {
	return XInt(strings.LastIndexFunc(string(s), f))
}

func (s XString) Trim(cutset XString) XString {
	return XString(strings.Trim(string(s), string(cutset)))
}

func (s XString) TrimLeft(cutset XString) XString {
	return XString(strings.TrimLeft(string(s), string(cutset)))
}

func (s XString) TrimRight(cutset XString) XString {
	return XString(strings.TrimRight(string(s), string(cutset)))
}

func (s XString) TrimSpace() XString {
	return XString(strings.TrimSpace(string(s)))
}

func (s XString) TrimPrefix(prefix XString) XString {
	return XString(strings.TrimPrefix(string(s), string(prefix)))
}

func (s XString) TrimSuffix(suffix XString) XString {
	return XString(strings.TrimSuffix(string(s), string(suffix)))
}

func (s XString) Replace(old, new XString, n XInt) XString {
	return XString(strings.Replace(string(s), string(old), string(new), int(n)))
}

func (s XString) ReplaceAll(old, new XString) XString {
	return XString(strings.ReplaceAll(string(s), string(old), string(new)))
}

func (s XString) EqualFold(t XString) bool {
	return strings.EqualFold(string(s), string(t))
}

func (s XString) ToUpper() XString {
	return XString(strings.ToUpper(string(s)))
}

func (s XString) ToLower() XString {
	return XString(strings.ToLower(string(s)))
}

func (s XString) ToTitle() XString {
	return XString(strings.ToTitle(string(s)))
}

func (s XString) IsEmpty() bool {
	return len(s) == 0
}

func (s XString) IsNotEmpty() bool {
	return len(s) != 0
}

func (s XString) SubStringStart(AStartIndex XInt) XString {
	return s[AStartIndex:]
}

func (s XString) SubStringStartEnd(AStartIndex, AEndIndex XInt) XString {
	return s[AStartIndex:AEndIndex]
}

func (s XString) SubStringBefore(delimiter XString) XString {
	if i := s.IndexOf(delimiter); i != -1 {
		return s[:i]
	} else {
		return s
	}
}

func (s XString) SubStringAfter(delimiter XString) XString {
	if i := s.IndexOf(delimiter); i != -1 {
		return s[int(i)+len(delimiter):]
	} else {
		return s
	}
}

func (s XString) SubStringBeforeLast(delimiter XString) XString {
	if i := s.LastIndexOf(delimiter); i != -1 {
		return s[:i]
	} else {
		return s
	}
}

func (s XString) SubStringAfterLast(delimiter XString) XString {
	if i := s.LastIndexOf(delimiter); i != -1 {
		return s[int(i)+len(delimiter):]
	} else {
		return s
	}
}

func (s XString) Insert(index XInt, substr XString) XString {
	ss := s[:index] + substr + s[index:]
	return ss
}

func (s XString) Delete(index XInt, count XInt) XString {
	ss := s[:index] + s[index+count:]
	return ss
}

func (s XString) Matches(pattern XString) bool {
	reg := regexp.MustCompile(string(pattern))
	return reg.MatchString(string(s))
}

func (s XString) Lines() XList[XString] {
	return s.Split("\n")
}

func (s XString) LinesNoEmpty() XList[XString] {
	return s.Split("\n").Filter(func(item XString) bool {
		return item.IsNotEmpty()
	})
}

func (s XString) ToBoolean() bool {
	return s.ToLower() == "true"
}

func (s XString) ToInt() XInt {
	if i, err := strconv.Atoi(string(s)); err == nil {
		return XInt(i)
	} else {
		return 0
	}
}

func (s XString) ToInt64() XInt64 {
	if i, err := strconv.ParseInt(string(s), 10, 64); err == nil {
		return XInt64(i)
	} else {
		return 0
	}
}

func (s XString) ToFloat() XFloat {
	if f, err := strconv.ParseFloat(string(s), 64); err == nil {
		return XFloat(f)
	} else {
		return 0.0
	}
}

func (s XString) ToFloat64() XFloat64 {
	if f, err := strconv.ParseFloat(string(s), 64); err == nil {
		return XFloat64(f)
	} else {
		return 0.0
	}
}

func (s XString) Drop(n XInt) XString {
	return s[n:]
}

func (s XString) DropLast(n XInt) XString {
	return s[:len(s)-int(n)]
}

func (s XString) Take(n XInt) XString {
	return s[:n]
}

func (s XString) TakeLast(n XInt) XString {
	return s[len(s)-int(n):]
}
