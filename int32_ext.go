package ext

import "strconv"

type Int32Ext int32

func ToInt32(i int32) Int32Ext {
	return Int32Ext(i)
}

func (data Int32Ext) Int32() int32 {
	return int32(data)
}

func (data Int32Ext) ToString(base int) StringExt {
	s := strconv.FormatInt(int64(data.Int32()), base)
	return ToString(s)
}
