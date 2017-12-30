package cobs

func Encode(input []byte) ([]byte, error) {
	// Initialize output array. If input array is empty, then output array will contain {0x01}
	var output = []byte {0x01}

	// Index of the last zero element
	var lastZeroIndex = 0

	// Difference between lastZeroIndex and current element
	var delta byte = 1

	// Iterate over input array
	for i := range input {
		// TODO: check if delta < 255
		if input[i] == 0x00 {
			output[lastZeroIndex] = delta
			output = append(output, 0x01)
			lastZeroIndex = len(output) - 1
			delta = 1
		} else {
			output = append(output, input[i])
			delta++
		}
	}
	output[lastZeroIndex] = delta
	return output, nil
}