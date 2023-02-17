package funfacts

import (
	"reflect"
	"testing"
)

/*
*

	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
func TestGetFunFacts(t *testing.T) {
	type test struct {
		input string // her må du skrive riktig type for input
		want  string // her må du skrive riktig type for returverdien
	}

	// Her må du legge inn korrekte testverdier
	//tests := []test{
	//  {input: , want: },
	//}
	tests := []test{
		{input: "sun", want: "Sola"},
		{input: "luna", want: "Månen"},
		{input: "terra", want: "Jorda"},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
