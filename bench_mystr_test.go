// Import StrToInt1 and do Benchmark tests
package StrToInt2

import (
	"testing"
     "StrToInt1" // go get StrToInt
)

func BenchmarkStrToIntAtoi(b *testing.B) {
	for i:= 0; i < b.N; i++{
		StrToInt1.MyStrToInt1("234")
	}
}

func BenchmarkStrToInt2Sscanf(b *testing.B) {
	for i:= 0; i < b.N; i++{
		myStrToInt2("234")
	}
}

//(base) MacBook-Air-Sasha:StrToInt2 sashavorobyova$ go test -bench=. bench_mystr_test.go mystr2.go
//goos: darwin
//goarch: amd64
//BenchmarkStrToIntAtoi-4         61564124                17.9 ns/op
//BenchmarkStrToInt2Sscanf-4       1657227               723 ns/op
//PASS
//ok      command-line-arguments  4.001s

