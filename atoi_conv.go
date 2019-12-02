// Implementation using Atoi
package convert

import (
	"fmt"
	"strconv"
)

func convAtoi(s string) (int,error) {
	if _, err := strconv.Atoi(s); err != nil {
		return 0, fmt.Errorf("Not correct input.")
	}
	return strconv.Atoi(s)
}