// Implementation using Atoi
package StrToInt1

import (
	"fmt"
	"strconv"
)

func MyStrToInt1(s string) (int,error) {
	if _, err := strconv.Atoi(s); err != nil {
		return 0, fmt.Errorf("Not correct input.")
	}
	return strconv.Atoi(s)
}

//(base) MacBook-Air-Sasha:StrToInt1 sashavorobyova$ go test
//PASS
//ok      StrToInt1       0.006s