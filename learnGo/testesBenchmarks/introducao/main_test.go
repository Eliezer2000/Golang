package main

import "testing"

func TestSoma(t *testing.T){
	test := soma(1, 2, 3)
	resul := 6
	if test != 6 {
		t.Error("Expected: ", resul, "Got: ", test)
	}
}


func TestMultiplica(t *testing.T){
	test := multiplica(10, 10)
	resul := 100
	if test != resul {
		t.Error("Expected: ", resul, "Got: ", test)
	}
}