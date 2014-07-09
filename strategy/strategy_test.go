package strategy

import (
	"testing"
)

func TestAddStrategy(t *testing.T) {
	expect := 2
	calc := CalcAdd{}
	result := calc.Execute(1, 1)
	
	if result != expect {
		t.Errorf("Expect result to %s, but %s", result, expect)
	}
}

func TestMultStrategy(t *testing.T) {
	expect := 1
	calc := CalcMult{}
	result := calc.Execute(1, 1)
	
	if result != expect {
		t.Errorf("Expect result to %s, but %s", result, expect)
	}
}
