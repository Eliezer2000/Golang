package main

import "testing"


type test struct{
	data []int
	answer int
}

func TestSomaEmTabela(t *testing.T ){
	tests := []test {
		test{data: []int{1, 2, 3}, answer: 6},
		test{[]int{10, 11, 12}, 33}, 
		test{[]int{-5, 0, 5, 10}, 10},
	}
	for _, v := range tests {
		z := soma(v.data...)
		if z != v.answer {
			t.Error("Expected: ", v.answer, "Got: ", z)
		} 
	}
}

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