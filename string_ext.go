package ext

import (
	"encoding/base64"
	"regexp"
	"strconv"
	"strings"
)

type StringExt string

func ToString(str string) StringExt {
	return StringExt(str)
}

func (str StringExt) String() string {
	return string(str)
}

func (str StringExt) ToBytes() Bytes {
	return ToBytes([]byte(str))
}

func (str StringExt) ToFloat64() Float64Ext {
	return ToFloat64(str.Float64())
}

func (str StringExt) ToInt() IntExt {
	return ToInt(str.Int())
}

func (str StringExt) ToInt32(base int) Int32Ext {
	return ToInt32(str.Int32(base))
}

func (str StringExt) ToInt64(base int) Int64Ext {
	return ToInt64(str.Int64(base))
}

func (str StringExt) Length() int {
	return len([]rune(str.String()))
}

func (str StringExt) Substring(index int, len int) StringExt {
	r := []rune(str.String())[index : index+len]
	return ToString(string(r))
}

func (str StringExt) Left(n int) StringExt {
	r := []rune(str.String())[:n]
	return ToString(string(r))
}

func (str StringExt) Right(n int) StringExt {
	r := []rune(str.String())[str.Length()-n:]
	return ToString(string(r))
}

func (str StringExt) RemoveLeft(n int) StringExt {
	r := []rune(str.String())[n:]
	return ToString(string(r))
}

func (str StringExt) RemoveRight(n int) StringExt {
	r := []rune(str.String())[:str.Length()-n]
	return ToString(string(r))
}

func (str StringExt) In(ignoreCase bool, strs ...StringExt) bool {
	var returnValue bool = false
	for _, v := range strs {
		if ignoreCase {
			if str.ToLower() == v.ToLower() {
				returnValue = true
				break
			}
		} else {
			if str == v {
				returnValue = true
				break
			}
		}
	}
	return returnValue
}

func (str StringExt) In2(ignoreCase bool, strs ...string) bool {
	var returnValue bool = false
	for _, v := range strs {
		if ignoreCase {
			if str.ToLower().String() == strings.ToLower(v) {
				returnValue = true
				break
			}
		} else {
			if str.String() == v {
				returnValue = true
				break
			}
		}
	}
	return returnValue
}

func (str StringExt) In3(ignoreCase bool, strs []string) bool {
	var returnValue bool = false
	for _, v := range strs {
		if ignoreCase {
			if str.ToLower().String() == strings.ToLower(v) {
				returnValue = true
				break
			}
		} else {
			if str.String() == v {
				returnValue = true
				break
			}
		}
	}
	return returnValue
}

func (str StringExt) InLike(ignoreCase bool, strs ...StringExt) bool {
	var returnValue bool = false
	for _, v := range strs {
		if ignoreCase {
			if str.ToLower().Index(v.ToLower()) != -1 {
				returnValue = true
				break
			}
		} else {
			if str.Index(v) != -1 {
				returnValue = true
				break
			}
		}
	}
	return returnValue
}

func (str StringExt) InLike2(ignoreCase bool, strs ...string) bool {
	var returnValue bool = false
	for _, v := range strs {
		if ignoreCase {
			if str.ToLower().Index(ToString(v).ToLower()) != -1 {
				returnValue = true
				break
			}
		} else {
			if str.Index(ToString(v)) != -1 {
				returnValue = true
				break
			}
		}
	}
	return returnValue
}

func (str StringExt) InLike3(ignoreCase bool, strs []string) bool {
	var returnValue bool = false
	for _, v := range strs {
		if ignoreCase {
			if str.ToLower().Index(ToString(v).ToLower()) != -1 {
				returnValue = true
				break
			}
		} else {
			if str.Index(ToString(v)) != -1 {
				returnValue = true
				break
			}
		}
	}
	return returnValue
}

func (str StringExt) BetweenAnd(minValue StringExt, maxValue StringExt) bool {
	var returnValue bool = false
	if str.Compare(minValue) > -1 && str.Compare(maxValue) < 1 {
		returnValue = true
	}
	return returnValue
}

func (str StringExt) BetweenAnd2(minValue string, maxValue string) bool {
	var returnValue bool = false
	if str.Compare(ToString(minValue)) > -1 && str.Compare(ToString(maxValue)) < 1 {
		returnValue = true
	}
	return returnValue
}

