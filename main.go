package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	fmtString string = "%v"
	incr      string
	delim     string = "\n"
	float     bool
	stderr    = log.New(os.Stderr, "", 0)
)

func init() {
	flag.Usage = func() {
		var def bytes.Buffer
		flag.CommandLine.SetOutput(&def)
		flag.PrintDefaults()

		stderr.Printf(
			mandoc,
			filepath.Base(os.Args[0]),
			def.String(),
			buildVersion,
			compiledBy,
			buildTimestamp,
		)
	}

	flag.Func("fmt", "use `fmt` to format the output", func(given string) error {
		given = strings.ReplaceAll(given, `\t`, "\t")
		given = strings.ReplaceAll(given, `\n`, "\n")
		if len(given) < 2 {
			stderr.Println(`using default -fmt:`, `%v`)
			return nil
		}
		fmtString = given
		return nil
	})
	flag.BoolVar(&float, "float", false, "treat inputs as floats")
	flag.StringVar(&incr, "incr", "1", "increment by `n`")
	flag.Func("delim", "separate numbers using `delim`", func(given string) error {
		given = strings.ReplaceAll(given, `\t`, "\t")
		given = strings.ReplaceAll(given, `\n`, "\n")
		delim = given
		return nil
	})

	flag.Parse()

	if flag.NArg() < 2 {
		log.Println(os.Args[0], "requires a start and stop")
		os.Exit(1)
	}
}

func main() {
	if float {
		loopFloat()
	} else {
		loopInt()
	}
}

const mandoc = `
NAME
  %[1]s - generate a sequence of numbers

SYNOPSIS
  %[1]s [-float] [-fmt "%%v"] [-incr 1] [-delim "\n"] start stop

DESCRIPTION
  %[1]s generates an inclusive sequence of numbers (ints or floats) that are
  printf formatted.

  The syntax for -fmt is Go's fmt.printf syntax. See also
  https://golang.org/pkg/fmt/. As a special case \t and \n are parsed as the tab
  and newline characters respectively.

  If -float is used, -fmt must be adjusted to account for floats.

  %[1]s is inclusive; note that if -incr moves the value beyond stop, the last
  number might not be present in the output.

OPTIONS
%[2]s
VERSION
  version:  %[3]s
  compiled: %[4]s
  built:    %[5]s

`
