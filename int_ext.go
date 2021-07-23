package ext

import "strconv"

type IntExt int

func ToInt(i int) IntExt {
	return IntExt(i)
}

func (data IntExt) Int() int {
	return int(data)
}

func (data IntExt) ToString() StringExt {
	s := strconv.Itoa(data.Int())
	return ToString(s)
}