func (str StringExt) IsInt() bool {
	var returnValue bool = false
	_, err := strconv.Atoi(str.String())
	if err == nil {
		returnValue = true
	}
	return returnValue
}

func (str StringExt) IsInt32(base int) bool {
	var returnValue bool = false
	_, err := strconv.ParseInt(str.String(), base, 32)
	if err == nil {
		returnValue = true
	}
	return returnValue
}

func (str StringExt) IsInt64(base int) bool {
	var returnValue bool = false
	_, err := strconv.ParseInt(str.String(), base, 64)
	if err == nil {
		returnValue = true
	}
	return returnValue
}

func (str StringExt) IsFloat64() bool {
	var returnValue bool = false
	_, err := strconv.ParseFloat(str.String(), 64)
	if err == nil {
		returnValue = true
	}
	return returnValue
}

func (str StringExt) Int() int {
	returnValue, err := strconv.Atoi(str.String())
	if err != nil {
		panic(err)
	}
	return returnValue
}

func (str StringExt) Int32(base int) int32 {
	r, err := strconv.ParseInt(str.String(), base, 32)
	if err != nil {
		panic(err)
	}
	returnValue := int32(r)
	return returnValue
}

func (str StringExt) Int64(base int) int64 {
	returnValue, err := strconv.ParseInt(str.String(), base, 64)
	if err != nil {
		panic(err)
	}
	return returnValue
}

func (str StringExt) Float64() float64 {
	returnValue, err := strconv.ParseFloat(str.String(), 64)
	if err != nil {
		panic(err)
	}
	return returnValue
}

