package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	length     = kingpin.Flag("length", "Password length").Short('l').Default("32").Int()
	count      = kingpin.Flag("count", "Numbers of passwords").Short('c').Default("6").Int()
	dictionary = kingpin.Flag("dictionary", "Dictionary what should be used on generation passwords").Short('d').String()
	entropy    = kingpin.Flag("entropy", "Using entropy for generation passwords").Short('e').Bool()
)

func main() {
	kingpin.Version("Version: " + version + "\nBuildTime: " + buildTime + "\nCommit: " + commit)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
}
