package main

import (
	"strings"
	"testing"
)

var args = []string{"abc", "dafd", "akldjo", "akljfa", "üijol", "joidh", "ohoadfh"}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for _, arg := range args {
			s += sep + arg
			sep = " "
		}
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
}
