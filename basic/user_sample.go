package basic

import (
	"fmt"
	"my-protobuf/protogen/basic"
)

func BasicUser() {
	user := &basic.User{Id: 1, UserName: "Imran Ali", IsActive: true}
	user.Password = []byte("mySecret")
	fmt.Println(user.String())
}
