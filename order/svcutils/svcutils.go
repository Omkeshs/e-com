package svcutils

import (
	"fmt"
	"strconv"
)

// GetIntPathVariable retrieves int value for the path variable
func GetIntPathVariable(vars map[string]string, variable string) (int32, error) {
	varStr, exists := vars[variable]
	if !exists {
		return 0, fmt.Errorf(fmt.Sprintf("invalid value for %v", variable))
	}

	i, conversionError := strconv.ParseInt(varStr, 10, 32)
	if conversionError != nil {
		return 0, fmt.Errorf(fmt.Sprintf("invalid value for %v", variable))
	}
	return int32(i), nil
}
