package accessprivate

import (
	"fmt"
	"reflect"
	"testing"
)

/*
在 Go 语言里，可寻址（addressable） 的概念对反射非常重要。如果你只是传递 p（一个非指针结构体），则得到的 reflect.Value 往往是 只读 的，无法进行赋值操作或者对其字段进行“set”操作。因为在 Go 的语义里，传递 p 得到的只是一个值拷贝，不是真正原始变量的引用。
  - 若只用 reflect.ValueOf(p)：

返回的 reflect.Value 并不具备“可寻址性”，也就无法修改原结构体的字段。
  - 若用 reflect.ValueOf(&p)：

返回的是一个“指向 p 的指针”的 reflect.Value。再通过 .Elem() 得到的才是真正指向 p 实体的可寻址 reflect.Value，这样才能对结构体字段进行读取和修改（若需要）。

简单来说，想要通过反射“读写”原始对象的字段（尤其是写），就必须让反射层拿到指针并且处于可寻址状态。这也是为什么示例中常见 reflect.ValueOf(&p).Elem() 这种用法。

在实际编程中，这个模式非常常见，因为多数反射应用场景不仅仅是读取，还包括“根据标签自动赋值”、“动态设置字段”等，因此 reflect.ValueOf(&p).Elem() 就变成了一个标准的范式。
*/
func TestGetPrivateVarThroughReflect(t *testing.T) {
	p := Person{name: "shixun", age: 28}

	v := reflect.ValueOf(&p).Elem()

	nameField := v.FieldByName("name")

	fmt.Println("name (private):", nameField.String())
}
