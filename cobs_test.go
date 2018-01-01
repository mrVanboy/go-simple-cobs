package cobs

import (
	"testing"
	"reflect"
	"math/rand"
	"encoding/hex"
)

type testpair struct {
	source []byte
	encoded []byte
}

var array253 = buildArrayWithoutNulls(253)
var array254 = buildArrayWithoutNulls(254)
var array255 = append(array254, 0x12)
var array256 = append([]byte {0x00}, array255...)

var tests = []testpair{
	// Tests cases for single packet arrays
	{[]byte{},						[]byte{0x01}},
	{[]byte{0x00},					[]byte{0x01, 0x01}},
	{[]byte{0x00, 0x00},				[]byte{0x01, 0x01, 0x01}},
	{[]byte{0x11, 0x22, 0x00, 0x33},	[]byte{0x03, 0x11, 0x22, 0x02, 0x33}},
	{[]byte{0x11, 0x22, 0x33, 0x44},	[]byte{0x05, 0x11, 0x22, 0x33, 0x44}},
	{[]byte{0x11, 0x00, 0x00, 0x00},	[]byte{0x02, 0x11, 0x01, 0x01, 0x01}},

	// Tests cases for length of packet near 254
	{array253,						append( []byte{0xFE}, 		array253... )},
	{array254,						append( []byte{0xFF}, 		array254... )},
	{array255,						append( []byte{0xFF}, 		append( array254, []byte{0x02, 0x12}... )... )},
	{array256,						append( []byte{0x01, 0xFF}, append( array254, []byte{0x02, 0x12}... )... )},
}

func TestPredefinedEncode(t *testing.T){
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
				"For \n", hex.Dump(pair.source),
				"got \n", hex.Dump(output),
				"expected \n", hex.Dump(pair.encoded),
			)
		}
	}
}

func TestPredefinedDecode(t *testing.T) {
	for _, pair := range tests {
		output, err := Decode(pair.encoded)
		if err != nil {
			t.Error(
				"For", hex.Dump(pair.encoded),
				"got error", err,
			)

		} else if !reflect.DeepEqual(output, pair.source) {
			t.Error(
				"For \n", hex.Dump(pair.encoded),
				"got \n", hex.Dump(output),
				"expected \n", hex.Dump(pair.source),
			)
		}
	}
}

func TestTwoWayEncoding(t *testing.T) {
	for i := 0; i < 100; i++ {
		data := buildArrayWithNulls(i*10)
		encoded, _ := Encode(data)
		decoded, err := Decode(encoded)
		if err != nil {
			t.Error(
				"For", hex.Dump(encoded),
				"got error", err,
			)

		} else if !reflect.DeepEqual(data, decoded) {
			t.Error(
				"Source data \n", hex.Dump(data),
				"encoded to \n", hex.Dump(encoded),
				"decoded to \n", hex.Dump(decoded),
			)
		}
	}
}

func TestDecodeNil(t *testing.T) {
	var data []byte
	output, err := Decode(data)
	if err == nil {
		t.Error("For nil input data expect error, but get: ", hex.Dump(output))
	}
}

func TestDecodeWithZeros(t *testing.T) {
	var data = buildArrayWithNulls(10)
	encoded, err := Encode(data)
	if err != nil {
		t.Error(
			"For", hex.Dump(encoded),
			"got error", err,
		)
		return
	}

	// Add zero at random place in encoded slice
	encoded[ rand.Intn( len( encoded ) - 1 ) + 1 ] = 0x00
	decoded, err := Decode(encoded)
	if err == nil {
			t.Error(
			"For encoded data with zero \n", hex.Dump(encoded),
			"get \n", hex.Dump(decoded),
			"expect error",
		)

	}
}

func buildArrayWithoutNulls(length int) []byte {
	output := make([]byte, length)
	for i := 0; i < length; i++ {
		r := rand.Intn(254) + 1
		output[i] = byte(r)
	}
	return output[:]
}

func buildArrayWithNulls(length int) []byte {
	output := make([]byte, length)
	for i := 0; i < length; i++ {
		r := rand.Intn(255)
		output[i] = byte(r)
	}
	return output[:]
}
