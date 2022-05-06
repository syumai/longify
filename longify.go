package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Longify(r io.Reader, w io.Writer, repeatLen, longifiedPos int) error {
	if longifiedPos == 0 {
		centerPos, tmpfile, err := getCenterPos(r)
		if tmpfile != nil {
			defer os.Remove(tmpfile.Name())
			defer tmpfile.Close()
		}
		if err != nil {
			return err
		}
		r = tmpfile
		longifiedPos = centerPos
	}

	var pos int
	scanner := bufio.NewScanner(r)
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

func getCenterPos(r io.Reader) (int, *os.File, error) {
	tmpfile, err := ioutil.TempFile("", "longify")
	if err != nil {
		return 0, nil, fmt.Errorf("error creating tmpfile: %w", err)
	}

	if _, err := io.Copy(tmpfile, r); err != nil {
		return 0, tmpfile, fmt.Errorf("error copying input to tmpfile: %w", err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		return 0, tmpfile, fmt.Errorf("error seeking tmpfile: %w", err)
	}

	var lines int
	for buf := make([]byte, 1); err != io.EOF; _, err = tmpfile.Read(buf) {
		if buf[0] == '\n' {
			lines++
		}
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		return 0, tmpfile, fmt.Errorf("error seeking tmpfile: %w", err)
	}
	return lines / 2, tmpfile, nil
}