func (str StringExt) Reverse() StringExt {
	runes := []rune(str.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return ToString(string(runes))
}

func (str StringExt) DecodeBase64() Bytes {
	d64, err := base64.StdEncoding.DecodeString(str.String())
	if err != nil {
		panic(err)
	}
	return ToBytes(d64)
}

func (str StringExt) RegexMatchString(s StringExt) bool {
	r := regexp.MustCompile(s.String())
	result := r.MatchString(str.String())
	return result
}

func (str StringExt) RegexFindString(s StringExt) StringExt {
	r := regexp.MustCompile(s.String())
	result := r.FindString(str.String())
	return ToString(result)
}

func (str StringExt) RegexFindStringIndex(s StringExt) []int {
	r := regexp.MustCompile(s.String())
	result := r.FindStringIndex(str.String())
	//
	if len(result) > 0 {
		prefix := []byte(str)[:result[0]]
		rs := []rune(string(prefix))
		result[0] = len(rs)
		//
		prefix = []byte(str)[:result[1]]
		rs = []rune(string(prefix))
		result[1] = len(rs)
	}
	//
	return result
}

func (str StringExt) RegexFindAllString(s StringExt, n int) []StringExt {
	r := regexp.MustCompile(s.String())
	result := r.FindAllString(str.String(), n)
	returnValue := make([]StringExt, len(result))
	for i := 0; i < len(result); i++ {
		returnValue[i] = ToString(result[i])
	}
	return returnValue
}

func (str StringExt) RegexFindAllStringIndex(s StringExt, n int) [][]int {
	r := regexp.MustCompile(s.String())
	result := r.FindAllStringIndex(str.String(), n)
	if len(result) > 0 {
		for i := 0; i < len(result); i++ {
			prefix := []byte(str)[:result[i][0]]
			rs := []rune(string(prefix))
			result[i][0] = len(rs)
			//
			prefix = []byte(str)[:result[i][1]]
			rs = []rune(string(prefix))
			result[i][1] = len(rs)
		}
	}
	return result
}

func (str StringExt) RegexReplaceAllString(s StringExt, repl StringExt) StringExt {
	r := regexp.MustCompile(s.String())
	result := r.ReplaceAllString(str.String(), repl.String())
	return ToString(result)
}

func (str StringExt) RegexReplaceAllStringFunc(s StringExt, repl func(string) string) StringExt {
	r := regexp.MustCompile(s.String())
	result := r.ReplaceAllStringFunc(str.String(), repl)
	return ToString(result)
}

func (str StringExt) DecodeUnicode() StringExt {
	returnValue := str.RegexReplaceAllStringFunc("\\\\u[a-zA-Z0-9]{4,4}", func(s string) string {
		i, err := strconv.ParseInt(s[2:], 16, 64)
		if err != nil {
			return s
		}
		return string(rune(i))
	})
	return returnValue
}

func (str StringExt) ExtractSingleString(beginStr StringExt, endStr StringExt) StringExt {
	returnValue := ToString("")
	result := str.ExtractStrings(beginStr, endStr, 1)
	if len(result) > 0 {
		returnValue = result[0]
	}
	return returnValue
}

func (str StringExt) ExtractStrings(beginStr StringExt, endStr StringExt, n int) []StringExt {
	returnValue := make([]StringExt, 0)
	strArray := str.RegexFindAllString(beginStr+".+?"+endStr, n)
	for i := 0; i < len(strArray); i++ {
		begin := strArray[i].RegexFindAllString(beginStr, -1)
		end := strArray[i].RegexFindAllString(endStr, -1)
		if strArray[i].Length() > 0 && len(begin) > 0 && len(end) > 0 {
			beginLen := begin[0].Length()
			endLen := end[len(end)-1].Length()
			if strArray[i].Length() >= beginLen+endLen {
				returnValue = append(returnValue, strArray[i].RemoveLeft(beginLen).RemoveRight(endLen))
			}
		}
	}
	return returnValue
}

//strings 基本語法

func (str StringExt) Compare(s StringExt) int {
	return strings.Compare(str.String(), s.String())
}

func (str StringExt) Contains(s StringExt) bool {
	return strings.Contains(str.String(), s.String())
}

func (str StringExt) Count(s StringExt) int {
	return strings.Count(str.String(), s.String())
}

func (str StringExt) HasPrefix(s StringExt) bool {
	return strings.HasPrefix(str.String(), s.String())
}

func (str StringExt) HasSuffix(s StringExt) bool {
	return strings.HasSuffix(str.String(), s.String())
}

func (str StringExt) Index(s StringExt) int {
	returnValue := strings.Index(str.String(), s.String())
	if returnValue >= 0 {
		prefix := []byte(str)[:returnValue]
		rs := []rune(string(prefix))
		returnValue = len(rs)
	}
	return returnValue
}

func (str StringExt) LastIndex(s StringExt) int {
	returnValue := strings.LastIndex(str.String(), s.String())
	if returnValue >= 0 {
		prefix := []byte(str)[:returnValue]
		rs := []rune(string(prefix))
		returnValue = len(rs)
	}
	return returnValue
}

func (str StringExt) RepeatLeft(s string, n int) StringExt {
	return ToString(strings.Repeat(s, n)) + str
}

func (str StringExt) RepeatRight(s string, n int) StringExt {
	return str + ToString(strings.Repeat(s, n))
}

func (str StringExt) Replace(old StringExt, new StringExt, n int) StringExt {
	return ToString(strings.Replace(str.String(), old.String(), new.String(), n))
}

func (str StringExt) Split(s StringExt) []string {
	return strings.Split(str.String(), s.String())
}

func (str StringExt) SplitAfter(s StringExt) []string {
	return strings.SplitAfter(str.String(), s.String())
}

func (str StringExt) SplitAfterN(s StringExt, n int) []string {
	return strings.SplitAfterN(str.String(), s.String(), n)
}

func (str StringExt) SplitN(s StringExt, n int) []string {
	return strings.SplitN(str.String(), s.String(), n)
}

func (str StringExt) Title() StringExt {
	return ToString(strings.Title(str.String()))
}

func (str StringExt) ToLower() StringExt {
	return ToString(strings.ToLower(str.String()))
}

func (str StringExt) ToTitle() StringExt {
	return ToString(strings.ToTitle(str.String()))
}

func (str StringExt) ToUpper() StringExt {
	return ToString(strings.ToUpper(str.String()))
}

func (str StringExt) Trim(s StringExt) StringExt {
	return ToString(strings.Trim(str.String(), s.String()))
}

func (str StringExt) TrimLeft(s StringExt) StringExt {
	return ToString(strings.TrimLeft(str.String(), s.String()))
}

func (str StringExt) TrimPrefix(s StringExt) StringExt {
	return ToString(strings.TrimPrefix(str.String(), s.String()))
}

func (str StringExt) TrimRight(s StringExt) StringExt {
	return ToString(strings.TrimRight(str.String(), s.String()))
}

func (str StringExt) TrimSpace() StringExt {
	return ToString(strings.TrimSpace(str.String()))
}

func (str StringExt) TrimSuffix(s StringExt) StringExt {
	return ToString(strings.TrimSuffix(str.String(), s.String()))
}
