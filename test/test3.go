package main

import "fmt"

func main() {
	// Define the byte array
	byteArray := []byte{123, 34, 99, 108, 105, 101, 110, 116, 95, 105, 100, 34, 58, 34, 107, 97, 120, 111, 49, 112, 117, 56, 116, 108, 56, 98, 105, 57, 56, 112, 48, 57, 119, 104, 106, 118, 111, 51, 103, 48, 119, 106, 101, 50, 49, 34, 44, 34, 99, 108, 105, 101, 110, 116, 95, 115, 101, 99, 114, 101, 116, 34, 58, 34, 48, 55, 99, 98, 51, 102, 55, 52, 49, 48, 56, 56, 100, 102, 52, 51, 57, 53, 99, 48, 49, 52, 101, 52, 55, 97, 49, 101, 57, 50, 57, 52, 53, 102, 48, 57, 102, 50, 57, 99, 102, 55, 97, 100, 56, 98, 51, 51, 99, 48, 50, 102, 52, 55, 49, 54, 48, 98, 98, 52, 99, 53, 50, 55, 34, 44, 34, 114, 101, 100, 105, 114, 101, 99, 116, 95, 117, 114, 105, 34, 58, 34, 104, 116, 116, 112, 115, 58, 47, 47, 99, 97, 114, 116, 45, 101, 100, 105, 116, 105, 110, 103, 45, 115, 116, 97, 103, 105, 110, 103, 46, 103, 114, 105, 116, 46, 115, 111, 102, 116, 119, 97, 114, 101, 47, 105, 110, 115, 116, 97, 108, 108, 34, 44, 34, 103, 114, 97, 110, 116, 95, 116, 121, 112, 101, 34, 58, 34, 97, 117, 116, 104, 111, 114, 105, 122, 97, 116, 105, 111, 110, 95, 99, 111, 100, 101, 34, 44, 34, 99, 111, 100, 101, 34, 58, 34, 50, 120, 121, 108, 55, 104, 105, 120, 108, 57, 53, 98, 112, 111, 117, 99, 108, 108, 106, 107, 113, 111, 114, 104, 57, 119, 116, 103, 104, 57, 50, 34, 44, 34, 115, 99, 111, 112, 101, 34, 58, 34, 115, 116, 111, 114, 101, 95, 97, 112, 112, 95, 101, 120, 116, 101, 110, 115, 105, 111, 110, 115, 95, 109, 97, 110, 97, 103, 101, 32, 115, 116, 111, 114, 101, 95, 99, 97, 114, 116, 32, 115, 116, 111, 114, 101, 95, 99, 104, 97, 110, 110, 101, 108, 95, 108, 105, 115, 116, 105, 110, 103, 115, 95, 114, 101, 97, 100, 95, 111, 110, 108, 121, 32, 115, 116, 111, 114, 101, 95, 99, 104, 97, 110, 110, 101, 108, 95, 115, 101, 116, 116, 105, 110, 103, 115, 95, 114, 101, 97, 100, 95, 111, 110, 108, 121, 32, 115, 116, 111, 114, 101, 95, 99, 104, 101, 99, 107, 111, 117, 116, 32, 115, 116, 111, 114, 101, 95, 99, 111, 110, 116, 101, 110, 116, 95, 99, 104, 101, 99, 107, 111, 117, 116, 32, 115, 116, 111, 114, 101, 95, 105, 110, 118, 101, 110, 116, 111, 114, 121, 95, 114, 101, 97, 100, 95, 111, 110, 108, 121, 32, 115, 116, 111, 114, 101, 95, 108, 111, 99, 97, 116, 105, 111, 110, 115, 95, 114, 101, 97, 100, 95, 111, 110, 108, 121, 32, 115, 116, 111, 114, 101, 95, 108, 111, 103, 115, 95, 114, 101, 97, 100, 95, 111, 110, 108, 121, 32, 115, 116, 111, 114, 101, 95, 112, 97, 121, 109, 101, 110, 116, 115, 95, 109, 101, 116, 104, 111, 100, 115, 95, 114, 101, 97, 100, 32, 115, 116, 111, 114, 101, 95, 115, 105, 116, 101, 115, 95, 114, 101, 97, 100, 95, 111, 110, 108, 121, 32, 115, 116, 111, 114, 101, 95, 115, 116, 111, 114, 101, 100, 95, 112, 97, 121, 109, 101, 110, 116, 95, 105, 110, 115, 116, 114, 117, 109, 101, 110, 116, 115, 95, 114, 101, 97, 100, 95, 111, 110, 108, 121, 32, 115, 116, 111, 114, 101, 95, 115, 116, 111, 114, 101, 102, 114, 111, 110, 116, 95, 97, 112, 105, 32, 115, 116, 111, 114, 101, 95, 115, 116, 111, 114, 101, 102, 114, 111, 110, 116, 95, 97, 112, 105, 95, 99, 117, 115, 116, 111, 109, 101, 114, 95, 105, 109, 112, 101, 114, 115, 111, 110, 97, 116, 105, 111, 110, 32, 115, 116, 111, 114, 101, 95, 116, 104, 101, 109, 101, 115, 95, 109, 97, 110, 97, 103, 101, 32, 115, 116, 111, 114, 101, 95, 118, 50, 95, 99, 111, 110, 116, 101, 110, 116, 32, 115, 116, 111, 114, 101, 95, 118, 50, 95, 99, 117, 115, 116, 111, 109, 101, 114, 115, 95, 114, 101, 97, 100, 95, 111, 110, 108, 121, 32, 115, 116, 111, 114, 101, 95, 118, 50, 95, 100, 101, 102, 97, 117, 108, 116, 32, 115, 116, 111, 114, 101, 95, 118, 50, 95, 105, 110, 102, 111, 114, 109, 97, 116, 105, 111, 110, 95, 114, 101, 97, 100, 95, 111, 110, 108, 121, 32, 115, 116, 111, 114, 101, 95, 118, 50, 95, 109, 97, 114, 107, 101, 116, 105, 110, 103, 32, 115, 116, 111, 114, 101, 95, 118, 50, 95, 111, 114, 100, 101, 114, 115, 95, 114, 101, 97, 100, 95, 111, 110, 108, 121, 32, 115, 116, 111, 114, 101, 95, 118, 50, 95, 112, 114, 111, 100, 117, 99, 116, 115, 32, 115, 116, 111, 114, 101, 95, 118, 50, 95, 116, 114, 97, 110, 115, 97, 99, 116, 105, 111, 110, 115, 95, 114, 101, 97, 100, 95, 111, 110, 108, 121, 34, 44, 34, 99, 111, 110, 116, 101, 120, 116, 34, 58, 34, 115, 116, 111, 114, 101, 115, 47, 50, 48, 113, 119, 52, 100, 105, 55, 106, 117, 34, 125}

	// Convert the byte array to a string
	str := string(byteArray)

	// Print the resulting string
	fmt.Println(str)
}