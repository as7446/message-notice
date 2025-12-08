package cliutil

import "strings"

// SplitCSV 将以逗号分隔的字符串拆分为切片，自动去除空元素与首尾空格。
func SplitCSV(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	var out []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		out = append(out, p)
	}
	return out
}
