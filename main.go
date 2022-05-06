package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	length       = flag.Int("l", 0, "length of repeated line")
	longifiedPos = flag.Int("p", 0, "position to be longified")
)

func main() {
	flag.Parse()
	repeatLen := *length
	if repeatLen == 0 {
		repeatLen = rand.Int()%10 + 3
	}
	if err := Longify(os.Stdin, os.Stdout, repeatLen, *longifiedPos); err != nil {
		fmt.Fprintln(os.Stderr, "error: %w", err)
		os.Exit(1)
	}
}
