package main

import "strings"

func defangIPaddr(address string) string {

	var sb strings.Builder

	for _, ch := range address {

		if ch == '.' {
			sb.WriteString("[.]")
		} else {
			sb.WriteRune(rune(ch))
		}

	}

	return sb.String()
}
