package main

import (
	b64 "encoding/base64"
	"flag"
	"fmt"
	"os"
)

var UsageText = `
Usage : 

  encdec -[action-option] -[algo-option] [string-input]

Where: 

  -[action-option] is either -e or -d to encode or decode the string
  -[algo-option] is the algorithm to use for encoding or decoding the string
  [string-input] is the string to encode or decode

Example:

  encdec -e -algo base64 "hello world"
  encdec -d -algo base64 "aGVsbG8gd29ybGQ="

Supported Algorithms:

  base64

For help or suggestions, please open an issue at - https://github.com/akashsharma99/encdec
`

// cli tool to take a string and encode it to base64
func main() {
	// define options as flags
	encode := flag.Bool("e", false, "Encode the string")
	decode := flag.Bool("d", false, "Decode the string")
	algo := flag.String("algo", "", "specify an algorithm")
	// parse the flags
	flag.Parse()
	inputString := flag.Arg(0)
	if (!*encode && !*decode) || (*encode && *decode) || *algo == "" || inputString == "" {
		fmt.Println("Invalid options")
		fmt.Println(UsageText)
		os.Exit(1)
	}
	// if encode is set then encode the string
	if *encode {
		encodedString := encodeString(inputString, *algo)
		fmt.Println(encodedString)
		os.Exit(0)
	}
	// if decode is set then decode the string
	if *decode {
		decodedString := decodeString(inputString, *algo)
		fmt.Println(decodedString)
		os.Exit(0)
	}

}
func encodeString(inputString string, algo string) string {
	if algo == "base64" {
		return encdecBase64(inputString, "e")
	}
	return ""
}
func decodeString(inputString string, algo string) string {
	if algo == "base64" {
		return encdecBase64(inputString, "d")
	}
	return ""
}
func encdecBase64(inputString string, action string) string {
	if action == "e" {
		return b64.StdEncoding.EncodeToString([]byte(inputString))
	} else if action == "d" {
		data, err := b64.StdEncoding.DecodeString(inputString)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		return string(data)
	}
	return ""
}
