package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func loopInt() {
	start, stop, incr := parseInts(flag.Arg(0), flag.Arg(1), incr)

	i := start
	for shouldStopInt(i, stop, start < stop) {
		fmt.Fprintf(os.Stdout, fmtString, i)
		i += incr

		if !shouldStopInt(i, stop, start < stop) {
			break
		}

		fmt.Fprint(os.Stdout, delim)
	}

	// for i := start; shouldStopInt(i, stop, start < stop); i += incr {
	// 	fmt.Fprintf(os.Stdout, fmtString, i)
	// 	fmt.Fprint(os.Stdout, delim)
	// }

	if delim != "\n" {
		fmt.Fprint(os.Stdout, "\n")
	}

}

func shouldStopInt(curr, stop int64, up bool) bool {
	if up && (curr > stop) {
		return false
	}
	if !up && (curr < stop) {
		return false
	}
	return true
}

func parseInts(i1, i2 string, incr string) (int64, int64, int64) {
	diff, err := strconv.ParseInt(incr, 10, 64)
	if err != nil {
		stderr.Fatal(err)
	}

	if diff == 0 {
		stderr.Fatal("using -incr 0 will never stop")
	}

	start, err := strconv.ParseInt(i1, 10, 64)
	if err != nil {
		stderr.Fatal(err)
	}

	stop, err := strconv.ParseInt(i2, 10, 64)
	if err != nil {
		stderr.Fatal(err)
	}

	if diff < 0 && (stop > start) {
		stderr.Fatal("using -incr <=0 with numbers that increase will never stop")
	}

	if diff > 0 && (start > stop) {
		stderr.Fatal("using -incr >=0 with numbers that decrease will never stop")
	}

	return start, stop, diff
}
