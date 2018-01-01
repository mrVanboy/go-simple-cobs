package cobs

import "errors"

func Encode(input []byte) ([]byte, error) {
	// Initialize output array. If input array is empty, then output array will contain {0x01}
	var output = []byte {0x01}

	// Index of the last zero element
	var lastZeroIndex = 0

	// Difference between lastZeroIndex and current element
	var delta byte = 1

	// Iterate over input array
	for i := range input {

		if input[i] == 0x00 {

			output[lastZeroIndex] = delta
			output = append(output, 0x01)
			lastZeroIndex = len(output) - 1
			delta = 1

		} else {

			if delta == 255 {

				output[lastZeroIndex] = delta
				output = append(output, 0x01)
				lastZeroIndex = len(output) - 1
				delta = 1

			}

			output = append(output, input[i])
			delta++
		}
	}

	output[lastZeroIndex] = delta
	return output, nil
}

func Decode(input []byte) ([]byte, error){

	if len(input) == 0 {
		return nil, errors.New("cobs decode: length of input array must be greater than zero")
	}

	var output = make([]byte, 0)
	var lastZeroIndex = 0

	for true {

		if lastZeroIndex == len(input){
			break
		}

		if input[lastZeroIndex] == 0x00 {
			return nil, errors.New("cobs decode: zero are not allowed in input array")
		}

		if int(input[lastZeroIndex]) > (len(input) - int(lastZeroIndex)){
			return nil, errors.New("cobs decode: zero value refer to outbound of array")
		}

		var nextZeroIndex = lastZeroIndex + int(input[lastZeroIndex])

		for i:=lastZeroIndex + 1; i < nextZeroIndex; i++ {

			if input[i] == 0x00 {
				return nil, errors.New("cobs decode: zero value are not allowed in input array")
			}

			output = append(output, input[i])
		}

		if nextZeroIndex < len(input) && input[lastZeroIndex] != 0xFF{
			output = append(output, 0x00)
		}

		lastZeroIndex = nextZeroIndex
	}

	return output, nil
}