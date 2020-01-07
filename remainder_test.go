package main

import (
	"fmt"
	"testing"
)

var r int

func benchmarkRemainder(numerators, denominators []int, b *testing.B) {
	for _, num := range numerators {
		for _, den := range denominators {
			b.Run(fmt.Sprintf("%d/%d", num, den), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					r = num % den
				}
			})
		}
	}
}

func BenchmarkRemainderPowersOf2(b *testing.B) {
	numerators := []int{256, 512, 1024, 1024 * 1024, 1024 * 1024 * 1024}
	denominators := []int{256, 512, 1024, 1024 * 1024, 1024 * 1024 * 1024}

	benchmarkRemainder(numerators, denominators, b)
}

func BenchmarkRemainderNonPowersOf2(b *testing.B) {
	numerators := []int{255, 500, 1023, 1021 * 1032, 1029 * 1010 * 1009}
	denominators := []int{255, 500, 1023, 1020 * 1024, 1029 * 1010 * 1009}

	benchmarkRemainder(numerators, denominators, b)
}
