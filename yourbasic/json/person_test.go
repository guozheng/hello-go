package json

import (
	"fmt"
	"testing"
)

func TestToAndFromJson(t *testing.T) {
	c := Company{Name: "Google"}
	p := Person{Name: "Mike", Age: 30, Employers: []Company{c}}

	encoded, err := ToJson(p)
	if err != nil {
		t.Errorf("ToJson should not get error")
	}
	fmt.Printf("encoded: %v\n", string(encoded))

	decoded, err := FromJson(encoded)
	if err != nil {
		t.Errorf("FromJson should not get error")
	}
	fmt.Printf("decoded: %v\n", decoded)

}