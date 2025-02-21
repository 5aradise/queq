package main

import "strconv"

const maxFloatPrecision = 16

func formatFloat(f float64) string {
	fmtF := strconv.FormatFloat(f, 'f', maxFloatPrecision, 64)
	ld := len(fmtF) - 1
	for fmtF[ld] == '0' {
		ld--
	}
	if fmtF[ld] == '.' {
		ld--
	}
	return fmtF[:ld+1]
}
