package convert

import (
	"errors"
	"testing"
)

func Test_byAtoi (t *testing.T){
	var tests = []struct{
		input string
		n int
		want error
	}{
		{"4", 4, nil},
		{"543725638", 543725638, nil},
		{"-6", -6, nil},
		{"0", 0, nil},
		{"0.0", 0, errors.New( "Not correct input.")},
		{"2.1", 0, errors.New( "Not correct input.")},
		{".1", 0, errors.New( "Not correct input.")},
		{"2a", 0, errors.New( "Not correct input.")},
		{"", 0, errors.New( "Not correct input.")},
		{"  ", 0, errors.New( "Not correct input.")},
		{"hello", 0, errors.New( "Not correct input.")},
		{"世界", 0, errors.New( "Not correct input.")},
	}
	for _,test :=range tests{
		if num, got := convAtoi(test.input); got != test.want {
			if num !=test.n {
				t.Errorf("convAtoi(%q) = %v", test.input, num)
			}
		}
	}
}

//(base) MacBook-Air-Sasha:convert sashavorobyova$ go test atoi_test.go atoi_conv.go
//ok      command-line-arguments  0.006s

