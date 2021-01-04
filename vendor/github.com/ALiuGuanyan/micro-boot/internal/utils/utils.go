package utils

import (
	"fmt"
	"strings"
)

func BuildFlagName(pre, deft string) string {
	if pre != "" {
		if strings.TrimSpace(deft) == "" {
			return pre
		}
		return fmt.Sprintf("%s-%s", pre, deft)
	}
	return deft
}

