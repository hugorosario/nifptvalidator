package main

import (
	"fmt"

	"github.com/hugorosario/nifptvalidator/validator"
)

func testValidation(nif string) {
	fmt.Println(nif, "=>", validator.IsValidNif(nif))
}

func main() {
	testValidation("501442600")
	testValidation("999999990") //not a real assigned NIF, but is valid and used as a Consumidor Final
	testValidation("740898379")
	testValidation("987654321")
	testValidation("987654321")
	testValidation("123456789") //seems strange, but this is a valid NIF according to the specs
}
