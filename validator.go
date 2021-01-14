/*
Package nifptvalidator implements a validator for the Portuguese NIF (Número de Identificação Fiscal).
It follow the specs defined in the wikipedia page at  https://pt.wikipedia.org/wiki/N%C3%BAmero_de_identifica%C3%A7%C3%A3o_fiscal

Example:

	package main

	import (
		"fmt"

		"github.com/hugorosario/nifptvalidator"
	)

	func main() {
		nif := "123456789"
		fmt.Println(nif, "=>", nifptvalidator.IsValidNif(nif))
	}

*/
package nifptvalidator

import (
	"strconv"
)

//IsValidNif validates a NIF according to the rules specified in https://pt.wikipedia.org/wiki/N%C3%BAmero_de_identifica%C3%A7%C3%A3o_fiscal
func IsValidNif(nif string) bool {
	//validate length
	if len(nif) != 9 {
		return false
	}

	//validate prefixes
	validPrefixes := map[string]bool{
		"1": true, "2": true, "3": true, "5": true, "6": true, "8": true,
		"45": true, "70": true, "71": true, "72": true, "74": true, "75": true, "77": true, "78": true, "79": true, "90": true, "91": true, "98": true, "99": true,
	}
	if !validPrefixes[nif[:1]] && !validPrefixes[nif[:2]] {
		return false
	}

	//calculate check-digit
	sum := 0
	for i := 1; i < 9; i++ {
		v, err := strconv.Atoi(string(nif[i-1]))
		if err != nil {
			return false
		}
		sum += v * (10 - i)
	}
	rmd := sum % 11
	ckd := 0
	switch rmd {
	case 0, 1:
		ckd = 0
	default:
		ckd = 11 - rmd
	}

	//compare the provided check digit with the calculated one
	compare, err := strconv.Atoi(string(nif[8]))
	if err != nil {
		return false
	}
	return compare == ckd
}
