package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lil-shimon/toudi"
)

const (
	toudisFile = "toudis.json"
)

func main() {

	add := flag.Bool("add", false, "add a new toudis")
	complete := flag.Int("complete", 0, "mark a todo as completed")

	flag.Parse()

	toudis := &toudi.Toudis{}

	if err := toudis.Load(toudisFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		toudis.Add("Test")
		err := toudis.Store(toudisFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *complete > 0:
		err := toudis.Complete(*complete)
		err = toudis.Store(toudisFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}

}
