package ext

import "strconv"

type Int64Ext int64

func ToInt64(i int64) Int64Ext {
	return Int64Ext(i)
}

func (data Int64Ext) Int64() int64 {
	return int64(data)
}

func (data Int64Ext) ToString(base int) StringExt {
	s := strconv.FormatInt(data.Int64(), base)
	return ToString(s)
}
