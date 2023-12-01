package utils

import (
	"flag"
	"os"
	"strings"
)

func Setup() ([]string, int) {
	fileName := flag.String("f", "input.txt", "-f <filename>")
	part := flag.Int("p", 1, "-p <partNr>")
	flag.Parse()

	data, err := os.ReadFile(*fileName)
	checkError(err)

	lines := strings.Split(string(data), "\n")

	return lines, *part

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
