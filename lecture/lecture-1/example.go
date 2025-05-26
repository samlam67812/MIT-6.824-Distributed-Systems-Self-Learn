package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Map(key string, value string) {
	// key: document name
	// value: document contents
	words := strings.Fields(value)
	for _, w := range words {
		// "1" as example
		EmitIntermediate(w, "1")
	}
}

func EmitIntermediate(key string, value string) {
	// Implementation to store the intermediate key-value pairs
}

func Reduce(key string, values []string) {
	result := 0

	for _, v := range values {
		count, err := strconv.Atoi(v)
		if err != nil {
			result += count
		}
	}

	// Emit the final count as string
	Emit(fmt.Sprintf("%d", result))
}

func Emit(value string) {
	// Implementation to output the final result
}