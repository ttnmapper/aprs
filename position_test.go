package aprs

import (
	"fmt"
	"testing"
)

var testPosition = Position{
	Latitude:    35.7,
	Longitude:   -78.7,
	Altitude:    100,
	SymbolTable: "D",
	Symbol:      ">",
	Message:     "Test Position",
}

//func init() {
//	testPosition.Zero()
//}

func TestPosition_String(t *testing.T) {
	p := testPosition
	result := fmt.Sprintf("%s", p)
	if result != "!3542.00ND07842.00W>/A=000328Test Position" {
		t.Fatalf("%s != !3542.00ND07842.00W>/A=000328Test Position", result)
	}

	p.Zero()
	result = fmt.Sprintf("%s", p)
	if result != "!0000.00N\\00000.00E." {
		t.Fatalf("%s != !0000.00N\\00000.00E.", result)
	}
}
