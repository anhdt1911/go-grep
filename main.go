package main

import "github.com/docopt/docopt-go"

const (
	Usage = "\nUsage: rep [options] <pattern> [<file> ...]\n"
)

func main() {
	args, err := docopt.ParseDoc(Usage)
	println(args, err)
}
