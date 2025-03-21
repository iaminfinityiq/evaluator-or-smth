package value_types

import (
	"strconv"
)

type Fraction struct {
	Numerator int64
	Denominator int64
}

func (f Fraction) String() string {
	return strconv.FormatInt(f.Numerator, 10) + "/" + strconv.FormatInt(f.Denominator, 10)
}