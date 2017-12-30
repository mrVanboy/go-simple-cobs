package cobs

import (
	"testing"
	"reflect"
)

type testpair struct {
	source []byte
	encoded []byte
}

var tests = []testpair{
	{[]byte{},						[]byte{0x01}},
	{[]byte{0x00},					[]byte{0x01, 0x01}},
	{[]byte{0x00, 0x00},				[]byte{0x01, 0x01, 0x01}},
	{[]byte{0x11, 0x22, 0x00, 0x33},	[]byte{0x03, 0x11, 0x22, 0x02, 0x33}},
	{[]byte{0x11, 0x22, 0x33, 0x44},	[]byte{0x05, 0x11, 0x22, 0x33, 0x44}},
	{[]byte{0x11, 0x00, 0x00, 0x00},	[]byte{0x02, 0x11, 0x01, 0x01, 0x01}},

}

func TestPreddefinedEncode(t *testing.T){
	for _, pair := range tests {
		output, err := Encode(pair.source)
		if err != nil {
			t.Error(
				"For", pair.source,
				"got error", err,
			)
		}
		if !reflect.DeepEqual(output, pair.encoded) {
			t.Error(
				"For", pair.source,
				"got",output,
				"expected", pair.encoded,
			)
		}
	}
}
