package main

import "strconv"

func ToString(source float64) string {
	return strconv.FormatFloat(source, 'f', 6, 64)
}




