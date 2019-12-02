//Realization using Sscanf
package StrToInt2

import (
	"fmt"
)

func myStrToInt2(s string) (int, error) {
	var i int
	if _, err := fmt.Sscanf(s, "%d ", &i); err != nil {
		return 0, fmt.Errorf("Not correct input.")
	}
	return i, nil
}

//Faster variant, but int64 returns
//func myStrToInt2(s string) (int64, error) {
//	if _, err := strconv.ParseInt(s, 10,0); err != nil {
//		return 0, fmt.Errorf("Not correct input.")
//	}
//	return strconv.ParseInt(s, 10,0)
//}

//tests are the same as for the first implementation
//(base) MacBook-Air-Sasha:StrToInt2 sashavorobyova$ go test
//PASS
//ok      StrToInt2       0.006s
