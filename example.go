package main

import (
	"fmt"

	"github.com/hugorosario/nifptvalidator/validator"
)

func testNif(nif string) {
	fmt.Println(nif, "=>", validator.IsValidNif(nif))
}

func main() {
	testNif("23232323A") //true
	testNif("501442600") //true
	testNif("999999990") //true, not a real assigned NIF, but is valid and is used as Consumidor Final
	testNif("740898379") //false
	testNif("987654321") //false
	testNif("987654321") //false
	testNif("123456789") //true, seems strange, but this is a valid NIF according to the specs
}
