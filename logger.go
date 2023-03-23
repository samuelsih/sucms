package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var length = 0

func info(msg string) {
	length++
	fmt.Printf("[%d] %s\n", length, msg)
}

func writeStream(r io.ReadCloser, f func(doneChan chan struct{})) {
	defer r.Close()

	scanner := bufio.NewScanner(r)
	done := make(chan struct{})

	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		done <- struct{}{}
	}()

	f(done)
}

func exit[T any](msg T) {
	fmt.Println(msg)
	os.Exit(1)
}
