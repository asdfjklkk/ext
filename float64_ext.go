package ext

import "strconv"

type Float64Ext float64

func ToFloat64(f float64) Float64Ext {
	return Float64Ext(f)
}

func (data Float64Ext) Float64() float64 {
	return float64(data)
}

func (data Float64Ext) ToString(fmt byte, prec int, bitSize int) StringExt {
	s := strconv.FormatFloat(data.Float64(), fmt, prec, bitSize)
	return ToString(s)
}
