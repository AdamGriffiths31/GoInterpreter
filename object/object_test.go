package object

import "testing"

func TestStringHashkey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.Hashkey() != hello2.Hashkey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.Hashkey() != diff2.Hashkey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.Hashkey() == diff1.Hashkey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
