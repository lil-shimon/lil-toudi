package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/lil-shimon/toudi"
)

const (
	toudisFile = "toudis.json"
)

func main() {

	add := flag.Bool("add", false, "add a new toudis")
	complete := flag.Int("complete", 0, "mark a toudis as completed")
	del := flag.Int("d", 0, "delete a toudis")
	list := flag.Bool("list", false, "list all toudis")
	l := flag.Bool("l", false, "list all toudis")

	flag.Parse()

	toudis := &toudi.Toudis{}

	if err := toudis.Load(toudisFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		toudi, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		toudis.Add(toudi)

		err = toudis.Store(toudisFile)
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
	case *del > 0:
		err := toudis.Delete(*del)
		err = toudis.Store(toudisFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *list:
	case *l:
		toudis.Print()
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}

}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	text := scanner.Text()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	if len(text) == 0 {
		return "", errors.New("empty input")
	}

	return text, nil
}
