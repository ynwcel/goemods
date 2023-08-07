package numx

import (
	"fmt"
	"testing"
)

func testEncode(t *testing.T, data int64) {
	types := map[string]string{
		"EncodeType2":       EncodeType2,
		"EncodeType8":       EncodeType8,
		"EncodeType16Lower": EncodeType16Lower,
		"EncodeType16Upper": EncodeType16Upper,
		"EncodeType26Lower": EncodeType26Lower,
		"EncodeType26Upper": EncodeType26Upper,
		"EncodeType32Lower": EncodeType32Lower,
		"EncodeType32Upper": EncodeType32Upper,
		"EncodeType36Lower": EncodeType36Lower,
		"EncodeType36Upper": EncodeType36Upper,
		"EncodeType52":      EncodeType52,
		"EncodeType58":      EncodeType58,
		"EncodeType62":      EncodeType62,
	}
	for tpe_name, tpe := range types {
		e := NewEncoding(tpe)
		v := e.Encode(data)
		r, err := e.Decode(v)
		if err != nil {
			t.Error(err)
		}
		tpe_name = fmt.Sprintf("[ %-20s ]", tpe_name)
		t.Log(tpe_name, "source =", data, "encode = ", v, "decode = ", r)
	}
}

func TestAll(t *testing.T) {
	datas := []int64{
		0, 1, 2, 4, 8,
		16, 32, 36, 52, 58, 62,
		202308070950,
		1257894000,
	}
	for _, v := range datas {
		testEncode(t, v)
	}
}
