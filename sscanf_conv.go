// Implementation using Sscanf
package convert

import "fmt"

func convSScanf(s string) (int, error) {
	var i int
	if _, err := fmt.Sscanf(s, "%d ", &i); err != nil {
		return 0, fmt.Errorf("Not correct input.")
	}
	return i, nil
}

