package StrToInt1

import (
	"errors"
	"testing"
)

func TestMystr (t *testing.T){
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
		if num, got := MyStrToInt1(test.input); got != test.want {
			if num !=test.n {
				t.Errorf("myStrToInt(%q) = %v", test.input, num)
			}
		}
	}
}
