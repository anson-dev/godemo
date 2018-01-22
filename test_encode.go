package main

import "encoding/base64"
import "os"

func main() {

	input := []byte("123456")
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(input)

	// Must close the encoder when finished to flush any partial blocks.
	// If you comment out the following line, the last partial block "r"
	// won't be encoded.
	encoder.Close()
	
}