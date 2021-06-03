package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func loopFloat() {
	start, stop, incr := parseFloats(flag.Arg(0), flag.Arg(1), incr)

	i := start
	for shouldStopFloat(i, stop, start < stop) {
		fmt.Fprintf(os.Stdout, fmtString, i)
		i += incr

		if !shouldStopFloat(i, stop, start < stop) {
			break
		}

		fmt.Fprint(os.Stdout, delim)
	}

	// for i := start; shouldStopFloat(i, stop, start < stop); i += incr {
	// 	fmt.Fprintf(os.Stdout, fmtString, i)
	// }

	fmt.Fprint(os.Stdout, "\n")

}

func shouldStopFloat(curr, stop float64, up bool) bool {
	if up && (curr > stop) {
		return false
	}
	if !up && (curr < stop) {
		return false
	}

	return true
}

func parseFloats(i1, i2 string, incr string) (float64, float64, float64) {
	diff, err := strconv.ParseFloat(incr, 64)
	if err != nil {
		stderr.Fatal(err)
	}

	if diff == 0 {
		stderr.Fatal("using -incr 0 will never stop")
	}

	start, err := strconv.ParseFloat(i1, 64)
	if err != nil {
		stderr.Fatal(err)
	}

	stop, err := strconv.ParseFloat(i2, 64)
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
