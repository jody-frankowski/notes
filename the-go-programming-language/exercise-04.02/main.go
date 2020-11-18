// Exercise 4.2: Write a program that prints the SHA256 hash of its standard input by default but
// supports a command-line flag to print the SHA384 or SHA512 hash instead.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when reading from stdin: %v", err)
		os.Exit(-1)
	}
	sha512Flag := flag.Bool("sha512", false, "Use SHA512")
	sha384Flag := flag.Bool("sha384", false, "Use SHA384")
	flag.Parse()
	if *sha512Flag {
		fmt.Printf("%x\n", sha512.Sum512(bytes))
	} else if *sha384Flag {
		fmt.Printf("%x\n", sha512.Sum384(bytes))
	} else {
		fmt.Printf("%x\n", sha256.Sum256(bytes))
	}

}
