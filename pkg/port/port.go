// Package port provides utilities for working with network ports.
package port

import (
	"fmt"
	"strconv"
)

// ToString converts an integer port to an address string.
func ToString(p int) string {
	strPort := strconv.Itoa(p)
	return fmt.Sprintf(":%s", strPort)
}
