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
	trimmedData := strings.TrimSpace(string(data))
	CheckError(err)

	lines := strings.Split(trimmedData, "\n")

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

func Filter[I interface{}](arr []I, fn func(I) bool) []I {
	result := make([]I, 0)

	for _, v := range arr {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

func Reduce[I interface{}, O interface{}](arr []I, fn func(O, I, int) O, initial O) O {
	result := initial

	for idx, v := range arr {
		result = fn(result, v, idx)
	}

	return result
}

func Every[I interface{}](arr []I, fn func(I) bool) bool {
	for _, v := range arr {
		if !fn(v) {
			return false
		}
	}

	return true
}
