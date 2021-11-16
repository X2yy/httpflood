package main

import (
	"os"
	"strings"
)

func Catar() string {
	banner := strings.ReplaceAll(Banner, "%s.us.path%", os.Args[0])
	return banner
}
