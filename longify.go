package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Longify(r io.Reader, w io.Writer, repeatLen, longifiedPos int) error {
	tmpfile, err := ioutil.TempFile("", "longify")
	if err != nil {
		return fmt.Errorf("error creating tmpfile: %w", err)
	}

	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	if _, err := io.Copy(tmpfile, r); err != nil {
		return fmt.Errorf("error copying input to tmpfile: %w", err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		return fmt.Errorf("error seeking tmpfile: %w", err)
	}

	if longifiedPos == 0 {
		longifiedPos = getCenterPos(tmpfile)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		return fmt.Errorf("error seeking tmpfile: %w", err)
	}

	var pos int
	scanner := bufio.NewScanner(tmpfile)
	for scanner.Scan() {
		pos++
		i := 1
		if pos == longifiedPos {
			i = repeatLen
		}
		for ; i > 0; i-- {
			if _, err := fmt.Fprintln(w, scanner.Text()); err != nil {
				return fmt.Errorf("error printing line to writer: %w", err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner err: %w", err)
	}
	return nil
}

func getCenterPos(r io.Reader) int {
	var (
		lines int
		err   error
	)
	for buf := make([]byte, 1); err != io.EOF; _, err = r.Read(buf) {
		if buf[0] == '\n' {
			lines++
		}
	}
	return lines / 2
}
