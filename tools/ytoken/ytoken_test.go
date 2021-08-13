package ytoken

import "testing"

func TestYToken_Timeout(t *testing.T) {
	yToken := New(&Constraint{
		Signer:      "",
		Expiry:      0,
		Serial:      "",
		Beneficiary: "",
	})
	println(yToken.Timeout())
}
