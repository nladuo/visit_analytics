package main

import (
	"strings"
)

func GetHostName(url string) string {
	host_name := strings.TrimLeft(url, "http://")
	host_name = strings.TrimLeft(host_name, "https://")
	strs := strings.Split(host_name, "/")
	if len(strs) == 0 {
		return ""
	}
	return strs[0]
}
