package utils

import (
	"flag"
	"os"
	"strings"
)

func Setup() []string {
	fileName := flag.String("f", "input.txt", "-f <filename>")
	flag.Parse()

	data, err := os.ReadFile(*fileName)
	CheckError(err)

	lines := strings.Split(string(data), "\n")

	return Mapper(lines, func(line string) string {
		return strings.TrimSpace(line)
	})

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Mapper[I interface{}, O interface{}](arr []I, fn func(I) O) []O {
	result := make([]O, len(arr))

	for i, v := range arr {
		result[i] = fn(v)
	}

	return result
}
