package nifptvalidator

import "testing"

func TestValidator(t *testing.T) {
	if !IsValidNif("501442600") {
		t.Error("Expected NIF to be valid")
	}
	if IsValidNif("740898379") {
		t.Error("Expected NIF to be invalid")
	}
	if IsValidNif("1111111111111111111") {
		t.Error("Expected NIF to be invalid length")
	}
	if IsValidNif("ASDFGHJKL") {
		t.Error("Expected NIF to have invalid chars")
	}
	if IsValidNif("2323B323A") {
		t.Error("Expected NIF to have an invalid chars")
	}
	if IsValidNif("23232323A") {
		t.Error("Expected NIF to have an invalid checksum")
	}
	if !IsValidNif("999999990") {
		t.Error("Expected 9999990 to be a valid NIF")
	}
	if !IsValidNif("123456789") {
		t.Error("Expected 123456789 to be a valid NIF")
	}
}
