package numx

import (
	"fmt"
	"math"
)

const (
	EncodeType2       = "01"
	EncodeType8       = "01234567"
	EncodeType10      = "0123456789"
	EncodeType16Lower = "0123456789abcdef"
	EncodeType16Upper = "0123456789ABCDEF"
	EncodeType26Lower = "abcdefghijklmnopqrstuvwxyz"
	EncodeType26Upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	EncodeType32Lower = "0123456789abcdefghjkmnpqrstvwxyz"
	EncodeType32Upper = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
	EncodeType36Lower = "0123456789abcdefghijklmnopqrstuvwxyz"
	EncodeType36Upper = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	EncodeType52      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	EncodeType58      = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	EncodeType62      = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Encoding struct {
	runes []rune
	nmaps map[rune]int
}

func NewEncoding(encode_type string) *Encoding {
	r := &Encoding{
		runes: []rune(encode_type),
	}
	r.nmaps = make(map[rune]int, len(r.runes))
	for i := 0; i < len(r.runes); i++ {
		r.nmaps[r.runes[i]] = i
	}
	return r
}

func (e *Encoding) Encode(num int64) string {
	base := int64(len(e.runes))
	r := make([]rune, 0, 64)

	for {
		s, y := num/int64(base), num%int64(base)
		r = append(r, e.runes[y])
		if s < int64(base) {
			if s != 0 {
				r = append(r, e.runes[s])
			}
			break
		} else {
			num = s
		}
	}
	result := make([]rune, len(r))
	for i := len(r) - 1; i >= 0; i-- {
		result[len(r)-i-1] = r[i]
	}
	return string(result)
}

func (e *Encoding) Decode(val string) (int64, error) {
	var (
		base         = int64(len(e.runes))
		vrunes       = []rune(val)
		r      int64 = 0
		err    error
	)
	for i, vlen := 0, len(vrunes); i < vlen; i++ {
		cur_rune := vrunes[i]
		if num, ok := e.nmaps[cur_rune]; ok {
			r = r + int64(num)*int64(math.Pow(float64(base), float64(vlen)-float64(i+1)))
		} else {
			r = 0
			err = fmt.Errorf("invalid letter `%s`", string(cur_rune))
			break
		}
	}
	return r, err
}
