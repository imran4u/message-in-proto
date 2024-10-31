package basic

import (
	"fmt"
	"my-protobuf/protogen/basic"
)

func BasicHello() {
	hello := &basic.Hello{Name: "Imran ali"}
	fmt.Println(hello.Name)
}
