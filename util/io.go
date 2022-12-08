package util

import (
	"bufio"
	"os"
)

func ReadByLines(path string) <-chan string {
	c := make(chan string)

	file, err := os.Open(path)
	Must(err)

	go func() {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			line := scanner.Text()
			c <- line
		}
		close(c)
	}()

	return c
}
