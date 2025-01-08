package accessprivate

import (
	"fmt"
	"reflect"
	"testing"
)


func TestGetPrivateVarThroughReflect(t *testing.T) {
	p := Person{name: "shixun", age: 28}

	v := reflect.ValueOf(&p).Elem()

	nameField := v.FieldByName("name")

	fmt.Println("name (private):", nameField.String())
}