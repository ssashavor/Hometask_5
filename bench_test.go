package convert

import (
	"testing"
)

func Benchmark_Atoi(b *testing.B) {
	for i:= 0; i < b.N; i++{
		convAtoi("234")
	}
}

func Benchmark_Sscanf(b *testing.B) {
	for i:= 0; i < b.N; i++{
		convSScanf("234")
	}
}

//(base) MacBook-Air-Sasha:convert sashavorobyova$ go test -bench=. bench_test.go atoi_conv.go sscanf_conv.go
//goos: darwin
//goarch: amd64
//Benchmark_Atoi-4        66309339                18.0 ns/op
//Benchmark_Sscanf-4       1666868               720 ns/op
//PASS
//ok      command-line-arguments  3.154s
