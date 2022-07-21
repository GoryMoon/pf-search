package utils

import (
	"regexp"
	"strconv"
	"strings"
)

var cleanRegex = regexp.MustCompile("[^a-zA-Z\\d-_]+")

func CleanString(str string) string {
	return cleanRegex.ReplaceAllString(str, "")
}

func SplitAndTrim(str string, sep string) (out []string) {
	out = make([]string, 0)
	for _, s := range strings.Split(str, sep) {
		s = strings.TrimSpace(s)
		if s != "" {
			out = append(out, s)
		}
	}
	return
}

func FilterEmptyStrings(s []string) (out []string) {
	for _, str := range s {
		if str != "" {
			out = append(out, str)
		}
	}
	return
}

func ParseInt(str string, def int) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		return def
	}
	return res
}

func ParseBool(str string, def bool) bool {
	res, err := strconv.ParseBool(str)
	if err != nil {
		return def
	}
	return res
}
