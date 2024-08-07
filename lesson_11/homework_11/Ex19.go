package main

import "strings"

type DashString string

func (ds DashString) Transform() DashString {
	return DashString(strings.ReplaceAll(string(ds), " ", "-"))
}
