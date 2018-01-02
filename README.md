
# Simple COBS (Consistent Overhead Byte Stuffing) [*ʷᶦᵏᶦ*](https://en.wikipedia.org/wiki/Consistent_Overhead_Byte_Stuffing)

Simple library for encoding and decoding byte slices with COBS algorithm.


Example of usage:

```go
package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/mrVanboy/go-simple-cobs"
)

func main() {
	data := []byte{0x01, 0x02}
	fmt.Println("Source data:")
	fmt.Println(hex.Dump(data))

	encoded, err := cobs.Encode(data)
	if err != nil {
		fmt.Printf("Error occured %s \n", err)
		os.Exit(-1)
	}
	fmt.Println("Encoded data:")
	fmt.Println(hex.Dump(encoded))

	decoded, err := cobs.Decode(encoded)
	if err != nil {
		fmt.Printf("Error occured %s \n", err)
		os.Exit(-1)
	}
	fmt.Println("Decoded data:")
	fmt.Println(hex.Dump(decoded))
}
```