package main

import (
	"fmt"
	"log"

	"github.com/seehuhn/fortuna"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	length     = kingpin.Flag("length", "Password length").Short('l').Default("32").Int()
	count      = kingpin.Flag("count", "Numbers of passwords").Short('c').Default("6").Int()
	dictionary = kingpin.Flag("dictionary", "Dictionary what should be used").Short('d').
			Default(`ABCDEFGHIKLMNOPQRSTVXYZabcdefghiklmnopqrstvxyz0123456789`).String()
	entropy = kingpin.Flag("entropy", "Using entropy for generation passwords").Short('e').String()
)

func main() {
	// Initialize kingpin command line parser
	kingpin.Version("Version: " + version + "\nBuildTime: " + buildTime + "\nCommit: " + commit)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	// Create new fortuna generator
	rng, err := fortuna.NewRNG(*entropy)
	if err != nil {
		log.Fatal("cannot initialise the RNG: " + err.Error())
	}
	defer rng.Close()

	// Create *count passwords from *dictionary
	for i := 0; i < *count; i++ {
		data := rng.RandomData(uint(*length))
		fmt.Printf("%v\n", encode(data))
	}
}

func encode(data []byte) string {
	result := ""
	for _, el := range data {
		l := int(el)
		if l > len(*dictionary)-1 {
			l = mod(int(el), len(*dictionary)-1)
		}
		result += string((*dictionary)[l])
	}
	return result
}

func mod(a, b int) int {
	for a > b {
		a -= b
	}
	return a
}
